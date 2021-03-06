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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-ovh-api/apis/iploadbalancing/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TcpFarmServerLister helps list TcpFarmServers.
// All objects returned here must be treated as read-only.
type TcpFarmServerLister interface {
	// List lists all TcpFarmServers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TcpFarmServer, err error)
	// TcpFarmServers returns an object that can list and get TcpFarmServers.
	TcpFarmServers(namespace string) TcpFarmServerNamespaceLister
	TcpFarmServerListerExpansion
}

// tcpFarmServerLister implements the TcpFarmServerLister interface.
type tcpFarmServerLister struct {
	indexer cache.Indexer
}

// NewTcpFarmServerLister returns a new TcpFarmServerLister.
func NewTcpFarmServerLister(indexer cache.Indexer) TcpFarmServerLister {
	return &tcpFarmServerLister{indexer: indexer}
}

// List lists all TcpFarmServers in the indexer.
func (s *tcpFarmServerLister) List(selector labels.Selector) (ret []*v1alpha1.TcpFarmServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TcpFarmServer))
	})
	return ret, err
}

// TcpFarmServers returns an object that can list and get TcpFarmServers.
func (s *tcpFarmServerLister) TcpFarmServers(namespace string) TcpFarmServerNamespaceLister {
	return tcpFarmServerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TcpFarmServerNamespaceLister helps list and get TcpFarmServers.
// All objects returned here must be treated as read-only.
type TcpFarmServerNamespaceLister interface {
	// List lists all TcpFarmServers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TcpFarmServer, err error)
	// Get retrieves the TcpFarmServer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TcpFarmServer, error)
	TcpFarmServerNamespaceListerExpansion
}

// tcpFarmServerNamespaceLister implements the TcpFarmServerNamespaceLister
// interface.
type tcpFarmServerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TcpFarmServers in the indexer for a given namespace.
func (s tcpFarmServerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TcpFarmServer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TcpFarmServer))
	})
	return ret, err
}

// Get retrieves the TcpFarmServer from the indexer for a given namespace and name.
func (s tcpFarmServerNamespaceLister) Get(name string) (*v1alpha1.TcpFarmServer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tcpfarmserver"), name)
	}
	return obj.(*v1alpha1.TcpFarmServer), nil
}
