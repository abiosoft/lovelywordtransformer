#!/bin/bash

for d in micro-*
    do bash -c "echo $d; cd $d && GOOS=linux go build -o app && cp ../Dockerfile Dockerfile && docker build -t $d:1.0 . && rm Dockerfile app" 
done