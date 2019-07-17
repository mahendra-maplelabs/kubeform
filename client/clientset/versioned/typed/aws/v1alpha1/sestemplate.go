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

// SesTemplatesGetter has a method to return a SesTemplateInterface.
// A group's client should implement this interface.
type SesTemplatesGetter interface {
	SesTemplates() SesTemplateInterface
}

// SesTemplateInterface has methods to work with SesTemplate resources.
type SesTemplateInterface interface {
	Create(*v1alpha1.SesTemplate) (*v1alpha1.SesTemplate, error)
	Update(*v1alpha1.SesTemplate) (*v1alpha1.SesTemplate, error)
	UpdateStatus(*v1alpha1.SesTemplate) (*v1alpha1.SesTemplate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SesTemplate, error)
	List(opts v1.ListOptions) (*v1alpha1.SesTemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesTemplate, err error)
	SesTemplateExpansion
}

// sesTemplates implements SesTemplateInterface
type sesTemplates struct {
	client rest.Interface
}

// newSesTemplates returns a SesTemplates
func newSesTemplates(c *AwsV1alpha1Client) *sesTemplates {
	return &sesTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the sesTemplate, and returns the corresponding sesTemplate object, and an error if there is any.
func (c *sesTemplates) Get(name string, options v1.GetOptions) (result *v1alpha1.SesTemplate, err error) {
	result = &v1alpha1.SesTemplate{}
	err = c.client.Get().
		Resource("sestemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SesTemplates that match those selectors.
func (c *sesTemplates) List(opts v1.ListOptions) (result *v1alpha1.SesTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SesTemplateList{}
	err = c.client.Get().
		Resource("sestemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sesTemplates.
func (c *sesTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("sestemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sesTemplate and creates it.  Returns the server's representation of the sesTemplate, and an error, if there is any.
func (c *sesTemplates) Create(sesTemplate *v1alpha1.SesTemplate) (result *v1alpha1.SesTemplate, err error) {
	result = &v1alpha1.SesTemplate{}
	err = c.client.Post().
		Resource("sestemplates").
		Body(sesTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sesTemplate and updates it. Returns the server's representation of the sesTemplate, and an error, if there is any.
func (c *sesTemplates) Update(sesTemplate *v1alpha1.SesTemplate) (result *v1alpha1.SesTemplate, err error) {
	result = &v1alpha1.SesTemplate{}
	err = c.client.Put().
		Resource("sestemplates").
		Name(sesTemplate.Name).
		Body(sesTemplate).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sesTemplates) UpdateStatus(sesTemplate *v1alpha1.SesTemplate) (result *v1alpha1.SesTemplate, err error) {
	result = &v1alpha1.SesTemplate{}
	err = c.client.Put().
		Resource("sestemplates").
		Name(sesTemplate.Name).
		SubResource("status").
		Body(sesTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the sesTemplate and deletes it. Returns an error if one occurs.
func (c *sesTemplates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("sestemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sesTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("sestemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sesTemplate.
func (c *sesTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SesTemplate, err error) {
	result = &v1alpha1.SesTemplate{}
	err = c.client.Patch(pt).
		Resource("sestemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}