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
	v1alpha1 "kubeform.dev/provider-ovh-api/apis/dbaas/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LogsOutputGraylogStreamLister helps list LogsOutputGraylogStreams.
// All objects returned here must be treated as read-only.
type LogsOutputGraylogStreamLister interface {
	// List lists all LogsOutputGraylogStreams in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LogsOutputGraylogStream, err error)
	// LogsOutputGraylogStreams returns an object that can list and get LogsOutputGraylogStreams.
	LogsOutputGraylogStreams(namespace string) LogsOutputGraylogStreamNamespaceLister
	LogsOutputGraylogStreamListerExpansion
}

// logsOutputGraylogStreamLister implements the LogsOutputGraylogStreamLister interface.
type logsOutputGraylogStreamLister struct {
	indexer cache.Indexer
}

// NewLogsOutputGraylogStreamLister returns a new LogsOutputGraylogStreamLister.
func NewLogsOutputGraylogStreamLister(indexer cache.Indexer) LogsOutputGraylogStreamLister {
	return &logsOutputGraylogStreamLister{indexer: indexer}
}

// List lists all LogsOutputGraylogStreams in the indexer.
func (s *logsOutputGraylogStreamLister) List(selector labels.Selector) (ret []*v1alpha1.LogsOutputGraylogStream, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LogsOutputGraylogStream))
	})
	return ret, err
}

// LogsOutputGraylogStreams returns an object that can list and get LogsOutputGraylogStreams.
func (s *logsOutputGraylogStreamLister) LogsOutputGraylogStreams(namespace string) LogsOutputGraylogStreamNamespaceLister {
	return logsOutputGraylogStreamNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LogsOutputGraylogStreamNamespaceLister helps list and get LogsOutputGraylogStreams.
// All objects returned here must be treated as read-only.
type LogsOutputGraylogStreamNamespaceLister interface {
	// List lists all LogsOutputGraylogStreams in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LogsOutputGraylogStream, err error)
	// Get retrieves the LogsOutputGraylogStream from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LogsOutputGraylogStream, error)
	LogsOutputGraylogStreamNamespaceListerExpansion
}

// logsOutputGraylogStreamNamespaceLister implements the LogsOutputGraylogStreamNamespaceLister
// interface.
type logsOutputGraylogStreamNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LogsOutputGraylogStreams in the indexer for a given namespace.
func (s logsOutputGraylogStreamNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LogsOutputGraylogStream, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LogsOutputGraylogStream))
	})
	return ret, err
}

// Get retrieves the LogsOutputGraylogStream from the indexer for a given namespace and name.
func (s logsOutputGraylogStreamNamespaceLister) Get(name string) (*v1alpha1.LogsOutputGraylogStream, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("logsoutputgraylogstream"), name)
	}
	return obj.(*v1alpha1.LogsOutputGraylogStream), nil
}
