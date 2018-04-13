#!/bin/bash

# Required parameters
# CPU
# Memory
# Storage
# Namespace

CPU="1"
MEMORY="1Gi"
STORAGE="400Gi"
NAMESPACE="default"
TEMPLATE=epic-3.4-176-namespace-template.yaml

while getopts c:m:s:n: option; do 
    case $option in 
       c) 
          CPU=$OPTARG
          ;;
       m) 
          MEMORY=$OPTARG
          ;;
       s) 
          STORAGE=$OPTARG
          ;;
       n) 
          NAMESPACE=$OPTARG
          ;;
       \?)
          echo "Invalid Option"
          exit 1
    esac
done

echo "NAMESPACE: $NAMESPACE, CPU: $CPU, MEMORY: $MEMORY, STORAGE: $STORAGE"

FILE=epic-for-$NAMESPACE.yaml
cp $TEMPLATE $FILE
sed -i "s/{{{{CPU}}}}/$CPU/g ; s/{{{{MEMORY}}}}/$MEMORY/g; s/{{{{STORAGE}}}}/$STORAGE/g; s/{{{{NAMESPACE}}}}/$NAMESPACE/g" $FILE
echo "Generated $FILE"
