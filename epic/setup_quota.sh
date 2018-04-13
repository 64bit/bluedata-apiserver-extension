#!/bin/bash

# Required parameters
# CPU
# Memory
# Storage
# Namespace


QUOTA_TEMPLATE=resource-quota-template.yaml

while getopts c:m:s:n: option; do
    case $option in
       c)
          CPU=$OPTARG
          ;;
       m)
          MEMORY=$(OPTARG *1024*1024*1024)
          ;;
       s)
          STORAGE=$(OPTARG *1024*1024*1024)
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

RESOURCE_QUOTA_FILE=resource-quota-for-$NAMESPACE.yaml
cp $QUOTA_TEMPLATE $RESOURCE_QUOTA_FILE
sed -i "s/{{{{CPU}}}}/$CPU/g ; s/{{{{MEMORY}}}}/$MEMORY/g; s/{{{{STORAGE}}}}/$STORAGE/g; s/{{{{NAMESPACE}}}}/$NAMESPACE/g" $RESOURCE_QUOTA_FILE
echo "Generated $RESOURCE_QUOTA_FILE"
