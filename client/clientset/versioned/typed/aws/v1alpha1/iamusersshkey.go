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

// IamUserSSHKeysGetter has a method to return a IamUserSSHKeyInterface.
// A group's client should implement this interface.
type IamUserSSHKeysGetter interface {
	IamUserSSHKeys(namespace string) IamUserSSHKeyInterface
}

// IamUserSSHKeyInterface has methods to work with IamUserSSHKey resources.
type IamUserSSHKeyInterface interface {
	Create(*v1alpha1.IamUserSSHKey) (*v1alpha1.IamUserSSHKey, error)
	Update(*v1alpha1.IamUserSSHKey) (*v1alpha1.IamUserSSHKey, error)
	UpdateStatus(*v1alpha1.IamUserSSHKey) (*v1alpha1.IamUserSSHKey, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IamUserSSHKey, error)
	List(opts v1.ListOptions) (*v1alpha1.IamUserSSHKeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamUserSSHKey, err error)
	IamUserSSHKeyExpansion
}

// iamUserSSHKeys implements IamUserSSHKeyInterface
type iamUserSSHKeys struct {
	client rest.Interface
	ns     string
}

// newIamUserSSHKeys returns a IamUserSSHKeys
func newIamUserSSHKeys(c *AwsV1alpha1Client, namespace string) *iamUserSSHKeys {
	return &iamUserSSHKeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iamUserSSHKey, and returns the corresponding iamUserSSHKey object, and an error if there is any.
func (c *iamUserSSHKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.IamUserSSHKey, err error) {
	result = &v1alpha1.IamUserSSHKey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamusersshkeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IamUserSSHKeys that match those selectors.
func (c *iamUserSSHKeys) List(opts v1.ListOptions) (result *v1alpha1.IamUserSSHKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IamUserSSHKeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamusersshkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iamUserSSHKeys.
func (c *iamUserSSHKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iamusersshkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iamUserSSHKey and creates it.  Returns the server's representation of the iamUserSSHKey, and an error, if there is any.
func (c *iamUserSSHKeys) Create(iamUserSSHKey *v1alpha1.IamUserSSHKey) (result *v1alpha1.IamUserSSHKey, err error) {
	result = &v1alpha1.IamUserSSHKey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iamusersshkeys").
		Body(iamUserSSHKey).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iamUserSSHKey and updates it. Returns the server's representation of the iamUserSSHKey, and an error, if there is any.
func (c *iamUserSSHKeys) Update(iamUserSSHKey *v1alpha1.IamUserSSHKey) (result *v1alpha1.IamUserSSHKey, err error) {
	result = &v1alpha1.IamUserSSHKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamusersshkeys").
		Name(iamUserSSHKey.Name).
		Body(iamUserSSHKey).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iamUserSSHKeys) UpdateStatus(iamUserSSHKey *v1alpha1.IamUserSSHKey) (result *v1alpha1.IamUserSSHKey, err error) {
	result = &v1alpha1.IamUserSSHKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamusersshkeys").
		Name(iamUserSSHKey.Name).
		SubResource("status").
		Body(iamUserSSHKey).
		Do().
		Into(result)
	return
}

// Delete takes name of the iamUserSSHKey and deletes it. Returns an error if one occurs.
func (c *iamUserSSHKeys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamusersshkeys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iamUserSSHKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamusersshkeys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iamUserSSHKey.
func (c *iamUserSSHKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamUserSSHKey, err error) {
	result = &v1alpha1.IamUserSSHKey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iamusersshkeys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
