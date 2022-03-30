/*
RabbitMQ Messaging Topology Kubernetes Operator
Copyright 2021 VMware, Inc.

This product is licensed to you under the Mozilla Public License 2.0 license (the "License").  You may not use this product except in compliance with the Mozilla 2.0 License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/rabbitmq/messaging-topology-operator/pkg/generated/clientset/versioned/typed/rabbitmq.com/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeRabbitmqV1alpha1 struct {
	*testing.Fake
}

func (c *FakeRabbitmqV1alpha1) SuperStreams(namespace string) v1alpha1.SuperStreamInterface {
	return &FakeSuperStreams{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeRabbitmqV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}