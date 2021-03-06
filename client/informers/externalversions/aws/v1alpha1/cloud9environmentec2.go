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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	awsv1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	versioned "kubeform.dev/kubeform/client/clientset/versioned"
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/kubeform/client/listers/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// Cloud9EnvironmentEc2Informer provides access to a shared informer and lister for
// Cloud9EnvironmentEc2s.
type Cloud9EnvironmentEc2Informer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.Cloud9EnvironmentEc2Lister
}

type cloud9EnvironmentEc2Informer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCloud9EnvironmentEc2Informer constructs a new informer for Cloud9EnvironmentEc2 type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCloud9EnvironmentEc2Informer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCloud9EnvironmentEc2Informer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCloud9EnvironmentEc2Informer constructs a new informer for Cloud9EnvironmentEc2 type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCloud9EnvironmentEc2Informer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().Cloud9EnvironmentEc2s(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AwsV1alpha1().Cloud9EnvironmentEc2s(namespace).Watch(options)
			},
		},
		&awsv1alpha1.Cloud9EnvironmentEc2{},
		resyncPeriod,
		indexers,
	)
}

func (f *cloud9EnvironmentEc2Informer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCloud9EnvironmentEc2Informer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cloud9EnvironmentEc2Informer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&awsv1alpha1.Cloud9EnvironmentEc2{}, f.defaultInformer)
}

func (f *cloud9EnvironmentEc2Informer) Lister() v1alpha1.Cloud9EnvironmentEc2Lister {
	return v1alpha1.NewCloud9EnvironmentEc2Lister(f.Informer().GetIndexer())
}
