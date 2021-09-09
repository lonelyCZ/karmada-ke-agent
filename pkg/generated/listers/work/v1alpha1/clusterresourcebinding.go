// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/xmwilldo/karmada-ke-agent/pkg/apis/work/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterResourceBindingLister helps list ClusterResourceBindings.
// All objects returned here must be treated as read-only.
type ClusterResourceBindingLister interface {
	// List lists all ClusterResourceBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterResourceBinding, err error)
	// Get retrieves the ClusterResourceBinding from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterResourceBinding, error)
	ClusterResourceBindingListerExpansion
}

// clusterResourceBindingLister implements the ClusterResourceBindingLister interface.
type clusterResourceBindingLister struct {
	indexer cache.Indexer
}

// NewClusterResourceBindingLister returns a new ClusterResourceBindingLister.
func NewClusterResourceBindingLister(indexer cache.Indexer) ClusterResourceBindingLister {
	return &clusterResourceBindingLister{indexer: indexer}
}

// List lists all ClusterResourceBindings in the indexer.
func (s *clusterResourceBindingLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterResourceBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterResourceBinding))
	})
	return ret, err
}

// Get retrieves the ClusterResourceBinding from the index for a given name.
func (s *clusterResourceBindingLister) Get(name string) (*v1alpha1.ClusterResourceBinding, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterresourcebinding"), name)
	}
	return obj.(*v1alpha1.ClusterResourceBinding), nil
}
