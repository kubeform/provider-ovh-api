/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	domainv1alpha1 "kubeform.dev/provider-ovh-api/apis/domain/v1alpha1"
	versioned "kubeform.dev/provider-ovh-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-ovh-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-ovh-api/client/listers/domain/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ZoneRecordInformer provides access to a shared informer and lister for
// ZoneRecords.
type ZoneRecordInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ZoneRecordLister
}

type zoneRecordInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewZoneRecordInformer constructs a new informer for ZoneRecord type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewZoneRecordInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredZoneRecordInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredZoneRecordInformer constructs a new informer for ZoneRecord type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredZoneRecordInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DomainV1alpha1().ZoneRecords(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DomainV1alpha1().ZoneRecords(namespace).Watch(context.TODO(), options)
			},
		},
		&domainv1alpha1.ZoneRecord{},
		resyncPeriod,
		indexers,
	)
}

func (f *zoneRecordInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredZoneRecordInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *zoneRecordInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&domainv1alpha1.ZoneRecord{}, f.defaultInformer)
}

func (f *zoneRecordInformer) Lister() v1alpha1.ZoneRecordLister {
	return v1alpha1.NewZoneRecordLister(f.Informer().GetIndexer())
}
