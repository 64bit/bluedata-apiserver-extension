


package bluedatacluster_test

import (
    "testing"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "k8s.io/client-go/rest"
    "github.com/kubernetes-sigs/kubebuilder/pkg/test"

    "bluedata-apiserver-extension/pkg/apis"
    "bluedata-apiserver-extension/pkg/client/clientset_generated/clientset"
    "bluedata-apiserver-extension/pkg/controller/sharedinformers"
    "bluedata-apiserver-extension/pkg/controller/bluedatacluster"
)

var testenv *test.TestEnvironment
var config *rest.Config
var cs *clientset.Clientset
var shutdown chan struct{}
var controller *bluedatacluster.BlueDataClusterController
var si *sharedinformers.SharedInformers

func TestBlueDataCluster(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecsWithDefaultAndCustomReporters(t, "BlueDataCluster Suite", []Reporter{test.NewlineReporter{}})
}

var _ = BeforeSuite(func() {
    testenv = &test.TestEnvironment{CRDs: apis.APIMeta.GetCRDs()}
    var err error
    config, err = testenv.Start()
    Expect(err).NotTo(HaveOccurred())
    cs = clientset.NewForConfigOrDie(config)

    shutdown = make(chan struct{})
    si = sharedinformers.NewSharedInformers(config, shutdown)
    controller = bluedatacluster.NewBlueDataClusterController(config, si)
    controller.Run(shutdown)
})

var _ = AfterSuite(func() {
    testenv.Stop()
})
