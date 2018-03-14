package controller

import (
	"bluedata-apiserver-extension/pkg/controller/bluedatacluster"
	"bluedata-apiserver-extension/pkg/controller/sharedinformers"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"k8s.io/client-go/rest"
)

func GetAllControllers(config *rest.Config) ([]controller.Controller, chan struct{}) {
	shutdown := make(chan struct{})
	si := sharedinformers.NewSharedInformers(config, shutdown)
	return []controller.Controller{
		bluedatacluster.NewBlueDataClusterController(config, si),
	}, shutdown
}
