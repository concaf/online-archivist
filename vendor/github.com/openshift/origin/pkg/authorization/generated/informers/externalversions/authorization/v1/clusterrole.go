// This file was automatically generated by informer-gen

package v1

import (
	authorization_v1 "github.com/openshift/origin/pkg/authorization/apis/authorization/v1"
	clientset "github.com/openshift/origin/pkg/authorization/generated/clientset"
	internalinterfaces "github.com/openshift/origin/pkg/authorization/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/origin/pkg/authorization/generated/listers/authorization/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterRoleInformer provides access to a shared informer and lister for
// ClusterRoles.
type ClusterRoleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterRoleLister
}

type clusterRoleInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newClusterRoleInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.AuthorizationV1().ClusterRoles().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.AuthorizationV1().ClusterRoles().Watch(options)
			},
		},
		&authorization_v1.ClusterRole{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *clusterRoleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&authorization_v1.ClusterRole{}, newClusterRoleInformer)
}

func (f *clusterRoleInformer) Lister() v1.ClusterRoleLister {
	return v1.NewClusterRoleLister(f.Informer().GetIndexer())
}
