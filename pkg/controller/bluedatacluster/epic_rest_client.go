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

var Session string
var once sync.Once
var BaseUrl string = "http://controller-svc.bluedata-system.svc.cluster.local:8080"
var client http.Client

func (c *EpicRestClient) Once() {
  once.Do(func() {
      client := http.Client{Timeout: time.Minute * 2}
      creds := make(map[string]string)
      creds["name"] = "demo.user"
      creds["password"] = "admin123"
      jsonValue, _ := json.Marshal(creds)
      request, _ := http.NewRequest("POST", BaseUrl + "/api/v1/login", bytes.NewBuffer(jsonValue))
      response, err := client.Do(request)
      Session = response.Header.Get("Location")
      log.Println("REST Login Session %s (Err: %s)", response.Header.Get("Location"), err)
  })
}

func (c *EpicRestClient) CreateCluster(cluster *v1alpha1.BlueDataCluster) {
  c.Once()
  jsonValue, _ := json.Marshal(cluster.Spec)
  request, _ := http.NewRequest("POST",
                            BaseUrl + "/api/v2/cluster",
                            bytes.NewBuffer(jsonValue))
  request.Header.Add("X-BDS-SESSION", Session)
  request.Header.Add("Content-Type", "application/json")
  response, err := client.Do(request)
  log.Println("REST CreateCluster %s Response: %s Err: %s",
              cluster.Name, response, err)
}
