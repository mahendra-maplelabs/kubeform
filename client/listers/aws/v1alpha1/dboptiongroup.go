/*
Copyright The Kubeform Authors.

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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DbOptionGroupLister helps list DbOptionGroups.
type DbOptionGroupLister interface {
	// List lists all DbOptionGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DbOptionGroup, err error)
	// DbOptionGroups returns an object that can list and get DbOptionGroups.
	DbOptionGroups(namespace string) DbOptionGroupNamespaceLister
	DbOptionGroupListerExpansion
}

// dbOptionGroupLister implements the DbOptionGroupLister interface.
type dbOptionGroupLister struct {
	indexer cache.Indexer
}

// NewDbOptionGroupLister returns a new DbOptionGroupLister.
func NewDbOptionGroupLister(indexer cache.Indexer) DbOptionGroupLister {
	return &dbOptionGroupLister{indexer: indexer}
}

// List lists all DbOptionGroups in the indexer.
func (s *dbOptionGroupLister) List(selector labels.Selector) (ret []*v1alpha1.DbOptionGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DbOptionGroup))
	})
	return ret, err
}

// DbOptionGroups returns an object that can list and get DbOptionGroups.
func (s *dbOptionGroupLister) DbOptionGroups(namespace string) DbOptionGroupNamespaceLister {
	return dbOptionGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DbOptionGroupNamespaceLister helps list and get DbOptionGroups.
type DbOptionGroupNamespaceLister interface {
	// List lists all DbOptionGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DbOptionGroup, err error)
	// Get retrieves the DbOptionGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DbOptionGroup, error)
	DbOptionGroupNamespaceListerExpansion
}

// dbOptionGroupNamespaceLister implements the DbOptionGroupNamespaceLister
// interface.
type dbOptionGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DbOptionGroups in the indexer for a given namespace.
func (s dbOptionGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DbOptionGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DbOptionGroup))
	})
	return ret, err
}

// Get retrieves the DbOptionGroup from the indexer for a given namespace and name.
func (s dbOptionGroupNamespaceLister) Get(name string) (*v1alpha1.DbOptionGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dboptiongroup"), name)
	}
	return obj.(*v1alpha1.DbOptionGroup), nil
}
