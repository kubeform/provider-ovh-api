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

// IploadbalancingLister helps list Iploadbalancings.
// All objects returned here must be treated as read-only.
type IploadbalancingLister interface {
	// List lists all Iploadbalancings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Iploadbalancing, err error)
	// Iploadbalancings returns an object that can list and get Iploadbalancings.
	Iploadbalancings(namespace string) IploadbalancingNamespaceLister
	IploadbalancingListerExpansion
}

// iploadbalancingLister implements the IploadbalancingLister interface.
type iploadbalancingLister struct {
	indexer cache.Indexer
}

// NewIploadbalancingLister returns a new IploadbalancingLister.
func NewIploadbalancingLister(indexer cache.Indexer) IploadbalancingLister {
	return &iploadbalancingLister{indexer: indexer}
}

// List lists all Iploadbalancings in the indexer.
func (s *iploadbalancingLister) List(selector labels.Selector) (ret []*v1alpha1.Iploadbalancing, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Iploadbalancing))
	})
	return ret, err
}

// Iploadbalancings returns an object that can list and get Iploadbalancings.
func (s *iploadbalancingLister) Iploadbalancings(namespace string) IploadbalancingNamespaceLister {
	return iploadbalancingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IploadbalancingNamespaceLister helps list and get Iploadbalancings.
// All objects returned here must be treated as read-only.
type IploadbalancingNamespaceLister interface {
	// List lists all Iploadbalancings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Iploadbalancing, err error)
	// Get retrieves the Iploadbalancing from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Iploadbalancing, error)
	IploadbalancingNamespaceListerExpansion
}

// iploadbalancingNamespaceLister implements the IploadbalancingNamespaceLister
// interface.
type iploadbalancingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Iploadbalancings in the indexer for a given namespace.
func (s iploadbalancingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Iploadbalancing, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Iploadbalancing))
	})
	return ret, err
}

// Get retrieves the Iploadbalancing from the indexer for a given namespace and name.
func (s iploadbalancingNamespaceLister) Get(name string) (*v1alpha1.Iploadbalancing, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iploadbalancing"), name)
	}
	return obj.(*v1alpha1.Iploadbalancing), nil
}
