<a href="https://www.instagram.com/ttomasferrari/">
    <img align="right" alt="Abhishek's Instagram" width="22px" 
    src="https://i.imgur.com/EzpyGdV.png" />
</a>
<a href="https://twitter.com/tomasferrari">
    <img align="right" alt="Abhishek Naidu | Twitter" width="22px"         
    src="https://i.imgur.com/eFVBTVz.png" />
</a>
<a href="https://www.linkedin.com/in/tomas-ferrari-devops/">
    <img align="right" alt="Abhishek's LinkedIN" width="22px" 
    src="https://i.imgur.com/pMzVPqj.png" />
</a>
<p align="right">
    <a >Find me here: </a>
</p>
<!-- <p align="right">
    <a  href="/docs/readme_es.md">Versión en Español</a>
</p> -->






HACER SCRIPT DE DESTROY

kubectl --namespace my-cluster \
    get secret my-cluster \
    --output jsonpath="{.data.kubeconfig}" \
    | base64 -d

kubectl --namespace my-cluster \
    get secret my-cluster \
    --output jsonpath="{.data.kubeconfig}" \
    | base64 -d > kubeconfig.yaml

<p title="Banner" align="center"> <img src="https://i.imgur.com/FbsIwSJ.jpg"> </p>

# INDEX

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Some Crossplane Concepts](#some-crossplane-concepts)
- [What We'll Be Deploying](#what-well-be-deploying)
  - [My App](#my-app)
  - [Standalone EKS Cluster](#standalone-eks-cluster)
- [Crossplane Packages](#optional-crossplane-packages)
- [Initial Setup](#initial-setup)
- [Run Lab](#run-lab)
- [Conclusion](#conclusion)

</br>
</br>

# INTRODUCTION
This is a spin-off of my [Automate All The Things](https://github.com/tferrari92/automate-all-the-things) DevOps project. While working on the [Nirvana Edition](https://github.com/tferrari92/automate-all-the-things-nirvana) I'm creating this smaller lab for anyone who wants to start experimenting with this tool.

....

We'll be using a GitOps methodology with Helm, ArgoCD and the App Of Apps Pattern. There is some extra information [here](/docs/argocd-notes.md), but you are expected to know about these things.

</br>
</br>

# PREREQUISITES
- Active DockerHub account
- Minikube installed
- kubectl installed
- Helm installed
- (Optional) docker cli installed
- (Optional) crossplane cli installed

</br>
</br>

# SOME CROSSPLANE CONCEPTS

<p title="Diagrama fundamentales" align="center"> <img src="https://i.imgur.com/rBLyH8I.jpg"> </p>

Wont go into details.

### Operations team

1. Create CompositeResourceDefinition: 
2. Create Composition: 
3. Create Providers:
4. Create ProviderConfig

### Dev team

Will only create a ClusterClaim manifest. They could create a Cluster but we usually want resources to be namescpace scoped so we use CLusterClaim over Cluster.

What will happen next:
1. 



1. CompositeResourceDefinitions: Extend k8s API by creating CRDs. Define what parameters are available and required for a Composition that uses this XRD
2. Compositions
3. CompositeResources

</br>
</br>

# WHAT WE'LL BE DEPLOYING
## My-App
My-app is composed of a frontend service and backend service.

### Frontend service
The fronted service is composed of a Deployment, a Service, an Horizontal Pod Autoscaler and...  But if we look at the [helm chart](/helm-charts/my-app/frontend/), we'll find only the [AppClaim manifest](/helm-charts/systems/my-app/backend/templates/app-claim.yaml). Why is this?

We are using Crossplane to define what a frontend application looks like. Through the use of the [Frontend Application Composition](/helm-charts/infra/crossplane-compositions/application/frontend-composition.yaml) we can define exactly how a frontend application must be deployed. [Here's a video](https://youtu.be/eIQpGXUGEow?si=nsm-uR1AyGZFbf6y) further explaining this concept.

</br>

### Backend service
The backend service will include the backend appication but also it's required database.

Let's take a look a the my-app backend helm chart. In this case we'll find two manifest templates: [app-claim.yaml](/helm-charts/my-app/backend/templates/app-claim.yaml) and [sql-claim.yaml](/helm-charts/my-app/backend/templates/sql-claim.yaml)

As for the frontend, we have a [Backend Application Composition](/helm-charts/infra/crossplane-compositions/application/backend-composition.yaml) which defines how a backend application is deployed.

But in this case we also have the [SQL Claim](/helm-charts/my-app/backend/templates/sql-claim.yaml).
junto c0n el  [SQLClaim](/helm-charts/systems/my-app/backend/templates/sql-claim.yaml) creamos el secret que contiene la password que el sqlclaim va a usar para ponerle a la db. ESTO DEBERIA ENCONTRA UNA FORMA MEJOR DE HACERLONAPRAS QUE NO QUED EEL SECRET AHI EXPUESTO
 RDS instance deployed in AWS.


what does the app backed composition include?
ProviderConfig (required for deploying Kuberntes objects within this same Minikube cluster), NECESITO EL PROVIDER Q SE INSTALA EN EL CHART DE PROVIDERSA??????deplyomenty and service. we could have nested within the backend composition an sql composition For the sake of simplicity and understability, we'll keep the backend's [AppClaim](/helm-charts/systems/my-app/backend/templates/app-claim.yaml) and [SQLClaim](/helm-charts/systems/my-app/backend/templates/sql-claim.yaml) separated. We could have included an SQLClaim within the [Backend App composition](/helm-charts/infra/crossplane-compositions/application/backend-composition.yaml).






</br>

## Standalone EKS Cluster
A standalone EKS cluster. This cluster is unrelated to our my-app applciation, 

Comes with Prometheus and ArgoCD
GitOps ready: https://youtu.be/AVHyltqgmSU?si=bV2U4OLCUFrgNhym


</br>
</br>


# (OPTIONAL) CROSSPLANE PACKAGES
<p title="Diagrama packages" align="center"> <img src="https://i.imgur.com/5CW8ZyB.jpg"> </p>

We wont use Crossplane Packages in this example so you can see all the moving parts. We could use packages like this.....
PROVIDERS, CONFIGURATIONS AND FUNCTIONS ARE ALL TYPES of "PACKAGES"
CAN BE LISTED WITH kubectl get pkgrev

INCLUIR COMANDOS PARA CREAR Y SUBIR PACKAGE

```bash
cd crossplane-configuration-packages

chmod +x build-and-push-package.sh
./build-and-push-package.sh
```

Explicar despues como uno utilizaria el package

</br>
</br>


# INITIAL SETUP

### Fork and clone the repo
Let's turn this whole deployment into your own thing.

1. Fork this repo. Keep the repository name "crossplane-minikube-lab".
1. Clone the repo from your fork:

```bash
git clone https://github.com/<your-github-username>/crossplane-minikube-lab.git
```

2. Move into the directory:

```bash
cd crossplane-minikube-lab
```

2. Run the initial setup script. Come back when you are done:

```bash
chmod +x initial-setup.sh
./initial-setup.sh
```

4. Commit and push your customized repo to GitHub:

```bash
git add .
git commit -m "customized repo"
git push
```

</br>
</br>

# RUN LAB
If you have a Minikube cluster running, delete it first with `minikube delete`.

Now run the deploy-in-minikube.sh script to get everything setup:
```bash
chmod +x deploy-in-minikube.sh
./deploy-in-minikube.sh
```
</br>

Now go to localhost:8080 on your browser to access the ArgoCD UI. You'll get the credentials from deploy script.

kubectl port-forward -n my-app service/my-app-frontend 8081:80


</br>
</br>

# CONCLUSION
That's it! This is your own Crossplane implementation now. 

For more DevOps and Platform Engineering goodness, check out my [Automate All The Things](https://github.com/tferrari92/automate-all-the-things) project.

Happy automating!
