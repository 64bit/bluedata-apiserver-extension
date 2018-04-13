#!/bin/bash

# Required parameters
# CPU
# Memory
# Storage
# Namespace

CPU="40"
MEMORY=$((40*1024*1024*1024))
STORAGE=$((500*1024*1024*1024))
NAMESPACE="default"
TEMPLATE=epic-3.4-178-namespace-template.yaml
QUOTA_TEMPLATE=resource-quota-template.yaml

while getopts c:m:s:n: option; do
    case $option in
       c)
          CPU=$OPTARG
          ;;
       m)
          MEMORY=$((OPTARG *1024*1024*1024))
          ;;
       s)
          STORAGE=$((OPTARG *1024*1024*1024))
          ;;
       n)
          NAMESPACE=$OPTARG
          ;;
       \?)
          echo "Invalid Option"
          exit 1
    esac
done

echo "NAMESPACE: $NAMESPACE, CPU: $((CPU/(1024*1024*1024))) G, MEMORY: $((MEMORY/(1024*1024*1024))) G, STORAGE: $((STORAGE/(1024*1024*1024))) G"

FILE=epic-for-$NAMESPACE.yaml
cp $TEMPLATE $FILE
sed -i "s/{{{{CPU}}}}/$CPU/g ; s/{{{{MEMORY}}}}/$MEMORY/g; s/{{{{STORAGE}}}}/$STORAGE/g; s/{{{{NAMESPACE}}}}/$NAMESPACE/g" $FILE
echo "Generated $FILE"
