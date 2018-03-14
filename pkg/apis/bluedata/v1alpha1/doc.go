


// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=bluedata-apiserver-extension/pkg/apis/bluedata
// +k8s:defaulter-gen=TypeMeta
// +groupName=bluedata.k8s.bluedata.com
package v1alpha1 // import "bluedata-apiserver-extension/pkg/apis/bluedata/v1alpha1"
