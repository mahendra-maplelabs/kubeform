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

// ApiManagementGroupUsersGetter has a method to return a ApiManagementGroupUserInterface.
// A group's client should implement this interface.
type ApiManagementGroupUsersGetter interface {
	ApiManagementGroupUsers(namespace string) ApiManagementGroupUserInterface
}

// ApiManagementGroupUserInterface has methods to work with ApiManagementGroupUser resources.
type ApiManagementGroupUserInterface interface {
	Create(*v1alpha1.ApiManagementGroupUser) (*v1alpha1.ApiManagementGroupUser, error)
	Update(*v1alpha1.ApiManagementGroupUser) (*v1alpha1.ApiManagementGroupUser, error)
	UpdateStatus(*v1alpha1.ApiManagementGroupUser) (*v1alpha1.ApiManagementGroupUser, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiManagementGroupUser, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiManagementGroupUserList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementGroupUser, err error)
	ApiManagementGroupUserExpansion
}

// apiManagementGroupUsers implements ApiManagementGroupUserInterface
type apiManagementGroupUsers struct {
	client rest.Interface
	ns     string
}

// newApiManagementGroupUsers returns a ApiManagementGroupUsers
func newApiManagementGroupUsers(c *AzurermV1alpha1Client, namespace string) *apiManagementGroupUsers {
	return &apiManagementGroupUsers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiManagementGroupUser, and returns the corresponding apiManagementGroupUser object, and an error if there is any.
func (c *apiManagementGroupUsers) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementGroupUser, err error) {
	result = &v1alpha1.ApiManagementGroupUser{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementgroupusers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiManagementGroupUsers that match those selectors.
func (c *apiManagementGroupUsers) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementGroupUserList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiManagementGroupUserList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementgroupusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiManagementGroupUsers.
func (c *apiManagementGroupUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementgroupusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiManagementGroupUser and creates it.  Returns the server's representation of the apiManagementGroupUser, and an error, if there is any.
func (c *apiManagementGroupUsers) Create(apiManagementGroupUser *v1alpha1.ApiManagementGroupUser) (result *v1alpha1.ApiManagementGroupUser, err error) {
	result = &v1alpha1.ApiManagementGroupUser{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apimanagementgroupusers").
		Body(apiManagementGroupUser).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiManagementGroupUser and updates it. Returns the server's representation of the apiManagementGroupUser, and an error, if there is any.
func (c *apiManagementGroupUsers) Update(apiManagementGroupUser *v1alpha1.ApiManagementGroupUser) (result *v1alpha1.ApiManagementGroupUser, err error) {
	result = &v1alpha1.ApiManagementGroupUser{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementgroupusers").
		Name(apiManagementGroupUser.Name).
		Body(apiManagementGroupUser).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiManagementGroupUsers) UpdateStatus(apiManagementGroupUser *v1alpha1.ApiManagementGroupUser) (result *v1alpha1.ApiManagementGroupUser, err error) {
	result = &v1alpha1.ApiManagementGroupUser{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementgroupusers").
		Name(apiManagementGroupUser.Name).
		SubResource("status").
		Body(apiManagementGroupUser).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiManagementGroupUser and deletes it. Returns an error if one occurs.
func (c *apiManagementGroupUsers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementgroupusers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiManagementGroupUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementgroupusers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiManagementGroupUser.
func (c *apiManagementGroupUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementGroupUser, err error) {
	result = &v1alpha1.ApiManagementGroupUser{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apimanagementgroupusers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
