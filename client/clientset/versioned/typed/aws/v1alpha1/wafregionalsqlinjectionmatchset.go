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

// WafregionalSQLInjectionMatchSetsGetter has a method to return a WafregionalSQLInjectionMatchSetInterface.
// A group's client should implement this interface.
type WafregionalSQLInjectionMatchSetsGetter interface {
	WafregionalSQLInjectionMatchSets(namespace string) WafregionalSQLInjectionMatchSetInterface
}

// WafregionalSQLInjectionMatchSetInterface has methods to work with WafregionalSQLInjectionMatchSet resources.
type WafregionalSQLInjectionMatchSetInterface interface {
	Create(*v1alpha1.WafregionalSQLInjectionMatchSet) (*v1alpha1.WafregionalSQLInjectionMatchSet, error)
	Update(*v1alpha1.WafregionalSQLInjectionMatchSet) (*v1alpha1.WafregionalSQLInjectionMatchSet, error)
	UpdateStatus(*v1alpha1.WafregionalSQLInjectionMatchSet) (*v1alpha1.WafregionalSQLInjectionMatchSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.WafregionalSQLInjectionMatchSet, error)
	List(opts v1.ListOptions) (*v1alpha1.WafregionalSQLInjectionMatchSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalSQLInjectionMatchSet, err error)
	WafregionalSQLInjectionMatchSetExpansion
}

// wafregionalSQLInjectionMatchSets implements WafregionalSQLInjectionMatchSetInterface
type wafregionalSQLInjectionMatchSets struct {
	client rest.Interface
	ns     string
}

// newWafregionalSQLInjectionMatchSets returns a WafregionalSQLInjectionMatchSets
func newWafregionalSQLInjectionMatchSets(c *AwsV1alpha1Client, namespace string) *wafregionalSQLInjectionMatchSets {
	return &wafregionalSQLInjectionMatchSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the wafregionalSQLInjectionMatchSet, and returns the corresponding wafregionalSQLInjectionMatchSet object, and an error if there is any.
func (c *wafregionalSQLInjectionMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.WafregionalSQLInjectionMatchSet, err error) {
	result = &v1alpha1.WafregionalSQLInjectionMatchSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalsqlinjectionmatchsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WafregionalSQLInjectionMatchSets that match those selectors.
func (c *wafregionalSQLInjectionMatchSets) List(opts v1.ListOptions) (result *v1alpha1.WafregionalSQLInjectionMatchSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.WafregionalSQLInjectionMatchSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalsqlinjectionmatchsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested wafregionalSQLInjectionMatchSets.
func (c *wafregionalSQLInjectionMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("wafregionalsqlinjectionmatchsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a wafregionalSQLInjectionMatchSet and creates it.  Returns the server's representation of the wafregionalSQLInjectionMatchSet, and an error, if there is any.
func (c *wafregionalSQLInjectionMatchSets) Create(wafregionalSQLInjectionMatchSet *v1alpha1.WafregionalSQLInjectionMatchSet) (result *v1alpha1.WafregionalSQLInjectionMatchSet, err error) {
	result = &v1alpha1.WafregionalSQLInjectionMatchSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("wafregionalsqlinjectionmatchsets").
		Body(wafregionalSQLInjectionMatchSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a wafregionalSQLInjectionMatchSet and updates it. Returns the server's representation of the wafregionalSQLInjectionMatchSet, and an error, if there is any.
func (c *wafregionalSQLInjectionMatchSets) Update(wafregionalSQLInjectionMatchSet *v1alpha1.WafregionalSQLInjectionMatchSet) (result *v1alpha1.WafregionalSQLInjectionMatchSet, err error) {
	result = &v1alpha1.WafregionalSQLInjectionMatchSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafregionalsqlinjectionmatchsets").
		Name(wafregionalSQLInjectionMatchSet.Name).
		Body(wafregionalSQLInjectionMatchSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *wafregionalSQLInjectionMatchSets) UpdateStatus(wafregionalSQLInjectionMatchSet *v1alpha1.WafregionalSQLInjectionMatchSet) (result *v1alpha1.WafregionalSQLInjectionMatchSet, err error) {
	result = &v1alpha1.WafregionalSQLInjectionMatchSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wafregionalsqlinjectionmatchsets").
		Name(wafregionalSQLInjectionMatchSet.Name).
		SubResource("status").
		Body(wafregionalSQLInjectionMatchSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the wafregionalSQLInjectionMatchSet and deletes it. Returns an error if one occurs.
func (c *wafregionalSQLInjectionMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafregionalsqlinjectionmatchsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *wafregionalSQLInjectionMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wafregionalsqlinjectionmatchsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched wafregionalSQLInjectionMatchSet.
func (c *wafregionalSQLInjectionMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WafregionalSQLInjectionMatchSet, err error) {
	result = &v1alpha1.WafregionalSQLInjectionMatchSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("wafregionalsqlinjectionmatchsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
