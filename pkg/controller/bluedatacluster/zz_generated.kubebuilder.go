package bluedatacluster

import (
	"bluedata-apiserver-extension/pkg/controller/sharedinformers"
	"github.com/golang/glog"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

// BlueDataClusterController implements the controller.BlueDataClusterController interface
type BlueDataClusterController struct {
	queue *controller.QueueWorker

	// Handles messages
	controller *BlueDataClusterControllerImpl

	Name string

	BeforeReconcile func(key string)
	AfterReconcile  func(key string, err error)

	Informers *sharedinformers.SharedInformers
}

// NewController returns a new BlueDataClusterController for responding to BlueDataCluster events
func NewBlueDataClusterController(config *rest.Config, si *sharedinformers.SharedInformers) *BlueDataClusterController {
	q := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "BlueDataCluster")

	queue := &controller.QueueWorker{q, 10, "BlueDataCluster", nil}
	c := &BlueDataClusterController{queue, nil, "BlueDataCluster", nil, nil, si}

	// For non-generated code to add events
	uc := &BlueDataClusterControllerImpl{}
	var ci sharedinformers.Controller = uc

	if i, ok := ci.(sharedinformers.ControllerInit); ok {
		i.Init(&sharedinformers.ControllerInitArgumentsImpl{si, config, c.LookupAndReconcile})
	}

	c.controller = uc

	queue.Reconcile = c.LookupAndReconcile
	if c.Informers.WorkerQueues == nil {
		c.Informers.WorkerQueues = map[string]*controller.QueueWorker{}
	}
	c.Informers.WorkerQueues["BlueDataCluster"] = queue
	si.Factory.Bluedata().V1alpha1().BlueDataClusters().Informer().
		AddEventHandler(&controller.QueueingEventHandler{q, nil, false})
	return c
}

func (c *BlueDataClusterController) GetName() string {
	return c.Name
}

func (c *BlueDataClusterController) LookupAndReconcile(key string) (err error) {
	var namespace, name string

	if c.BeforeReconcile != nil {
		c.BeforeReconcile(key)
	}
	if c.AfterReconcile != nil {
		// Wrap in a function so err is evaluated after it is set
		defer func() { c.AfterReconcile(key, err) }()
	}

	namespace, name, err = cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return
	}

	u, err := c.controller.Get(namespace, name)
	if errors.IsNotFound(err) {
		glog.Infof("Not doing work for BlueDataCluster %v because it has been deleted", key)
		// Set error so it is picked up by AfterReconcile and the return function
		err = nil
		return
	}
	if err != nil {
		glog.Errorf("Unable to retrieve BlueDataCluster %v from store: %v", key, err)
		return
	}

	// Set error so it is picked up by AfterReconcile and the return function
	err = c.controller.Reconcile(u)

	return
}

func (c *BlueDataClusterController) Run(stopCh <-chan struct{}) {
	for _, q := range c.Informers.WorkerQueues {
		q.Run(stopCh)
	}
	controller.GetDefaults(c.controller).Run(stopCh)
}
