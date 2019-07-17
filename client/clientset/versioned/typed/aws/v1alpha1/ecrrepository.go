/*
Copyright 2019 The Kubeform Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// EcrRepositoriesGetter has a method to return a EcrRepositoryInterface.
// A group's client should implement this interface.
type EcrRepositoriesGetter interface {
	EcrRepositories() EcrRepositoryInterface
}

// EcrRepositoryInterface has methods to work with EcrRepository resources.
type EcrRepositoryInterface interface {
	Create(*v1alpha1.EcrRepository) (*v1alpha1.EcrRepository, error)
	Update(*v1alpha1.EcrRepository) (*v1alpha1.EcrRepository, error)
	UpdateStatus(*v1alpha1.EcrRepository) (*v1alpha1.EcrRepository, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EcrRepository, error)
	List(opts v1.ListOptions) (*v1alpha1.EcrRepositoryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EcrRepository, err error)
	EcrRepositoryExpansion
}

// ecrRepositories implements EcrRepositoryInterface
type ecrRepositories struct {
	client rest.Interface
}

// newEcrRepositories returns a EcrRepositories
func newEcrRepositories(c *AwsV1alpha1Client) *ecrRepositories {
	return &ecrRepositories{
		client: c.RESTClient(),
	}
}

// Get takes name of the ecrRepository, and returns the corresponding ecrRepository object, and an error if there is any.
func (c *ecrRepositories) Get(name string, options v1.GetOptions) (result *v1alpha1.EcrRepository, err error) {
	result = &v1alpha1.EcrRepository{}
	err = c.client.Get().
		Resource("ecrrepositories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EcrRepositories that match those selectors.
func (c *ecrRepositories) List(opts v1.ListOptions) (result *v1alpha1.EcrRepositoryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EcrRepositoryList{}
	err = c.client.Get().
		Resource("ecrrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ecrRepositories.
func (c *ecrRepositories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ecrrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ecrRepository and creates it.  Returns the server's representation of the ecrRepository, and an error, if there is any.
func (c *ecrRepositories) Create(ecrRepository *v1alpha1.EcrRepository) (result *v1alpha1.EcrRepository, err error) {
	result = &v1alpha1.EcrRepository{}
	err = c.client.Post().
		Resource("ecrrepositories").
		Body(ecrRepository).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ecrRepository and updates it. Returns the server's representation of the ecrRepository, and an error, if there is any.
func (c *ecrRepositories) Update(ecrRepository *v1alpha1.EcrRepository) (result *v1alpha1.EcrRepository, err error) {
	result = &v1alpha1.EcrRepository{}
	err = c.client.Put().
		Resource("ecrrepositories").
		Name(ecrRepository.Name).
		Body(ecrRepository).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ecrRepositories) UpdateStatus(ecrRepository *v1alpha1.EcrRepository) (result *v1alpha1.EcrRepository, err error) {
	result = &v1alpha1.EcrRepository{}
	err = c.client.Put().
		Resource("ecrrepositories").
		Name(ecrRepository.Name).
		SubResource("status").
		Body(ecrRepository).
		Do().
		Into(result)
	return
}

// Delete takes name of the ecrRepository and deletes it. Returns an error if one occurs.
func (c *ecrRepositories) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ecrrepositories").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ecrRepositories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ecrrepositories").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ecrRepository.
func (c *ecrRepositories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EcrRepository, err error) {
	result = &v1alpha1.EcrRepository{}
	err = c.client.Patch(pt).
		Resource("ecrrepositories").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}