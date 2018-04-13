#!/bin/bash

# Required parameters
# CPU
# Memory
# Storage
# Namespace

CPU="10"
MEMORY=$((10*1024*1024*1024))
STORAGE=$((400*1024*1024*1024))
NAMESPACE="default"
TEMPLATE=epic-3.4-178-namespace-template.yaml

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

FILE=epic-for-$NAMESPACE.yaml
cp $TEMPLATE $FILE
sed -i "s/{{{{CPU}}}}/$CPU/g ; s/{{{{MEMORY}}}}/$MEMORY/g; s/{{{{STORAGE}}}}/$STORAGE/g; s/{{{{NAMESPACE}}}}/$NAMESPACE/g" $FILE
echo "Generated $FILE"
