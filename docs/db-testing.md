# TESTING SQL DEPLOYMENT

Different Hyperscaler return different secrets. To find the DB endpoit in each of them use:
- For AWS run ```export HOST_KEY=endpoint```
- For GCP run ```export HOST_KEY=host```
- For Azure run ```export HOST_KEY=publicIP```

Chose environment:
- For dev run ```export ENVIRONMENT=dev```
- For stage run ```export HOST_KEY=stage```
- For prod run ```export HOST_KEY=prod```

Then
```bash
export DB=my-app-backend-dev-db

export PGUSER=$(kubectl --namespace my-app-$ENVIRONMENT \
    get secret $DB --output jsonpath="{.data.username}" \
    | base64 -d)

export PGPASSWORD=$(kubectl --namespace my-app-$ENVIRONMENT \
    get secret $DB --output jsonpath="{.data.password}" \
    | base64 -d)

export PGHOST=$(kubectl --namespace my-app-$ENVIRONMENT \
    get secret $DB --output jsonpath="{.data.$HOST_KEY}" \
    | base64 -d)

kubectl run postgresql-client --rm -ti --restart='Never' \
    --image docker.io/bitnami/postgresql:16 \
    --env PGPASSWORD=$PGPASSWORD --env PGHOST=$PGHOST \
    --env PGUSER=$PGUSER --command -- sh
```

#### Inside the pod run:
```bash
psql --host $PGHOST -U $PGUSER -d postgres -p 5432
```
 
#### Inside the DB run:
```bash
\l
```
Output should be something like:
```sql
                                                          List of databases
   Name    |   Owner    | Encoding | Locale Provider |   Collate   |    Ctype    | ICU Locale | ICU Rules |     Access privileges
-----------+------------+----------+-----------------+-------------+-------------+------------+-----------+---------------------------
 my-db     | masteruser | UTF8     | libc            | en_US.UTF-8 | en_US.UTF-8 |            |           |
 postgres  | masteruser | UTF8     | libc            | en_US.UTF-8 | en_US.UTF-8 |            |           |
 rdsadmin  | rdsadmin   | UTF8     | libc            | en_US.UTF-8 | en_US.UTF-8 |            |           | rdsadmin=CTc/rdsadmin
 template0 | rdsadmin   | UTF8     | libc            | en_US.UTF-8 | en_US.UTF-8 |            |           | =c/rdsadmin              +
           |            |          |                 |             |             |            |           | rdsadmin=CTc/rdsadmin
 template1 | masteruser | UTF8     | libc            | en_US.UTF-8 | en_US.UTF-8 |            |           | =c/masteruser            +
           |            |          |                 |             |             |            |           | masteruser=CTc/masteruser
```

Where "my-db" is the DB created through the composition.

#### Delete postgresql-client pod:
```bash
kubectl delete pod postgresql-client
```


### When deleting some resources will be left hanging
```bash
kubectl patch database.postgresql.sql.crossplane.io $DB \
    --patch '{"metadata":{"finalizers":[]}}' --type=merge

kubectl patch object.kubernetes.crossplane.io $DB \
    --patch '{"metadata":{"finalizers":[]}}' --type=merge
```
