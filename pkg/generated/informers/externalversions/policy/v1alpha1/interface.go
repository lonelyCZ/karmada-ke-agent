// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/xmwilldo/karmada-ke-agent/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterOverridePolicies returns a ClusterOverridePolicyInformer.
	ClusterOverridePolicies() ClusterOverridePolicyInformer
	// ClusterPropagationPolicies returns a ClusterPropagationPolicyInformer.
	ClusterPropagationPolicies() ClusterPropagationPolicyInformer
	// OverridePolicies returns a OverridePolicyInformer.
	OverridePolicies() OverridePolicyInformer
	// PropagationPolicies returns a PropagationPolicyInformer.
	PropagationPolicies() PropagationPolicyInformer
	// ReplicaSchedulingPolicies returns a ReplicaSchedulingPolicyInformer.
	ReplicaSchedulingPolicies() ReplicaSchedulingPolicyInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterOverridePolicies returns a ClusterOverridePolicyInformer.
func (v *version) ClusterOverridePolicies() ClusterOverridePolicyInformer {
	return &clusterOverridePolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterPropagationPolicies returns a ClusterPropagationPolicyInformer.
func (v *version) ClusterPropagationPolicies() ClusterPropagationPolicyInformer {
	return &clusterPropagationPolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// OverridePolicies returns a OverridePolicyInformer.
func (v *version) OverridePolicies() OverridePolicyInformer {
	return &overridePolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PropagationPolicies returns a PropagationPolicyInformer.
func (v *version) PropagationPolicies() PropagationPolicyInformer {
	return &propagationPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ReplicaSchedulingPolicies returns a ReplicaSchedulingPolicyInformer.
func (v *version) ReplicaSchedulingPolicies() ReplicaSchedulingPolicyInformer {
	return &replicaSchedulingPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
