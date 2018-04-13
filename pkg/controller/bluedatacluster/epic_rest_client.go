package bluedatacluster

import (
  "log"
  "bytes"
  "encoding/json"
  "net/http"
  "time"
  "sync"
  "bluedata-apiserver-extension/pkg/apis/bluedata/v1alpha1"
)

type EpicRestClient struct {
}

var client http.Client
var credsJson []byte
var once sync.Once

func (c *EpicRestClient) Once() {
  once.Do(func() {
      client = http.Client{Timeout: time.Minute * 2}
      creds := make(map[string]string)
      creds["name"] = "demo.user"
      creds["password"] = "admin123"
      credsJson, _ = json.Marshal(creds)
  })
}

func (c *EpicRestClient) Login(namespace string) (string, string) {
    c.Once()
    BaseUrl :=  "http://controller-svc." + namespace + ".svc.cluster.local:8080"
    request, _ := http.NewRequest("POST", BaseUrl + "/api/v1/login",
                                  bytes.NewBuffer(credsJson))
    response, err := client.Do(request)
    if response == nil {
      return "", BaseUrl
    }
    Session := response.Header.Get("Location")
    log.Println("REST Login Session:", response.Header.Get("Location"),
                ". Error: ", err)
    return Session, BaseUrl
}

func (c *EpicRestClient) CreateCluster(cluster *v1alpha1.BlueDataCluster) {
    Session, BaseUrl := c.Login(cluster.ObjectMeta.Namespace)
    log.Println("BaseUrl: ", BaseUrl)
    if Session == "" {
        log.Println("Empty session, doing nothing..")
        return
    }
    jsonValue, _ := json.Marshal(cluster.Spec)
    request, _ := http.NewRequest("POST", BaseUrl + "/api/v2/cluster",
                                  bytes.NewBuffer(jsonValue))
    request.Header.Add("X-BDS-SESSION", Session)
    request.Header.Add("X-NAMESPACE", cluster.ObjectMeta.Namespace)
    request.Header.Add("Content-Type", "application/json")
    log.Println("Namespace: ", cluster.ObjectMeta.Namespace)
    response, err := client.Do(request)
    log.Println("REST CreateCluster:", cluster.Name,
                ". Response: ", response,
                ". Err: ", err)
}
