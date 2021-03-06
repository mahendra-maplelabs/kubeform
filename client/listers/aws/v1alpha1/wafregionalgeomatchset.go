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

// WafregionalGeoMatchSetLister helps list WafregionalGeoMatchSets.
type WafregionalGeoMatchSetLister interface {
	// List lists all WafregionalGeoMatchSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.WafregionalGeoMatchSet, err error)
	// WafregionalGeoMatchSets returns an object that can list and get WafregionalGeoMatchSets.
	WafregionalGeoMatchSets(namespace string) WafregionalGeoMatchSetNamespaceLister
	WafregionalGeoMatchSetListerExpansion
}

// wafregionalGeoMatchSetLister implements the WafregionalGeoMatchSetLister interface.
type wafregionalGeoMatchSetLister struct {
	indexer cache.Indexer
}

// NewWafregionalGeoMatchSetLister returns a new WafregionalGeoMatchSetLister.
func NewWafregionalGeoMatchSetLister(indexer cache.Indexer) WafregionalGeoMatchSetLister {
	return &wafregionalGeoMatchSetLister{indexer: indexer}
}

// List lists all WafregionalGeoMatchSets in the indexer.
func (s *wafregionalGeoMatchSetLister) List(selector labels.Selector) (ret []*v1alpha1.WafregionalGeoMatchSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafregionalGeoMatchSet))
	})
	return ret, err
}

// WafregionalGeoMatchSets returns an object that can list and get WafregionalGeoMatchSets.
func (s *wafregionalGeoMatchSetLister) WafregionalGeoMatchSets(namespace string) WafregionalGeoMatchSetNamespaceLister {
	return wafregionalGeoMatchSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WafregionalGeoMatchSetNamespaceLister helps list and get WafregionalGeoMatchSets.
type WafregionalGeoMatchSetNamespaceLister interface {
	// List lists all WafregionalGeoMatchSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.WafregionalGeoMatchSet, err error)
	// Get retrieves the WafregionalGeoMatchSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.WafregionalGeoMatchSet, error)
	WafregionalGeoMatchSetNamespaceListerExpansion
}

// wafregionalGeoMatchSetNamespaceLister implements the WafregionalGeoMatchSetNamespaceLister
// interface.
type wafregionalGeoMatchSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WafregionalGeoMatchSets in the indexer for a given namespace.
func (s wafregionalGeoMatchSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WafregionalGeoMatchSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafregionalGeoMatchSet))
	})
	return ret, err
}

// Get retrieves the WafregionalGeoMatchSet from the indexer for a given namespace and name.
func (s wafregionalGeoMatchSetNamespaceLister) Get(name string) (*v1alpha1.WafregionalGeoMatchSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("wafregionalgeomatchset"), name)
	}
	return obj.(*v1alpha1.WafregionalGeoMatchSet), nil
}
