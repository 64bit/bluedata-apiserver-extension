


package bluedatacluster_test

import (
    . "bluedata-apiserver-extension/pkg/apis/bluedata/v1alpha1"
    . "bluedata-apiserver-extension/pkg/client/clientset_generated/clientset/typed/bluedata/v1alpha1"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement controller logic tests

var _ = Describe("BlueDataCluster controller", func() {
    var instance BlueDataCluster
    var expectedKey string
    var client BlueDataClusterInterface

    BeforeEach(func() {
        instance = BlueDataCluster{}
        instance.Name = "instance-1"
        expectedKey = "default/instance-1"
    })

    AfterEach(func() {
        client.Delete(instance.Name, &metav1.DeleteOptions{})
    })

    Describe("when creating a new object", func() {
        It("invoke the reconcile method", func() {
            after := make(chan struct{})
            controller.AfterReconcile = func(key string, err error) {
                defer func() {
                    // Recover in case the key is reconciled multiple times
                    defer func() { recover() }()
                    close(after)
                }()
                Expect(key).To(Equal(expectedKey))
                Expect(err).ToNot(HaveOccurred())
            }

            // Create the instance
            client = cs.BluedataV1alpha1().BlueDataClusters("default")
            _, err := client.Create(&instance)
            Expect(err).ShouldNot(HaveOccurred())

            // Wait for reconcile to happen
            Eventually(after).Should(BeClosed())

            // INSERT YOUR CODE HERE - test conditions post reconcile
        })
    })
})
