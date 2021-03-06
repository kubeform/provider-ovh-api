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

// ServerInstallTaskLister helps list ServerInstallTasks.
// All objects returned here must be treated as read-only.
type ServerInstallTaskLister interface {
	// List lists all ServerInstallTasks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServerInstallTask, err error)
	// ServerInstallTasks returns an object that can list and get ServerInstallTasks.
	ServerInstallTasks(namespace string) ServerInstallTaskNamespaceLister
	ServerInstallTaskListerExpansion
}

// serverInstallTaskLister implements the ServerInstallTaskLister interface.
type serverInstallTaskLister struct {
	indexer cache.Indexer
}

// NewServerInstallTaskLister returns a new ServerInstallTaskLister.
func NewServerInstallTaskLister(indexer cache.Indexer) ServerInstallTaskLister {
	return &serverInstallTaskLister{indexer: indexer}
}

// List lists all ServerInstallTasks in the indexer.
func (s *serverInstallTaskLister) List(selector labels.Selector) (ret []*v1alpha1.ServerInstallTask, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServerInstallTask))
	})
	return ret, err
}

// ServerInstallTasks returns an object that can list and get ServerInstallTasks.
func (s *serverInstallTaskLister) ServerInstallTasks(namespace string) ServerInstallTaskNamespaceLister {
	return serverInstallTaskNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServerInstallTaskNamespaceLister helps list and get ServerInstallTasks.
// All objects returned here must be treated as read-only.
type ServerInstallTaskNamespaceLister interface {
	// List lists all ServerInstallTasks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServerInstallTask, err error)
	// Get retrieves the ServerInstallTask from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServerInstallTask, error)
	ServerInstallTaskNamespaceListerExpansion
}

// serverInstallTaskNamespaceLister implements the ServerInstallTaskNamespaceLister
// interface.
type serverInstallTaskNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServerInstallTasks in the indexer for a given namespace.
func (s serverInstallTaskNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServerInstallTask, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServerInstallTask))
	})
	return ret, err
}

// Get retrieves the ServerInstallTask from the indexer for a given namespace and name.
func (s serverInstallTaskNamespaceLister) Get(name string) (*v1alpha1.ServerInstallTask, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("serverinstalltask"), name)
	}
	return obj.(*v1alpha1.ServerInstallTask), nil
}
