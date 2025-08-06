#!/bin/bash

export GOOS=linux
export GOARCH=amd64

go build -o ../../../bin/linux/auth_microservice ../../cmd/auth_microservice/main.go

if [ $? -eq 0 ]; then
  cp ../../config/auth_microservice.conf ../../../bin/linux/auth_microservice.conf
  echo "Build successful!"
else
  echo "Build failed!"
fi
