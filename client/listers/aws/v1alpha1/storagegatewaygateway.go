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

// StoragegatewayGatewayLister helps list StoragegatewayGateways.
type StoragegatewayGatewayLister interface {
	// List lists all StoragegatewayGateways in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StoragegatewayGateway, err error)
	// StoragegatewayGateways returns an object that can list and get StoragegatewayGateways.
	StoragegatewayGateways(namespace string) StoragegatewayGatewayNamespaceLister
	StoragegatewayGatewayListerExpansion
}

// storagegatewayGatewayLister implements the StoragegatewayGatewayLister interface.
type storagegatewayGatewayLister struct {
	indexer cache.Indexer
}

// NewStoragegatewayGatewayLister returns a new StoragegatewayGatewayLister.
func NewStoragegatewayGatewayLister(indexer cache.Indexer) StoragegatewayGatewayLister {
	return &storagegatewayGatewayLister{indexer: indexer}
}

// List lists all StoragegatewayGateways in the indexer.
func (s *storagegatewayGatewayLister) List(selector labels.Selector) (ret []*v1alpha1.StoragegatewayGateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StoragegatewayGateway))
	})
	return ret, err
}

// StoragegatewayGateways returns an object that can list and get StoragegatewayGateways.
func (s *storagegatewayGatewayLister) StoragegatewayGateways(namespace string) StoragegatewayGatewayNamespaceLister {
	return storagegatewayGatewayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StoragegatewayGatewayNamespaceLister helps list and get StoragegatewayGateways.
type StoragegatewayGatewayNamespaceLister interface {
	// List lists all StoragegatewayGateways in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StoragegatewayGateway, err error)
	// Get retrieves the StoragegatewayGateway from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StoragegatewayGateway, error)
	StoragegatewayGatewayNamespaceListerExpansion
}

// storagegatewayGatewayNamespaceLister implements the StoragegatewayGatewayNamespaceLister
// interface.
type storagegatewayGatewayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StoragegatewayGateways in the indexer for a given namespace.
func (s storagegatewayGatewayNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StoragegatewayGateway, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StoragegatewayGateway))
	})
	return ret, err
}

// Get retrieves the StoragegatewayGateway from the indexer for a given namespace and name.
func (s storagegatewayGatewayNamespaceLister) Get(name string) (*v1alpha1.StoragegatewayGateway, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagegatewaygateway"), name)
	}
	return obj.(*v1alpha1.StoragegatewayGateway), nil
}
