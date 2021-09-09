// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	workv1alpha1 "github.com/xmwilldo/karmada-ke-agent/pkg/apis/work/v1alpha1"
	versioned "github.com/xmwilldo/karmada-ke-agent/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/xmwilldo/karmada-ke-agent/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/xmwilldo/karmada-ke-agent/pkg/generated/listers/work/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ResourceBindingInformer provides access to a shared informer and lister for
// ResourceBindings.
type ResourceBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ResourceBindingLister
}

type resourceBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResourceBindingInformer constructs a new informer for ResourceBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResourceBindingInformer constructs a new informer for ResourceBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorkV1alpha1().ResourceBindings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorkV1alpha1().ResourceBindings(namespace).Watch(context.TODO(), options)
			},
		},
		&workv1alpha1.ResourceBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&workv1alpha1.ResourceBinding{}, f.defaultInformer)
}

func (f *resourceBindingInformer) Lister() v1alpha1.ResourceBindingLister {
	return v1alpha1.NewResourceBindingLister(f.Informer().GetIndexer())
}
