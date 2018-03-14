package apis

import (
	"bluedata-apiserver-extension/pkg/apis/bluedata"
	bluedatav1alpha1 "bluedata-apiserver-extension/pkg/apis/bluedata/v1alpha1"
	"github.com/kubernetes-sigs/kubebuilder/pkg/builders"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type MetaData struct{}

var APIMeta = MetaData{}

// GetAllApiBuilders returns all known APIGroupBuilders
// so they can be registered with the apiserver
func (MetaData) GetAllApiBuilders() []*builders.APIGroupBuilder {
	return []*builders.APIGroupBuilder{
		GetBluedataAPIBuilder(),
	}
}

// GetCRDs returns all the CRDs for known resource types
func (MetaData) GetCRDs() []v1beta1.CustomResourceDefinition {
	return []v1beta1.CustomResourceDefinition{
		bluedatav1alpha1.BlueDataClusterCRD,
	}
}

func (MetaData) GetRules() []rbacv1.PolicyRule {
	return []rbacv1.PolicyRule{
		{
			APIGroups: []string{"bluedata.k8s.bluedata.com"},
			Resources: []string{"*"},
			Verbs:     []string{"*"},
		},
	}
}

func (MetaData) GetGroupVersions() []schema.GroupVersion {
	return []schema.GroupVersion{
		{
			Group:   "bluedata.k8s.bluedata.com",
			Version: "v1alpha1",
		},
	}
}

var bluedataApiGroup = builders.NewApiGroupBuilder(
	"bluedata.k8s.bluedata.com",
	"bluedata-apiserver-extension/pkg/apis/bluedata").
	WithUnVersionedApi(bluedata.ApiVersion).
	WithVersionedApis(
		bluedatav1alpha1.ApiVersion,
	).
	WithRootScopedKinds()

func GetBluedataAPIBuilder() *builders.APIGroupBuilder {
	return bluedataApiGroup
}
