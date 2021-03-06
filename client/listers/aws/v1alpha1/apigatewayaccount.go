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

// ApiGatewayAccountLister helps list ApiGatewayAccounts.
type ApiGatewayAccountLister interface {
	// List lists all ApiGatewayAccounts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayAccount, err error)
	// ApiGatewayAccounts returns an object that can list and get ApiGatewayAccounts.
	ApiGatewayAccounts(namespace string) ApiGatewayAccountNamespaceLister
	ApiGatewayAccountListerExpansion
}

// apiGatewayAccountLister implements the ApiGatewayAccountLister interface.
type apiGatewayAccountLister struct {
	indexer cache.Indexer
}

// NewApiGatewayAccountLister returns a new ApiGatewayAccountLister.
func NewApiGatewayAccountLister(indexer cache.Indexer) ApiGatewayAccountLister {
	return &apiGatewayAccountLister{indexer: indexer}
}

// List lists all ApiGatewayAccounts in the indexer.
func (s *apiGatewayAccountLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayAccount))
	})
	return ret, err
}

// ApiGatewayAccounts returns an object that can list and get ApiGatewayAccounts.
func (s *apiGatewayAccountLister) ApiGatewayAccounts(namespace string) ApiGatewayAccountNamespaceLister {
	return apiGatewayAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiGatewayAccountNamespaceLister helps list and get ApiGatewayAccounts.
type ApiGatewayAccountNamespaceLister interface {
	// List lists all ApiGatewayAccounts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayAccount, err error)
	// Get retrieves the ApiGatewayAccount from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiGatewayAccount, error)
	ApiGatewayAccountNamespaceListerExpansion
}

// apiGatewayAccountNamespaceLister implements the ApiGatewayAccountNamespaceLister
// interface.
type apiGatewayAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiGatewayAccounts in the indexer for a given namespace.
func (s apiGatewayAccountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayAccount))
	})
	return ret, err
}

// Get retrieves the ApiGatewayAccount from the indexer for a given namespace and name.
func (s apiGatewayAccountNamespaceLister) Get(name string) (*v1alpha1.ApiGatewayAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apigatewayaccount"), name)
	}
	return obj.(*v1alpha1.ApiGatewayAccount), nil
}
