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

// TcpRouteRuleLister helps list TcpRouteRules.
// All objects returned here must be treated as read-only.
type TcpRouteRuleLister interface {
	// List lists all TcpRouteRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TcpRouteRule, err error)
	// TcpRouteRules returns an object that can list and get TcpRouteRules.
	TcpRouteRules(namespace string) TcpRouteRuleNamespaceLister
	TcpRouteRuleListerExpansion
}

// tcpRouteRuleLister implements the TcpRouteRuleLister interface.
type tcpRouteRuleLister struct {
	indexer cache.Indexer
}

// NewTcpRouteRuleLister returns a new TcpRouteRuleLister.
func NewTcpRouteRuleLister(indexer cache.Indexer) TcpRouteRuleLister {
	return &tcpRouteRuleLister{indexer: indexer}
}

// List lists all TcpRouteRules in the indexer.
func (s *tcpRouteRuleLister) List(selector labels.Selector) (ret []*v1alpha1.TcpRouteRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TcpRouteRule))
	})
	return ret, err
}

// TcpRouteRules returns an object that can list and get TcpRouteRules.
func (s *tcpRouteRuleLister) TcpRouteRules(namespace string) TcpRouteRuleNamespaceLister {
	return tcpRouteRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TcpRouteRuleNamespaceLister helps list and get TcpRouteRules.
// All objects returned here must be treated as read-only.
type TcpRouteRuleNamespaceLister interface {
	// List lists all TcpRouteRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TcpRouteRule, err error)
	// Get retrieves the TcpRouteRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TcpRouteRule, error)
	TcpRouteRuleNamespaceListerExpansion
}

// tcpRouteRuleNamespaceLister implements the TcpRouteRuleNamespaceLister
// interface.
type tcpRouteRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TcpRouteRules in the indexer for a given namespace.
func (s tcpRouteRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TcpRouteRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TcpRouteRule))
	})
	return ret, err
}

// Get retrieves the TcpRouteRule from the indexer for a given namespace and name.
func (s tcpRouteRuleNamespaceLister) Get(name string) (*v1alpha1.TcpRouteRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tcprouterule"), name)
	}
	return obj.(*v1alpha1.TcpRouteRule), nil
}