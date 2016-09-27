#!/bin/bash

apps=micro-*

if [ ! -z "$1" ]; then
    apps=$1
fi

for d in $apps
    do bash -c "echo $d; cd $d && GOOS=linux go build -o app && cp ../Dockerfile Dockerfile && docker build -t $d:1.0 . && rm Dockerfile app" 
done
