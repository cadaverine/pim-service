#!/bin/bash

for i in $(ls build/k8s/)
    do kubectl apply -f build/k8s/$i
done