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

// EcrRepositoryLister helps list EcrRepositories.
type EcrRepositoryLister interface {
	// List lists all EcrRepositories in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.EcrRepository, err error)
	// EcrRepositories returns an object that can list and get EcrRepositories.
	EcrRepositories(namespace string) EcrRepositoryNamespaceLister
	EcrRepositoryListerExpansion
}

// ecrRepositoryLister implements the EcrRepositoryLister interface.
type ecrRepositoryLister struct {
	indexer cache.Indexer
}

// NewEcrRepositoryLister returns a new EcrRepositoryLister.
func NewEcrRepositoryLister(indexer cache.Indexer) EcrRepositoryLister {
	return &ecrRepositoryLister{indexer: indexer}
}

// List lists all EcrRepositories in the indexer.
func (s *ecrRepositoryLister) List(selector labels.Selector) (ret []*v1alpha1.EcrRepository, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EcrRepository))
	})
	return ret, err
}

// EcrRepositories returns an object that can list and get EcrRepositories.
func (s *ecrRepositoryLister) EcrRepositories(namespace string) EcrRepositoryNamespaceLister {
	return ecrRepositoryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EcrRepositoryNamespaceLister helps list and get EcrRepositories.
type EcrRepositoryNamespaceLister interface {
	// List lists all EcrRepositories in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.EcrRepository, err error)
	// Get retrieves the EcrRepository from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.EcrRepository, error)
	EcrRepositoryNamespaceListerExpansion
}

// ecrRepositoryNamespaceLister implements the EcrRepositoryNamespaceLister
// interface.
type ecrRepositoryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EcrRepositories in the indexer for a given namespace.
func (s ecrRepositoryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EcrRepository, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EcrRepository))
	})
	return ret, err
}

// Get retrieves the EcrRepository from the indexer for a given namespace and name.
func (s ecrRepositoryNamespaceLister) Get(name string) (*v1alpha1.EcrRepository, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ecrrepository"), name)
	}
	return obj.(*v1alpha1.EcrRepository), nil
}
