package fake

import (
	v1alpha1 "bluedata-apiserver-extension/pkg/client/clientset_generated/clientset/typed/bluedata/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeBluedataV1alpha1 struct {
	*testing.Fake
}

func (c *FakeBluedataV1alpha1) BlueDataClusters(namespace string) v1alpha1.BlueDataClusterInterface {
	return &FakeBlueDataClusters{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBluedataV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
