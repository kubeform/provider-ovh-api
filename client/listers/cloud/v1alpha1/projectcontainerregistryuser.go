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
	v1alpha1 "kubeform.dev/provider-ovh-api/apis/cloud/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProjectContainerregistryUserLister helps list ProjectContainerregistryUsers.
// All objects returned here must be treated as read-only.
type ProjectContainerregistryUserLister interface {
	// List lists all ProjectContainerregistryUsers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProjectContainerregistryUser, err error)
	// ProjectContainerregistryUsers returns an object that can list and get ProjectContainerregistryUsers.
	ProjectContainerregistryUsers(namespace string) ProjectContainerregistryUserNamespaceLister
	ProjectContainerregistryUserListerExpansion
}

// projectContainerregistryUserLister implements the ProjectContainerregistryUserLister interface.
type projectContainerregistryUserLister struct {
	indexer cache.Indexer
}

// NewProjectContainerregistryUserLister returns a new ProjectContainerregistryUserLister.
func NewProjectContainerregistryUserLister(indexer cache.Indexer) ProjectContainerregistryUserLister {
	return &projectContainerregistryUserLister{indexer: indexer}
}

// List lists all ProjectContainerregistryUsers in the indexer.
func (s *projectContainerregistryUserLister) List(selector labels.Selector) (ret []*v1alpha1.ProjectContainerregistryUser, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProjectContainerregistryUser))
	})
	return ret, err
}

// ProjectContainerregistryUsers returns an object that can list and get ProjectContainerregistryUsers.
func (s *projectContainerregistryUserLister) ProjectContainerregistryUsers(namespace string) ProjectContainerregistryUserNamespaceLister {
	return projectContainerregistryUserNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProjectContainerregistryUserNamespaceLister helps list and get ProjectContainerregistryUsers.
// All objects returned here must be treated as read-only.
type ProjectContainerregistryUserNamespaceLister interface {
	// List lists all ProjectContainerregistryUsers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProjectContainerregistryUser, err error)
	// Get retrieves the ProjectContainerregistryUser from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ProjectContainerregistryUser, error)
	ProjectContainerregistryUserNamespaceListerExpansion
}

// projectContainerregistryUserNamespaceLister implements the ProjectContainerregistryUserNamespaceLister
// interface.
type projectContainerregistryUserNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProjectContainerregistryUsers in the indexer for a given namespace.
func (s projectContainerregistryUserNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProjectContainerregistryUser, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProjectContainerregistryUser))
	})
	return ret, err
}

// Get retrieves the ProjectContainerregistryUser from the indexer for a given namespace and name.
func (s projectContainerregistryUserNamespaceLister) Get(name string) (*v1alpha1.ProjectContainerregistryUser, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("projectcontainerregistryuser"), name)
	}
	return obj.(*v1alpha1.ProjectContainerregistryUser), nil
}
