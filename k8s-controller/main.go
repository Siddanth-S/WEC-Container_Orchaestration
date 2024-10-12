package main

import (
    "log"
    "time"

    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/tools/cache"
    "k8s.io/client-go/informers"
    corev1 "k8s.io/api/core/v1"
)

func main() {
    // Load the kubeconfig
    kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
        clientcmd.NewDefaultClientConfigLoadingRules(),
        &clientcmd.ConfigOverrides{},
    )
    config, err := kubeconfig.ClientConfig()
    if err != nil {
        log.Fatalf("Error creating Kubernetes client config: %v", err)
    }

    // Create a Kubernetes clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        log.Fatalf("Error creating Kubernetes clientset: %v", err)
    }

    // Create a shared informer for ConfigMaps
    factory := informers.NewSharedInformerFactory(clientset, time.Minute)
    configMapInformer := factory.Core().V1().ConfigMaps().Informer()

    // Add event handlers for ConfigMap updates
    configMapInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
        UpdateFunc: func(oldObj, newObj interface{}) {
            newConfigMap := newObj.(*corev1.ConfigMap)

            // Log ConfigMap update
            log.Printf("ConfigMap %s updated at %v", newConfigMap.Name, time.Now())
        },
    })

    // Start the informer
    stopCh := make(chan struct{})
    defer close(stopCh)
    factory.Start(stopCh)

    // Wait forever
    select {}
}