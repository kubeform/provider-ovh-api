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
	v1alpha1 "kubeform.dev/provider-ovh-api/apis/me/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IpxeScriptLister helps list IpxeScripts.
// All objects returned here must be treated as read-only.
type IpxeScriptLister interface {
	// List lists all IpxeScripts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IpxeScript, err error)
	// IpxeScripts returns an object that can list and get IpxeScripts.
	IpxeScripts(namespace string) IpxeScriptNamespaceLister
	IpxeScriptListerExpansion
}

// ipxeScriptLister implements the IpxeScriptLister interface.
type ipxeScriptLister struct {
	indexer cache.Indexer
}

// NewIpxeScriptLister returns a new IpxeScriptLister.
func NewIpxeScriptLister(indexer cache.Indexer) IpxeScriptLister {
	return &ipxeScriptLister{indexer: indexer}
}

// List lists all IpxeScripts in the indexer.
func (s *ipxeScriptLister) List(selector labels.Selector) (ret []*v1alpha1.IpxeScript, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IpxeScript))
	})
	return ret, err
}

// IpxeScripts returns an object that can list and get IpxeScripts.
func (s *ipxeScriptLister) IpxeScripts(namespace string) IpxeScriptNamespaceLister {
	return ipxeScriptNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IpxeScriptNamespaceLister helps list and get IpxeScripts.
// All objects returned here must be treated as read-only.
type IpxeScriptNamespaceLister interface {
	// List lists all IpxeScripts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IpxeScript, err error)
	// Get retrieves the IpxeScript from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IpxeScript, error)
	IpxeScriptNamespaceListerExpansion
}

// ipxeScriptNamespaceLister implements the IpxeScriptNamespaceLister
// interface.
type ipxeScriptNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IpxeScripts in the indexer for a given namespace.
func (s ipxeScriptNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IpxeScript, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IpxeScript))
	})
	return ret, err
}

// Get retrieves the IpxeScript from the indexer for a given namespace and name.
func (s ipxeScriptNamespaceLister) Get(name string) (*v1alpha1.IpxeScript, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ipxescript"), name)
	}
	return obj.(*v1alpha1.IpxeScript), nil
}
