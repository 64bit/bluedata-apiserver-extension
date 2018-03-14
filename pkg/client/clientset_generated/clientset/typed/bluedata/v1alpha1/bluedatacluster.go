package v1alpha1

import (
	v1alpha1 "bluedata-apiserver-extension/pkg/apis/bluedata/v1alpha1"
	scheme "bluedata-apiserver-extension/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BlueDataClustersGetter has a method to return a BlueDataClusterInterface.
// A group's client should implement this interface.
type BlueDataClustersGetter interface {
	BlueDataClusters(namespace string) BlueDataClusterInterface
}

// BlueDataClusterInterface has methods to work with BlueDataCluster resources.
type BlueDataClusterInterface interface {
	Create(*v1alpha1.BlueDataCluster) (*v1alpha1.BlueDataCluster, error)
	Update(*v1alpha1.BlueDataCluster) (*v1alpha1.BlueDataCluster, error)
	UpdateStatus(*v1alpha1.BlueDataCluster) (*v1alpha1.BlueDataCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BlueDataCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.BlueDataClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BlueDataCluster, err error)
	BlueDataClusterExpansion
}

// blueDataClusters implements BlueDataClusterInterface
type blueDataClusters struct {
	client rest.Interface
	ns     string
}

// newBlueDataClusters returns a BlueDataClusters
func newBlueDataClusters(c *BluedataV1alpha1Client, namespace string) *blueDataClusters {
	return &blueDataClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the blueDataCluster, and returns the corresponding blueDataCluster object, and an error if there is any.
func (c *blueDataClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.BlueDataCluster, err error) {
	result = &v1alpha1.BlueDataCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bluedataclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BlueDataClusters that match those selectors.
func (c *blueDataClusters) List(opts v1.ListOptions) (result *v1alpha1.BlueDataClusterList, err error) {
	result = &v1alpha1.BlueDataClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bluedataclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested blueDataClusters.
func (c *blueDataClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bluedataclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a blueDataCluster and creates it.  Returns the server's representation of the blueDataCluster, and an error, if there is any.
func (c *blueDataClusters) Create(blueDataCluster *v1alpha1.BlueDataCluster) (result *v1alpha1.BlueDataCluster, err error) {
	result = &v1alpha1.BlueDataCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bluedataclusters").
		Body(blueDataCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a blueDataCluster and updates it. Returns the server's representation of the blueDataCluster, and an error, if there is any.
func (c *blueDataClusters) Update(blueDataCluster *v1alpha1.BlueDataCluster) (result *v1alpha1.BlueDataCluster, err error) {
	result = &v1alpha1.BlueDataCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bluedataclusters").
		Name(blueDataCluster.Name).
		Body(blueDataCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *blueDataClusters) UpdateStatus(blueDataCluster *v1alpha1.BlueDataCluster) (result *v1alpha1.BlueDataCluster, err error) {
	result = &v1alpha1.BlueDataCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bluedataclusters").
		Name(blueDataCluster.Name).
		SubResource("status").
		Body(blueDataCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the blueDataCluster and deletes it. Returns an error if one occurs.
func (c *blueDataClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bluedataclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *blueDataClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bluedataclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched blueDataCluster.
func (c *blueDataClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BlueDataCluster, err error) {
	result = &v1alpha1.BlueDataCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bluedataclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
