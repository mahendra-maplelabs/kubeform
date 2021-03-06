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

// ApiGatewayModelLister helps list ApiGatewayModels.
type ApiGatewayModelLister interface {
	// List lists all ApiGatewayModels in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayModel, err error)
	// ApiGatewayModels returns an object that can list and get ApiGatewayModels.
	ApiGatewayModels(namespace string) ApiGatewayModelNamespaceLister
	ApiGatewayModelListerExpansion
}

// apiGatewayModelLister implements the ApiGatewayModelLister interface.
type apiGatewayModelLister struct {
	indexer cache.Indexer
}

// NewApiGatewayModelLister returns a new ApiGatewayModelLister.
func NewApiGatewayModelLister(indexer cache.Indexer) ApiGatewayModelLister {
	return &apiGatewayModelLister{indexer: indexer}
}

// List lists all ApiGatewayModels in the indexer.
func (s *apiGatewayModelLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayModel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayModel))
	})
	return ret, err
}

// ApiGatewayModels returns an object that can list and get ApiGatewayModels.
func (s *apiGatewayModelLister) ApiGatewayModels(namespace string) ApiGatewayModelNamespaceLister {
	return apiGatewayModelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiGatewayModelNamespaceLister helps list and get ApiGatewayModels.
type ApiGatewayModelNamespaceLister interface {
	// List lists all ApiGatewayModels in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayModel, err error)
	// Get retrieves the ApiGatewayModel from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiGatewayModel, error)
	ApiGatewayModelNamespaceListerExpansion
}

// apiGatewayModelNamespaceLister implements the ApiGatewayModelNamespaceLister
// interface.
type apiGatewayModelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiGatewayModels in the indexer for a given namespace.
func (s apiGatewayModelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayModel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayModel))
	})
	return ret, err
}

// Get retrieves the ApiGatewayModel from the indexer for a given namespace and name.
func (s apiGatewayModelNamespaceLister) Get(name string) (*v1alpha1.ApiGatewayModel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apigatewaymodel"), name)
	}
	return obj.(*v1alpha1.ApiGatewayModel), nil
}
