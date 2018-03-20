


package bluedatacluster

import (
    "log"
    "github.com/kubernetes-sigs/kubebuilder/pkg/builders"
    "bluedata-apiserver-extension/pkg/apis/bluedata/v1alpha1"
    "bluedata-apiserver-extension/pkg/controller/sharedinformers"
    "k8s.io/client-go/tools/cache"
    listers "bluedata-apiserver-extension/pkg/client/listers_generated/bluedata/v1alpha1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement controller logic for the BlueDataCluster resource API

// Reconcile handles enqueued messages
func (c *BlueDataClusterControllerImpl) Reconcile(u *v1alpha1.BlueDataCluster) error {
    // INSERT YOUR CODE HERE - implement controller logic to reconcile observed and desired state of the object
    log.Printf("Running RECONCILE BlueDataCluster for %s\n", u.Name)
    return nil
}

// +controller:group=bluedata,version=v1alpha1,kind=BlueDataCluster,resource=bluedataclusters
type BlueDataClusterControllerImpl struct {
    builders.DefaultControllerFns

    // lister indexes properties about BlueDataCluster
    lister listers.BlueDataClusterLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *BlueDataClusterControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
    // INSERT YOUR CODE HERE - add logic for initializing the controller as needed

    // Use the lister for indexing bluedataclusters labels
    c.lister = arguments.GetSharedInformers().Factory.Bluedata().V1alpha1().BlueDataClusters().Lister()

    // To watch other resource types, uncomment this function and replace Foo with the resource name to watch.
    // Must define the func FooToBlueDataCluster(i interface{}) (string, error) {} that returns the BlueDataCluster
    // "namespace/name"" to reconcile in response to the updated Foo
    // Note: To watch Kubernetes resources, you must also update the StartAdditionalInformers function in
    // pkg/controllers/sharedinformers/informers.go
    //
    // arguments.Watch("BlueDataClusterFoo",
    //     arguments.GetSharedInformers().Factory.Bar().V1beta1().Bars().Informer(),
    //     c.FooToBlueDataCluster)

    erc := EpicRestClient{}
    arguments.GetSharedInformers().Factory.Bluedata().V1alpha1().BlueDataClusters().Informer().
  		AddEventHandler(cache.ResourceEventHandlerFuncs{
             AddFunc: func(obj interface{}) {
                    bluedatacluster, _ := obj.(*v1alpha1.BlueDataCluster)
                    log.Printf("ADD BlueDataCluster: %s", bluedatacluster.Name)
                    erc.CreateCluster(bluedatacluster)
             },
             DeleteFunc: func(obj interface{}) {
                    bluedatacluster, _ := obj.(*v1alpha1.BlueDataCluster)
                    log.Printf("DELETE BlueDataCluster: %s ", bluedatacluster.Name)
             },
             UpdateFunc: func(oldObj, newObj interface{}) {
                    log.Printf("\n\nUPDATE \n\n     Old: %+v \n\n      New: %+v\n", oldObj, newObj)
             },
      })

}

func (c *BlueDataClusterControllerImpl) Get(namespace, name string) (*v1alpha1.BlueDataCluster, error) {
    return c.lister.BlueDataClusters(namespace).Get(name)
}
