# client-go
* MY First k8s controller which helps you to expose nginx deployment cluster wide!

## How to use?
```bash
# 1. Go build
go build -o ekspose

# Build and push 
podman/docker build -t <user-name>/ekspose:v0.0.1

podman/docker push <user-name>/ekspose:v0.0.1

sed -i "s|IMAGE_CHANGEME|<user-name>/ekspose:v0.0.1|" manifests/deployment.yaml

or use my image (quay.io/praveen4g0/ekspose:v0.0.1) by default

oc apply -f manifests/
```
* will help you run controller on any k8s cluster, so now when you create depolyment `nginx` controller will expose service and ingress resource, and also deletes it when user deletes deployment!