#!/bin/bash

kubectl patch object.kubernetes.crossplane.io my-app-backend-db \
    --patch '{"metadata":{"finalizers":[]}}' --type=merge

kubectl patch database.postgresql.sql.crossplane.io my-app-backend-db \
    --patch '{"metadata":{"finalizers":[]}}' --type=merge

kubectl delete -n argocd applications.argoproj.io my-app-backend
kubectl delete -n argocd applications.argoproj.io my-cluster

minikube delete