/*
Copyright (c) 2016-2017 Bitnami

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package fake

import (
	v1beta1 "github.com/kubeless/kubeless/pkg/client/clientset/versioned/typed/kubeless/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKubelessV1beta1 struct {
	*testing.Fake
}

func (c *FakeKubelessV1beta1) CronJobTriggers(namespace string) v1beta1.CronJobTriggerInterface {
	return &FakeCronJobTriggers{c, namespace}
}

func (c *FakeKubelessV1beta1) Functions(namespace string) v1beta1.FunctionInterface {
	return &FakeFunctions{c, namespace}
}

func (c *FakeKubelessV1beta1) HTTPTriggers(namespace string) v1beta1.HTTPTriggerInterface {
	return &FakeHTTPTriggers{c, namespace}
}

func (c *FakeKubelessV1beta1) KafkaTriggers(namespace string) v1beta1.KafkaTriggerInterface {
	return &FakeKafkaTriggers{c, namespace}
}

func (c *FakeKubelessV1beta1) KinesisTriggers(namespace string) v1beta1.KinesisTriggerInterface {
	return &FakeKinesisTriggers{c, namespace}
}

func (c *FakeKubelessV1beta1) NATSTriggers(namespace string) v1beta1.NATSTriggerInterface {
	return &FakeNATSTriggers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKubelessV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
