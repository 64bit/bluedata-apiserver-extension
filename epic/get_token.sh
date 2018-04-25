#!/bin/bash 
export TOKEN=`kubectl describe secret -n bluedata-system $(kubectl get secrets -n bluedata-system | cut -f1 -d ' ' | tail -n1) | tail -n1 | awk '{print $2}'`
echo $TOKEN
