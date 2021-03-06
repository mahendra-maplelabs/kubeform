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

// IamOpenidConnectProviderLister helps list IamOpenidConnectProviders.
type IamOpenidConnectProviderLister interface {
	// List lists all IamOpenidConnectProviders in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IamOpenidConnectProvider, err error)
	// IamOpenidConnectProviders returns an object that can list and get IamOpenidConnectProviders.
	IamOpenidConnectProviders(namespace string) IamOpenidConnectProviderNamespaceLister
	IamOpenidConnectProviderListerExpansion
}

// iamOpenidConnectProviderLister implements the IamOpenidConnectProviderLister interface.
type iamOpenidConnectProviderLister struct {
	indexer cache.Indexer
}

// NewIamOpenidConnectProviderLister returns a new IamOpenidConnectProviderLister.
func NewIamOpenidConnectProviderLister(indexer cache.Indexer) IamOpenidConnectProviderLister {
	return &iamOpenidConnectProviderLister{indexer: indexer}
}

// List lists all IamOpenidConnectProviders in the indexer.
func (s *iamOpenidConnectProviderLister) List(selector labels.Selector) (ret []*v1alpha1.IamOpenidConnectProvider, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamOpenidConnectProvider))
	})
	return ret, err
}

// IamOpenidConnectProviders returns an object that can list and get IamOpenidConnectProviders.
func (s *iamOpenidConnectProviderLister) IamOpenidConnectProviders(namespace string) IamOpenidConnectProviderNamespaceLister {
	return iamOpenidConnectProviderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamOpenidConnectProviderNamespaceLister helps list and get IamOpenidConnectProviders.
type IamOpenidConnectProviderNamespaceLister interface {
	// List lists all IamOpenidConnectProviders in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IamOpenidConnectProvider, err error)
	// Get retrieves the IamOpenidConnectProvider from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IamOpenidConnectProvider, error)
	IamOpenidConnectProviderNamespaceListerExpansion
}

// iamOpenidConnectProviderNamespaceLister implements the IamOpenidConnectProviderNamespaceLister
// interface.
type iamOpenidConnectProviderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamOpenidConnectProviders in the indexer for a given namespace.
func (s iamOpenidConnectProviderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamOpenidConnectProvider, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamOpenidConnectProvider))
	})
	return ret, err
}

// Get retrieves the IamOpenidConnectProvider from the indexer for a given namespace and name.
func (s iamOpenidConnectProviderNamespaceLister) Get(name string) (*v1alpha1.IamOpenidConnectProvider, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iamopenidconnectprovider"), name)
	}
	return obj.(*v1alpha1.IamOpenidConnectProvider), nil
}
