package fake

import (
	v1alpha1 "bluedata-apiserver-extension/pkg/apis/bluedata/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBlueDataClusters implements BlueDataClusterInterface
type FakeBlueDataClusters struct {
	Fake *FakeBluedataV1alpha1
	ns   string
}

var bluedataclustersResource = schema.GroupVersionResource{Group: "bluedata.k8s.bluedata.com", Version: "v1alpha1", Resource: "bluedataclusters"}

var bluedataclustersKind = schema.GroupVersionKind{Group: "bluedata.k8s.bluedata.com", Version: "v1alpha1", Kind: "BlueDataCluster"}

// Get takes name of the blueDataCluster, and returns the corresponding blueDataCluster object, and an error if there is any.
func (c *FakeBlueDataClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.BlueDataCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bluedataclustersResource, c.ns, name), &v1alpha1.BlueDataCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlueDataCluster), err
}

// List takes label and field selectors, and returns the list of BlueDataClusters that match those selectors.
func (c *FakeBlueDataClusters) List(opts v1.ListOptions) (result *v1alpha1.BlueDataClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bluedataclustersResource, bluedataclustersKind, c.ns, opts), &v1alpha1.BlueDataClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BlueDataClusterList{}
	for _, item := range obj.(*v1alpha1.BlueDataClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested blueDataClusters.
func (c *FakeBlueDataClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bluedataclustersResource, c.ns, opts))

}

// Create takes the representation of a blueDataCluster and creates it.  Returns the server's representation of the blueDataCluster, and an error, if there is any.
func (c *FakeBlueDataClusters) Create(blueDataCluster *v1alpha1.BlueDataCluster) (result *v1alpha1.BlueDataCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bluedataclustersResource, c.ns, blueDataCluster), &v1alpha1.BlueDataCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlueDataCluster), err
}

// Update takes the representation of a blueDataCluster and updates it. Returns the server's representation of the blueDataCluster, and an error, if there is any.
func (c *FakeBlueDataClusters) Update(blueDataCluster *v1alpha1.BlueDataCluster) (result *v1alpha1.BlueDataCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bluedataclustersResource, c.ns, blueDataCluster), &v1alpha1.BlueDataCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlueDataCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBlueDataClusters) UpdateStatus(blueDataCluster *v1alpha1.BlueDataCluster) (*v1alpha1.BlueDataCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bluedataclustersResource, "status", c.ns, blueDataCluster), &v1alpha1.BlueDataCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlueDataCluster), err
}

// Delete takes name of the blueDataCluster and deletes it. Returns an error if one occurs.
func (c *FakeBlueDataClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bluedataclustersResource, c.ns, name), &v1alpha1.BlueDataCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBlueDataClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bluedataclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BlueDataClusterList{})
	return err
}

// Patch applies the patch and returns the patched blueDataCluster.
func (c *FakeBlueDataClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BlueDataCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bluedataclustersResource, c.ns, name, data, subresources...), &v1alpha1.BlueDataCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlueDataCluster), err
}
