## Inspired by following example:
https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/

## Prerequisites:
1. Installed nginx ingress controller 
(for minikube: minikube addons enable ingress )

## Instruction
1. Create Deployment 1 ```make dep1```
2. Expose Deployment 1```make exposeDep1```
3. Create Deployment 2 ```make dep2```
4. Expose Deployment 2```make exposeDep2```
5. Create Ingress ```make ingress```
6. Edit /etc/hosts file with following line
``` <minikube ipv4> hello-world.info```
6. verify routing using curl
```curl hello-world.info && curl hello-world.info/v2```