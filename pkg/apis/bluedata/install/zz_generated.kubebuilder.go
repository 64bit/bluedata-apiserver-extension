package install

import (
	"bluedata-apiserver-extension/pkg/apis"
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	"k8s.io/apimachinery/pkg/runtime"
)

func Install(
	groupFactoryRegistry announced.APIGroupFactoryRegistry,
	registry *registered.APIRegistrationManager,
	scheme *runtime.Scheme) {

	apis.GetBluedataAPIBuilder().Install(groupFactoryRegistry, registry, scheme)
}
