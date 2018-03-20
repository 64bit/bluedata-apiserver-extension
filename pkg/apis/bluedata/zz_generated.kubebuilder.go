package bluedata

import (
	"fmt"
	"github.com/kubernetes-sigs/kubebuilder/pkg/builders"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
)

var (
	InternalBlueDataCluster = builders.NewInternalResource(
		"bluedataclusters",
		"BlueDataCluster",
		func() runtime.Object { return &BlueDataCluster{} },
		func() runtime.Object { return &BlueDataClusterList{} },
	)
	InternalBlueDataClusterStatus = builders.NewInternalResourceStatus(
		"bluedataclusters",
		"BlueDataClusterStatus",
		func() runtime.Object { return &BlueDataCluster{} },
		func() runtime.Object { return &BlueDataClusterList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("bluedata.k8s.bluedata.com").WithKinds(
		InternalBlueDataCluster,
		InternalBlueDataClusterStatus,
	)

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BlueDataCluster struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   BlueDataClusterSpec
	Status BlueDataClusterStatus
}

type BlueDataClusterSpec struct {
	Label          BlueDataLabel
	Nodegroup      BlueDataNodegroup
	TwoPhaseDelete bool
}

type BlueDataClusterStatus struct {
}

type BlueDataNodegroup struct {
	CatalogEntryDistroId string
	RoleConfigs          []BlueDataNodegroupRoleConfig
}

type BlueDataLabel struct {
	Name        string
	Description string
}

type BlueDataNodegroupRoleConfig struct {
	RoleId    string
	NodeCount int
	Flavor    string
}

//
// BlueDataCluster Functions and Structs
//
// +k8s:deepcopy-gen=false
type BlueDataClusterStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type BlueDataClusterStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BlueDataClusterList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []BlueDataCluster
}

func (BlueDataCluster) NewStatus() interface{} {
	return BlueDataClusterStatus{}
}

func (pc *BlueDataCluster) GetStatus() interface{} {
	return pc.Status
}

func (pc *BlueDataCluster) SetStatus(s interface{}) {
	pc.Status = s.(BlueDataClusterStatus)
}

func (pc *BlueDataCluster) GetSpec() interface{} {
	return pc.Spec
}

func (pc *BlueDataCluster) SetSpec(s interface{}) {
	pc.Spec = s.(BlueDataClusterSpec)
}

func (pc *BlueDataCluster) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *BlueDataCluster) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc BlueDataCluster) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store BlueDataCluster.
// +k8s:deepcopy-gen=false
type BlueDataClusterRegistry interface {
	ListBlueDataClusters(ctx request.Context, options *internalversion.ListOptions) (*BlueDataClusterList, error)
	GetBlueDataCluster(ctx request.Context, id string, options *metav1.GetOptions) (*BlueDataCluster, error)
	CreateBlueDataCluster(ctx request.Context, id *BlueDataCluster) (*BlueDataCluster, error)
	UpdateBlueDataCluster(ctx request.Context, id *BlueDataCluster) (*BlueDataCluster, error)
	DeleteBlueDataCluster(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewBlueDataClusterRegistry(sp builders.StandardStorageProvider) BlueDataClusterRegistry {
	return &storageBlueDataCluster{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageBlueDataCluster struct {
	builders.StandardStorageProvider
}

func (s *storageBlueDataCluster) ListBlueDataClusters(ctx request.Context, options *internalversion.ListOptions) (*BlueDataClusterList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*BlueDataClusterList), err
}

func (s *storageBlueDataCluster) GetBlueDataCluster(ctx request.Context, id string, options *metav1.GetOptions) (*BlueDataCluster, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*BlueDataCluster), nil
}

func (s *storageBlueDataCluster) CreateBlueDataCluster(ctx request.Context, object *BlueDataCluster) (*BlueDataCluster, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, true)
	if err != nil {
		return nil, err
	}
	return obj.(*BlueDataCluster), nil
}

func (s *storageBlueDataCluster) UpdateBlueDataCluster(ctx request.Context, object *BlueDataCluster) (*BlueDataCluster, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*BlueDataCluster), nil
}

func (s *storageBlueDataCluster) DeleteBlueDataCluster(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}
