
## 1.Creating a Kubernetes Cluster..
Creating a Kubernetes cluster using Kind (Kubernetes IN Docker) allows us to run clusters in Docker containers for local development. You can quickly set up a single-node cluster with the command :-
```
kind create cluster --name localwec
```
<img width="842" alt="Screenshot 2024-10-12 at 3 18 41 PM" src="https://github.com/user-attachments/assets/c1c327c6-326f-4dde-b2b1-d14d4b024fc6">



 This provides us  an isolated environment for testing Kubernetes features without the overhead of a cloud provider.

 ## 2.Creating a Deployment..
After  creating a Kubernetes cluster with Kind, the next step is typically to deploy your applications. 

Here’s the deployment manifest file:

Applying the deployment to kind cluster:
```
kubectl apply -f nodeapp-deployment.yaml

```
This file starts a replicaset which in turn starts the specified number of pods.



Pods Running Successfully!!

## 3.Configmaps
ConfigMaps in Kubernetes are used to store configuration data in key-value pairs, allowing us to decouple environment-specific settings from your containerized applications. They are required to manage configuration changes without rebuilding container images, enabling easier updates and flexibility.

Here’s the Configmap file:

Applying the Configmap to kind cluster

## 4.Service 
Service in Kubernetes allows us to expose an applications externally.There are several Services offered by Kubernetes. But we are using Nodeport Service here.

Here’s the Service file:



