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

// CognitoResourceServersGetter has a method to return a CognitoResourceServerInterface.
// A group's client should implement this interface.
type CognitoResourceServersGetter interface {
	CognitoResourceServers(namespace string) CognitoResourceServerInterface
}

// CognitoResourceServerInterface has methods to work with CognitoResourceServer resources.
type CognitoResourceServerInterface interface {
	Create(*v1alpha1.CognitoResourceServer) (*v1alpha1.CognitoResourceServer, error)
	Update(*v1alpha1.CognitoResourceServer) (*v1alpha1.CognitoResourceServer, error)
	UpdateStatus(*v1alpha1.CognitoResourceServer) (*v1alpha1.CognitoResourceServer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CognitoResourceServer, error)
	List(opts v1.ListOptions) (*v1alpha1.CognitoResourceServerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CognitoResourceServer, err error)
	CognitoResourceServerExpansion
}

// cognitoResourceServers implements CognitoResourceServerInterface
type cognitoResourceServers struct {
	client rest.Interface
	ns     string
}

// newCognitoResourceServers returns a CognitoResourceServers
func newCognitoResourceServers(c *AwsV1alpha1Client, namespace string) *cognitoResourceServers {
	return &cognitoResourceServers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cognitoResourceServer, and returns the corresponding cognitoResourceServer object, and an error if there is any.
func (c *cognitoResourceServers) Get(name string, options v1.GetOptions) (result *v1alpha1.CognitoResourceServer, err error) {
	result = &v1alpha1.CognitoResourceServer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cognitoresourceservers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CognitoResourceServers that match those selectors.
func (c *cognitoResourceServers) List(opts v1.ListOptions) (result *v1alpha1.CognitoResourceServerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CognitoResourceServerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cognitoresourceservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cognitoResourceServers.
func (c *cognitoResourceServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cognitoresourceservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cognitoResourceServer and creates it.  Returns the server's representation of the cognitoResourceServer, and an error, if there is any.
func (c *cognitoResourceServers) Create(cognitoResourceServer *v1alpha1.CognitoResourceServer) (result *v1alpha1.CognitoResourceServer, err error) {
	result = &v1alpha1.CognitoResourceServer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cognitoresourceservers").
		Body(cognitoResourceServer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cognitoResourceServer and updates it. Returns the server's representation of the cognitoResourceServer, and an error, if there is any.
func (c *cognitoResourceServers) Update(cognitoResourceServer *v1alpha1.CognitoResourceServer) (result *v1alpha1.CognitoResourceServer, err error) {
	result = &v1alpha1.CognitoResourceServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cognitoresourceservers").
		Name(cognitoResourceServer.Name).
		Body(cognitoResourceServer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cognitoResourceServers) UpdateStatus(cognitoResourceServer *v1alpha1.CognitoResourceServer) (result *v1alpha1.CognitoResourceServer, err error) {
	result = &v1alpha1.CognitoResourceServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cognitoresourceservers").
		Name(cognitoResourceServer.Name).
		SubResource("status").
		Body(cognitoResourceServer).
		Do().
		Into(result)
	return
}

// Delete takes name of the cognitoResourceServer and deletes it. Returns an error if one occurs.
func (c *cognitoResourceServers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cognitoresourceservers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cognitoResourceServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cognitoresourceservers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cognitoResourceServer.
func (c *cognitoResourceServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CognitoResourceServer, err error) {
	result = &v1alpha1.CognitoResourceServer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cognitoresourceservers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
