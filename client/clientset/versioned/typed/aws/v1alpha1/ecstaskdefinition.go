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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EcsTaskDefinitionsGetter has a method to return a EcsTaskDefinitionInterface.
// A group's client should implement this interface.
type EcsTaskDefinitionsGetter interface {
	EcsTaskDefinitions(namespace string) EcsTaskDefinitionInterface
}

// EcsTaskDefinitionInterface has methods to work with EcsTaskDefinition resources.
type EcsTaskDefinitionInterface interface {
	Create(*v1alpha1.EcsTaskDefinition) (*v1alpha1.EcsTaskDefinition, error)
	Update(*v1alpha1.EcsTaskDefinition) (*v1alpha1.EcsTaskDefinition, error)
	UpdateStatus(*v1alpha1.EcsTaskDefinition) (*v1alpha1.EcsTaskDefinition, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EcsTaskDefinition, error)
	List(opts v1.ListOptions) (*v1alpha1.EcsTaskDefinitionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EcsTaskDefinition, err error)
	EcsTaskDefinitionExpansion
}

// ecsTaskDefinitions implements EcsTaskDefinitionInterface
type ecsTaskDefinitions struct {
	client rest.Interface
	ns     string
}

// newEcsTaskDefinitions returns a EcsTaskDefinitions
func newEcsTaskDefinitions(c *AwsV1alpha1Client, namespace string) *ecsTaskDefinitions {
	return &ecsTaskDefinitions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ecsTaskDefinition, and returns the corresponding ecsTaskDefinition object, and an error if there is any.
func (c *ecsTaskDefinitions) Get(name string, options v1.GetOptions) (result *v1alpha1.EcsTaskDefinition, err error) {
	result = &v1alpha1.EcsTaskDefinition{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ecstaskdefinitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EcsTaskDefinitions that match those selectors.
func (c *ecsTaskDefinitions) List(opts v1.ListOptions) (result *v1alpha1.EcsTaskDefinitionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EcsTaskDefinitionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ecstaskdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ecsTaskDefinitions.
func (c *ecsTaskDefinitions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ecstaskdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ecsTaskDefinition and creates it.  Returns the server's representation of the ecsTaskDefinition, and an error, if there is any.
func (c *ecsTaskDefinitions) Create(ecsTaskDefinition *v1alpha1.EcsTaskDefinition) (result *v1alpha1.EcsTaskDefinition, err error) {
	result = &v1alpha1.EcsTaskDefinition{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ecstaskdefinitions").
		Body(ecsTaskDefinition).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ecsTaskDefinition and updates it. Returns the server's representation of the ecsTaskDefinition, and an error, if there is any.
func (c *ecsTaskDefinitions) Update(ecsTaskDefinition *v1alpha1.EcsTaskDefinition) (result *v1alpha1.EcsTaskDefinition, err error) {
	result = &v1alpha1.EcsTaskDefinition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ecstaskdefinitions").
		Name(ecsTaskDefinition.Name).
		Body(ecsTaskDefinition).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ecsTaskDefinitions) UpdateStatus(ecsTaskDefinition *v1alpha1.EcsTaskDefinition) (result *v1alpha1.EcsTaskDefinition, err error) {
	result = &v1alpha1.EcsTaskDefinition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ecstaskdefinitions").
		Name(ecsTaskDefinition.Name).
		SubResource("status").
		Body(ecsTaskDefinition).
		Do().
		Into(result)
	return
}

// Delete takes name of the ecsTaskDefinition and deletes it. Returns an error if one occurs.
func (c *ecsTaskDefinitions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ecstaskdefinitions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ecsTaskDefinitions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ecstaskdefinitions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ecsTaskDefinition.
func (c *ecsTaskDefinitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EcsTaskDefinition, err error) {
	result = &v1alpha1.EcsTaskDefinition{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ecstaskdefinitions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
