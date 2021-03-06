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
	v1alpha1 "kubeform.dev/provider-ovh-api/apis/dedicated/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CephACLLister helps list CephACLs.
// All objects returned here must be treated as read-only.
type CephACLLister interface {
	// List lists all CephACLs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CephACL, err error)
	// CephACLs returns an object that can list and get CephACLs.
	CephACLs(namespace string) CephACLNamespaceLister
	CephACLListerExpansion
}

// cephACLLister implements the CephACLLister interface.
type cephACLLister struct {
	indexer cache.Indexer
}

// NewCephACLLister returns a new CephACLLister.
func NewCephACLLister(indexer cache.Indexer) CephACLLister {
	return &cephACLLister{indexer: indexer}
}

// List lists all CephACLs in the indexer.
func (s *cephACLLister) List(selector labels.Selector) (ret []*v1alpha1.CephACL, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CephACL))
	})
	return ret, err
}

// CephACLs returns an object that can list and get CephACLs.
func (s *cephACLLister) CephACLs(namespace string) CephACLNamespaceLister {
	return cephACLNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CephACLNamespaceLister helps list and get CephACLs.
// All objects returned here must be treated as read-only.
type CephACLNamespaceLister interface {
	// List lists all CephACLs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CephACL, err error)
	// Get retrieves the CephACL from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CephACL, error)
	CephACLNamespaceListerExpansion
}

// cephACLNamespaceLister implements the CephACLNamespaceLister
// interface.
type cephACLNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CephACLs in the indexer for a given namespace.
func (s cephACLNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CephACL, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CephACL))
	})
	return ret, err
}

// Get retrieves the CephACL from the indexer for a given namespace and name.
func (s cephACLNamespaceLister) Get(name string) (*v1alpha1.CephACL, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cephacl"), name)
	}
	return obj.(*v1alpha1.CephACL), nil
}
