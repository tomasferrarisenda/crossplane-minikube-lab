#!/bin/bash

export DB=my-app-backend-db

kubectl patch object.kubernetes.crossplane.io $DB \
    --patch '{"metadata":{"finalizers":[]}}' --type=merge

kubectl patch database.postgresql.sql.crossplane.io $DB \
    --patch '{"metadata":{"finalizers":[]}}' --type=merge
