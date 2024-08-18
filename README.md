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

PROVIDERS, CONFIGURATIONS AND FUNCTIONS ARE ALL TYPES of "PACKAGES"
CAN BE LISTED WITH kubectl get pkgrev


1. CompositeResourceDefinitions: Extend k8s API by creating CRDs. Define what parameters are available and required for a Composition that uses this XRD
2. Compositions
3. CompositeResources


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
- [Initial Setup](#initial-setup)
- [Initial Setup](#initial-setup)
- [Run Backstage Locally](#run-backstage-locally)
- [Customising Backstage](#customising-backstage)
  - [OAuth With GitHub](#oauth-with-github)
- [Run Backstage In Minikube](#run-backstage-in-minikube)
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
- minikube installed
- kubectl installed
- helm installed

</br>
</br>

# WHAT WE'LL BE DEPLOYING

## My-App
My-app is composed of a frontend service and backend service.

### Frontend service
The fronted service is composed of a Deployment, a Service, an Horizontal Pod Autoscaler and...  But if we look at the [helm chart](/helm-charts/my-app/frontend/), we'll find only the [AppClaim manifest](/helm-charts/systems/my-app/backend/templates/app-claim.yaml). Why is this?

We are using Crossplane to define what a frontend application looks like. Through the use of the [Frontend Application Composition](/helm-charts/infra/crossplane-compositions/application/frontend-composition.yaml) we can define exactly how a frontend application must be deployed. [Here's a video](https://youtu.be/eIQpGXUGEow?si=nsm-uR1AyGZFbf6y) further explaining this concept.

### Backend service
The backend service will include the backend appication but also it's required database.

Let's take a look a the my-app backend helm chart. In this case we'll find two manifest templates: [app-claim.yaml](/helm-charts/my-app/backend/templates/app-claim.yaml) and [sql-claim.yaml](/helm-charts/my-app/backend/templates/sql-claim.yaml)

As for the frontend, we have a [Backend Application Composition](/helm-charts/infra/crossplane-compositions/application/backend-composition.yaml) which defines how a backend application is deployed.

But in this case we also have the [SQL Claim](/helm-charts/my-app/backend/templates/sql-claim.yaml).
junto c0n el  [SQLClaim](/helm-charts/systems/my-app/backend/templates/sql-claim.yaml) creamos el secret que contiene la password que el sqlclaim va a usar para ponerle a la db. ESTO DEBERIA ENCONTRA UNA FORMA MEJOR DE HACERLONAPRAS QUE NO QUED EEL SECRET AHI EXPUESTO
 RDS instance deployed in AWS.


what does the app backed composition include?
ProviderConfig (required for deploying Kuberntes objects within this same Minikube cluster), NECESITO EL PROVIDER Q SE INSTALA EN EL CHART DE PROVIDERSA??????deplyomenty and service. we could have nested within the backend composition an sql composition For the sake of simplicity and understability, we'll keep the backend's [AppClaim](/helm-charts/systems/my-app/backend/templates/app-claim.yaml) and [SQLClaim](/helm-charts/systems/my-app/backend/templates/sql-claim.yaml) separated. We could have included an SQLClaim within the [Backend App composition](/helm-charts/infra/crossplane-compositions/application/backend-composition.yaml).





kubectl port-forward -n my-app service/my-app-frontend 8081:80

## EKS Cluster
A standalone EKS cluster. This cluster is unrelated to our my-app applciation, 


## No Packages
we dont use packages in this example so you can see all the moving parts. We could use packages like this.....


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


</br>
</br>

# RUN BACKSTAGE IN MINIKUBE
Ok, lets run Backstage in Minikube. `Ctrl + C` to kill the `yarn dev` process.

We first need to build and push the Backstage Docker image. Login to Docker
```bash
docker login
```

Then run the build-push-image.sh script
```bash
chmod +x build-push-image.sh
./build-push-image.sh
```

`cd` to the root of the repo:
```bash
cd ../..
```

Update the value of backstage.image.tag in the backstage values-custom.yaml 
```bash
vim helm-charts/infra/backstage/values-custom.yaml
```

Save and push to repo
```bash
git add .
git commit -m "Updated backstage image tag"
git push
```

If you have a Minikube cluster running, delete it first with `minikube delete`.

Now run the deploy-in-minikube.sh script to get everything setup:
```bash
chmod +x deploy-in-minikube.sh
./deploy-in-minikube.sh
```
</br>

Now go to localhost:8080 on your browser and Voilá!

You should be able to access ArgoCD UI on localhost:8081 server to check everything is runnin fine.

Grafana will also be exposed on localhost:8082. The credentials are:
- user: admin
- password: automate-all-the-things

</br>
</br>

# CONCLUSION
That's it! This is your own Backstage implementation now. 

Feel free to add your own plugins, templates and whatever else you might think of. Customize it to fit your own needs.

For more DevOps and Platform Engineering goodness, check out my [Automate All The Things](https://github.com/tferrari92/automate-all-the-things) project.

Happy automating!




<!-- 
##### Info interesante:
https://backstage.spotify.com/learn/backstage-for-all/software-catalog/4-modeling/
https://backstage.spotify.com/learn/standing-up-backstage/putting-backstage-into-action/8-integration/
https://backstage.spotify.com/learn/onboarding-software-to-backstage/onboarding-software-to-backstage/5-register-component/

##### Info datallada sobre objetos de tipo template:
https://backstage.io/docs/features/software-catalog/descriptor-format#kind-template
##### Aqui las acciones q puede hacer el template:
http://localhost:3000/create/actions
##### Para acciones q no existen default:
https://backstage.io/docs/features/software-templates/writing-custom-actions/
##### A note on RepoUrlPicker
In the template.yaml file of the template we created, you must have noticed ui:field: RepoUrlPicker in the spec.parameters field. This is known as Scaffolder Field Extensions.

These field extensions are used in taking certain types of input from users like GitHub repository URL, teams registered in catalog for the owners field, etc. Such field extensions can also be customized for your own organization. See https://backstage.io/docs/features/software-templates/writing-custom-field-extensions/

##### Aca hay ejemplos de templates:
https://github.com/backstage/software-templates

##### Software Templates at Spotify
At Spotify, we have dozens of Software Templates. We divide them into several disciples like Backend, Frontend, Data pipelines, etc. Inside Spotify, we also have stakeholder groups for Web, Backend, Data, etc. separately. These Software Templates are hosted on our internal GitHub enterprise, maintained and reviewed by the concerned experts in the discipline.

The Technical Architecture Group (TAG) at Spotify is the body responsible for reducing fragmentation by deciding on the various Backend, Frontend, Data frameworks to be used inside Spotify. Hence, new Software Templates with completely new frameworks are carefully discussed and reviewed.

Our Software Templates are fundamental to the concept of Golden Paths at Spotify. The Golden Path is the opinionated and supported way to build something (for example, build a backend service, put up a website, create a data pipeline). The Golden Path Tutorial is a step-by-step instructions that walks you through this opinionated and supported path.

The blessed tools — those on the Golden Path — are visualized in the Explore section of Backstage. Read more https://engineering.atspotify.com/2020/08/how-we-use-golden-paths-to-solve-fragmentation-in-our-software-ecosystem/



Searching through App Metadata with Backstage Search
The Backstage Search feature allows you to integrate custom search engine providers. You can also use any of the three default search engines: Lunr, Postgres, or Elasticsearch. Lunr is the current search engine enabled on your Backstage app. However, the documentation does not recommend this setup for a production environment because this search engine may not perform indexing well enough when the volume of app metadata and documentation increases.
https://www.kosli.com/blog/implementing-backstage-2-using-the-core-features/

Optimizing Search Highlighting
For a better search highlighting experience, add these lines of config to app-config.yaml:
```yaml
search:
  pg:
    highlightOptions:
      useHighlight: true
      maxWord: 35 # Used to set the longest headlines to output. The default value is 35.
      minWord: 15 # Used to set the shortest headlines to output. The default value is 15.
      shortWord: 3 # Words of this length or less will be dropped at the start and end of a headline, unless they are query terms. The default value of three (3) eliminates common English articles.
      highlightAll: false # If true the whole document will be used as the headline, ignoring the preceding three parameters. The default is false.
      maxFragments: 0 # Maximum number of text fragments to display. The default value of zero selects a non-fragment-based headline generation method. A value greater than zero selects fragment-based headline generation (see the linked documentation above for more details).
      fragmentDelimiter: ' ... ' # Delimiter string used to concatenate fragments. Defaults to " ... ".
```
https://www.kosli.com/blog/implementing-backstage-2-using-the-core-features/ -->



<!-- VER PORQ EL RESOURCE REDIS NO APARECE BAJO OWNERSHIP DEL GRUPO REDIS
PORQ My-App Redis Subteam no muestra ownership de resource redis??? http://localhost:3000/catalog/default/group/my-app-redis-subteam
# BACKSTAGE
If the only change you've made is to the app-config.yaml (or other configuration files) and not to the application code itself, you don't necessarily need to run yarn build or yarn build:backend. The Docker image build process should copy the updated configuration files into the image.

AGREGARLE DESCRIPTION AL REPO DE GHUB
ARREGLAR DLO DE LOS TAGS EN LOS TEMPLATES DE CREAR SERVICIOS
AGREGAR DEPENDS ON EN TEMPLATE -->
