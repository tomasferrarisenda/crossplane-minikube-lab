#!/bin/bash

# Prompt the user for their GitHub token
read -p "Enter your AWS access key ID: " AWS_ACCESS_KEY_ID
read -p "Enter your AWS secret access key: " AWS_SECRET_ACCESS_KEY

# Start cluster. Extra beefy.
minikube start --cpus 4 --memory 4096

# Install ArgoCD
helm install argocd -n argocd helm-charts/infra/argo-cd --values helm-charts/infra/argo-cd/values-custom.yaml --dependency-update --create-namespace

# Get ArgoCD admin password
until kubectl -n argocd get secret argocd-initial-admin-secret &> /dev/null; do
  echo "Waiting for secret 'argocd-initial-admin-secret' to be available..."
  sleep 3
done
echo "#############################################################################"
echo "#############################################################################"
echo "#############################################################################"
echo " "
echo "ACCESS THE ARGOCD DASHBOARD:"
echo "Go to http://localhost:8080/"
echo " "
echo "user: admin"
echo "password: $(kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d)"
echo " "
echo "#############################################################################"
echo "#############################################################################"
echo "#############################################################################"

# Create all ArgoCD Applications
kubectl apply -n argocd -f argo-cd

# Then we create an application that will monitor the helm-charts/infra/argo-cd directory, the same we used to deploy ArgoCD, making ArgoCD self-managed. Any changes we apply in the helm/infra/argocd directory will be automatically applied.
# kubectl create -n argocd -f argo-cd/infra/argocd-application.yaml  
# kubectl create -n argocd -f argo-cd/infra/crossplane-application.yaml  
# kubectl create -n argocd -f argo-cd/infra/crossplane-composite-resource-definitions-application.yaml  
# kubectl create -n argocd -f argo-cd/infra/crossplane-compositions-application.yaml  
# # kubectl create -n argocd -f argo-cd/infra/crossplane-configurations-application.yaml  
# kubectl create -n argocd -f argo-cd/infra/crossplane-provider-configs-application.yaml  
# kubectl create -n argocd -f argo-cd/infra/crossplane-providers-application.yaml  
# kubectl create -n argocd -f argo-cd/infra/my-cluster-application.yaml  

# Finally, we create an application that will automatically deploy any ArgoCD Applications we specify in the argo-cd/applications directory (App of Apps pattern).
# kubectl create -n argocd -f argo-cd/self-manage/argocd-app-of-apps-application.yaml  

# We expose argocd on port 8080 in the background 
kubectl port-forward -n argocd service/argocd-server 8080:443 &

# Create a secret with AWS credentials
echo -e "[default]\naws_access_key_id = $AWS_ACCESS_KEY_ID\naws_secret_access_key = $AWS_SECRET_ACCESS_KEY" > aws-credentials.txt
kubectl create ns crossplane-system
kubectl create secret generic aws-secret -n crossplane-system --from-file=creds=./aws-credentials.txt
rm aws-credentials.txt

# Wait for backend database to be ready
while [[ $(kubectl get database.postgresql.sql.crossplane.io/my-app-backend-db -o jsonpath='{.status.conditions[?(@.type=="Ready")].status}') != "True" ]]; do
    kubectl get database.postgresql.sql.crossplane.io/my-app-backend-db
    sleep 5
done

# Port forward the frontend service
kubectl port-forward -n my-app service/my-app-frontend 8081:80

echo "#############################################################################"
echo "#############################################################################"
echo "#############################################################################"
echo " "
echo "ACCESS THE FRONTEND:"
echo "Go to http://localhost:8081/"
echo " "
echo "#############################################################################"
echo "#############################################################################"
echo "#############################################################################"