// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/xmwilldo/karmada-ke-agent/pkg/generated/clientset/versioned/typed/policy/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakePolicyV1alpha1 struct {
	*testing.Fake
}

func (c *FakePolicyV1alpha1) ClusterOverridePolicies() v1alpha1.ClusterOverridePolicyInterface {
	return &FakeClusterOverridePolicies{c}
}

func (c *FakePolicyV1alpha1) ClusterPropagationPolicies() v1alpha1.ClusterPropagationPolicyInterface {
	return &FakeClusterPropagationPolicies{c}
}

func (c *FakePolicyV1alpha1) OverridePolicies(namespace string) v1alpha1.OverridePolicyInterface {
	return &FakeOverridePolicies{c, namespace}
}

func (c *FakePolicyV1alpha1) PropagationPolicies(namespace string) v1alpha1.PropagationPolicyInterface {
	return &FakePropagationPolicies{c, namespace}
}

func (c *FakePolicyV1alpha1) ReplicaSchedulingPolicies(namespace string) v1alpha1.ReplicaSchedulingPolicyInterface {
	return &FakeReplicaSchedulingPolicies{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePolicyV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
