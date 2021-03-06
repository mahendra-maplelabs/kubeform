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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SchedulerJobCollectionsGetter has a method to return a SchedulerJobCollectionInterface.
// A group's client should implement this interface.
type SchedulerJobCollectionsGetter interface {
	SchedulerJobCollections(namespace string) SchedulerJobCollectionInterface
}

// SchedulerJobCollectionInterface has methods to work with SchedulerJobCollection resources.
type SchedulerJobCollectionInterface interface {
	Create(*v1alpha1.SchedulerJobCollection) (*v1alpha1.SchedulerJobCollection, error)
	Update(*v1alpha1.SchedulerJobCollection) (*v1alpha1.SchedulerJobCollection, error)
	UpdateStatus(*v1alpha1.SchedulerJobCollection) (*v1alpha1.SchedulerJobCollection, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SchedulerJobCollection, error)
	List(opts v1.ListOptions) (*v1alpha1.SchedulerJobCollectionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SchedulerJobCollection, err error)
	SchedulerJobCollectionExpansion
}

// schedulerJobCollections implements SchedulerJobCollectionInterface
type schedulerJobCollections struct {
	client rest.Interface
	ns     string
}

// newSchedulerJobCollections returns a SchedulerJobCollections
func newSchedulerJobCollections(c *AzurermV1alpha1Client, namespace string) *schedulerJobCollections {
	return &schedulerJobCollections{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the schedulerJobCollection, and returns the corresponding schedulerJobCollection object, and an error if there is any.
func (c *schedulerJobCollections) Get(name string, options v1.GetOptions) (result *v1alpha1.SchedulerJobCollection, err error) {
	result = &v1alpha1.SchedulerJobCollection{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("schedulerjobcollections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SchedulerJobCollections that match those selectors.
func (c *schedulerJobCollections) List(opts v1.ListOptions) (result *v1alpha1.SchedulerJobCollectionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SchedulerJobCollectionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("schedulerjobcollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested schedulerJobCollections.
func (c *schedulerJobCollections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("schedulerjobcollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a schedulerJobCollection and creates it.  Returns the server's representation of the schedulerJobCollection, and an error, if there is any.
func (c *schedulerJobCollections) Create(schedulerJobCollection *v1alpha1.SchedulerJobCollection) (result *v1alpha1.SchedulerJobCollection, err error) {
	result = &v1alpha1.SchedulerJobCollection{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("schedulerjobcollections").
		Body(schedulerJobCollection).
		Do().
		Into(result)
	return
}

// Update takes the representation of a schedulerJobCollection and updates it. Returns the server's representation of the schedulerJobCollection, and an error, if there is any.
func (c *schedulerJobCollections) Update(schedulerJobCollection *v1alpha1.SchedulerJobCollection) (result *v1alpha1.SchedulerJobCollection, err error) {
	result = &v1alpha1.SchedulerJobCollection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("schedulerjobcollections").
		Name(schedulerJobCollection.Name).
		Body(schedulerJobCollection).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *schedulerJobCollections) UpdateStatus(schedulerJobCollection *v1alpha1.SchedulerJobCollection) (result *v1alpha1.SchedulerJobCollection, err error) {
	result = &v1alpha1.SchedulerJobCollection{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("schedulerjobcollections").
		Name(schedulerJobCollection.Name).
		SubResource("status").
		Body(schedulerJobCollection).
		Do().
		Into(result)
	return
}

// Delete takes name of the schedulerJobCollection and deletes it. Returns an error if one occurs.
func (c *schedulerJobCollections) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("schedulerjobcollections").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *schedulerJobCollections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("schedulerjobcollections").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched schedulerJobCollection.
func (c *schedulerJobCollections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SchedulerJobCollection, err error) {
	result = &v1alpha1.SchedulerJobCollection{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("schedulerjobcollections").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
