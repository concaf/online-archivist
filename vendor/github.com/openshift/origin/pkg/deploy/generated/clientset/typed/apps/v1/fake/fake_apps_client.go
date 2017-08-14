package fake

import (
	v1 "github.com/openshift/origin/pkg/deploy/generated/clientset/typed/apps/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAppsV1 struct {
	*testing.Fake
}

func (c *FakeAppsV1) DeploymentConfigs(namespace string) v1.DeploymentConfigInterface {
	return &FakeDeploymentConfigs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAppsV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
