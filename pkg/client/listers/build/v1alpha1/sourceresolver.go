/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package v1alpha1

import (
	v1alpha1 "github.com/pivotal/build-service-system/pkg/apis/build/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SourceResolverLister helps list SourceResolvers.
type SourceResolverLister interface {
	// List lists all SourceResolvers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SourceResolver, err error)
	// SourceResolvers returns an object that can list and get SourceResolvers.
	SourceResolvers(namespace string) SourceResolverNamespaceLister
	SourceResolverListerExpansion
}

// sourceResolverLister implements the SourceResolverLister interface.
type sourceResolverLister struct {
	indexer cache.Indexer
}

// NewSourceResolverLister returns a new SourceResolverLister.
func NewSourceResolverLister(indexer cache.Indexer) SourceResolverLister {
	return &sourceResolverLister{indexer: indexer}
}

// List lists all SourceResolvers in the indexer.
func (s *sourceResolverLister) List(selector labels.Selector) (ret []*v1alpha1.SourceResolver, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SourceResolver))
	})
	return ret, err
}

// SourceResolvers returns an object that can list and get SourceResolvers.
func (s *sourceResolverLister) SourceResolvers(namespace string) SourceResolverNamespaceLister {
	return sourceResolverNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SourceResolverNamespaceLister helps list and get SourceResolvers.
type SourceResolverNamespaceLister interface {
	// List lists all SourceResolvers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SourceResolver, err error)
	// Get retrieves the SourceResolver from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SourceResolver, error)
	SourceResolverNamespaceListerExpansion
}

// sourceResolverNamespaceLister implements the SourceResolverNamespaceLister
// interface.
type sourceResolverNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SourceResolvers in the indexer for a given namespace.
func (s sourceResolverNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SourceResolver, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SourceResolver))
	})
	return ret, err
}

// Get retrieves the SourceResolver from the indexer for a given namespace and name.
func (s sourceResolverNamespaceLister) Get(name string) (*v1alpha1.SourceResolver, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sourceresolver"), name)
	}
	return obj.(*v1alpha1.SourceResolver), nil
}