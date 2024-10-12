
## 1.Creating a Kubernetes Cluster..
Creating a Kubernetes cluster using Kind (Kubernetes IN Docker) allows us to run clusters in Docker containers for local development. You can quickly set up a single-node cluster with the command :-
```
kind create cluster --name localwec
```
<img width="842" alt="Screenshot 2024-10-12 at 3 18 41 PM" src="https://github.com/user-attachments/assets/c1c327c6-326f-4dde-b2b1-d14d4b024fc6">



 This provides us  an isolated environment for testing Kubernetes features without the overhead of a cloud provider.

 ## 2.Dockerizing the application..
Dockerizing an application involves creating a Docker image that contains the application’s code, dependencies, and environment. This allows the application to run consistently across different environments.Docker image is created from a [Dockerfile](https://github.com/Siddanth-S/wecgtest-new/blob/main/dockerfile)

*Build successful!!*

<img width="1109" alt="Screenshot 2024-10-12 at 3 24 45 PM" src="https://github.com/user-attachments/assets/93ac4f21-cfb7-42ba-810d-eb592f41cf7e">

<img width="902" alt="Screenshot 2024-10-12 at 3 25 45 PM" src="https://github.com/user-attachments/assets/54cdcd13-bf37-465b-9a47-094d7a53fdfe">


 ## 3.Creating a Deployment..
After  creating a Kubernetes cluster with Kind, the next step is typically to deploy your applications. 

Here’s the deployment manifest file:- [Deployment manifest file](https://github.com/Siddanth-S/wecgtest-new/blob/main/nodeapp-deployment.yml)

Applying the deployment to kind cluster:-
```
kubectl apply -f nodeapp-deployment.yaml

```
This file starts a replicaset which in turn starts the specified number of pods.

<img width="761" alt="Screenshot 2024-10-12 at 3 33 41 PM" src="https://github.com/user-attachments/assets/c25ec8d5-118d-4308-b0eb-56799e188306">


*Pods Running Successfully!!*

## 4.Configmaps
ConfigMaps in Kubernetes are used to store configuration data in key-value pairs, allowing us to decouple environment-specific settings from your containerized applications. They are required to manage configuration changes without rebuilding container images, enabling easier updates and flexibility.

Here’s the Configmap file:- [Configmaps manifest file](https://github.com/Siddanth-S/wecgtest-new/blob/main/nodeapp-configmap.yml)

Applying the Configmap to kind cluster:-
```
kubectl apply -f nodeapp-configmap.yaml

```

## 5.Service 
Service in Kubernetes allows us to expose an applications externally.There are several Services offered by Kubernetes. But we are using Nodeport Service here.

Here’s the Service file: [Service manifest file](https://github.com/Siddanth-S/wecgtest-new/edit/main/nodeapp-service.yml)

Applying the Service to the kind cluster:-

```
kubectl apply -f nodeapp-service.yaml

```

We can access the service using Nodeip 
So if the NodeIP is 192.168.99.100 and the NodePort is 3000, you can access the service at:

```
http://192.168.99.100:3000

```
<img width="600" alt="Screenshot 2024-10-12 at 3 43 01 PM" src="https://github.com/user-attachments/assets/046eb6e4-2843-43bf-8c16-a3594112332f">




