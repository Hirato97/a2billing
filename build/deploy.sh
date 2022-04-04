#!/bin/sh
echo "Copy environment file"
yes | cp -rf build/a2billing-go-api-env /root/go/env/a2billing-go-api-env
echo "Build go application"
go mod tidy
go get .
GOOS=linux GOARCH=amd64 go build -o a2billing-go-api main.go
echo "Restart service"
systemctl restart a2billing-go-api
