#!/bin/bash 

kubectl config set-context john-dev --cluster=kubernetes --user=john --namespace=dev
kubectl config set-context mary-prod --cluster=kubernetes --user=mary --namespace=prod
