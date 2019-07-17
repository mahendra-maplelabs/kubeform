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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// ComputeRegionAutoscalersGetter has a method to return a ComputeRegionAutoscalerInterface.
// A group's client should implement this interface.
type ComputeRegionAutoscalersGetter interface {
	ComputeRegionAutoscalers() ComputeRegionAutoscalerInterface
}

// ComputeRegionAutoscalerInterface has methods to work with ComputeRegionAutoscaler resources.
type ComputeRegionAutoscalerInterface interface {
	Create(*v1alpha1.ComputeRegionAutoscaler) (*v1alpha1.ComputeRegionAutoscaler, error)
	Update(*v1alpha1.ComputeRegionAutoscaler) (*v1alpha1.ComputeRegionAutoscaler, error)
	UpdateStatus(*v1alpha1.ComputeRegionAutoscaler) (*v1alpha1.ComputeRegionAutoscaler, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeRegionAutoscaler, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeRegionAutoscalerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRegionAutoscaler, err error)
	ComputeRegionAutoscalerExpansion
}

// computeRegionAutoscalers implements ComputeRegionAutoscalerInterface
type computeRegionAutoscalers struct {
	client rest.Interface
}

// newComputeRegionAutoscalers returns a ComputeRegionAutoscalers
func newComputeRegionAutoscalers(c *GoogleV1alpha1Client) *computeRegionAutoscalers {
	return &computeRegionAutoscalers{
		client: c.RESTClient(),
	}
}

// Get takes name of the computeRegionAutoscaler, and returns the corresponding computeRegionAutoscaler object, and an error if there is any.
func (c *computeRegionAutoscalers) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeRegionAutoscaler, err error) {
	result = &v1alpha1.ComputeRegionAutoscaler{}
	err = c.client.Get().
		Resource("computeregionautoscalers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeRegionAutoscalers that match those selectors.
func (c *computeRegionAutoscalers) List(opts v1.ListOptions) (result *v1alpha1.ComputeRegionAutoscalerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeRegionAutoscalerList{}
	err = c.client.Get().
		Resource("computeregionautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeRegionAutoscalers.
func (c *computeRegionAutoscalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("computeregionautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeRegionAutoscaler and creates it.  Returns the server's representation of the computeRegionAutoscaler, and an error, if there is any.
func (c *computeRegionAutoscalers) Create(computeRegionAutoscaler *v1alpha1.ComputeRegionAutoscaler) (result *v1alpha1.ComputeRegionAutoscaler, err error) {
	result = &v1alpha1.ComputeRegionAutoscaler{}
	err = c.client.Post().
		Resource("computeregionautoscalers").
		Body(computeRegionAutoscaler).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeRegionAutoscaler and updates it. Returns the server's representation of the computeRegionAutoscaler, and an error, if there is any.
func (c *computeRegionAutoscalers) Update(computeRegionAutoscaler *v1alpha1.ComputeRegionAutoscaler) (result *v1alpha1.ComputeRegionAutoscaler, err error) {
	result = &v1alpha1.ComputeRegionAutoscaler{}
	err = c.client.Put().
		Resource("computeregionautoscalers").
		Name(computeRegionAutoscaler.Name).
		Body(computeRegionAutoscaler).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeRegionAutoscalers) UpdateStatus(computeRegionAutoscaler *v1alpha1.ComputeRegionAutoscaler) (result *v1alpha1.ComputeRegionAutoscaler, err error) {
	result = &v1alpha1.ComputeRegionAutoscaler{}
	err = c.client.Put().
		Resource("computeregionautoscalers").
		Name(computeRegionAutoscaler.Name).
		SubResource("status").
		Body(computeRegionAutoscaler).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeRegionAutoscaler and deletes it. Returns an error if one occurs.
func (c *computeRegionAutoscalers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("computeregionautoscalers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeRegionAutoscalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("computeregionautoscalers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeRegionAutoscaler.
func (c *computeRegionAutoscalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeRegionAutoscaler, err error) {
	result = &v1alpha1.ComputeRegionAutoscaler{}
	err = c.client.Patch(pt).
		Resource("computeregionautoscalers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}