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

package fake

import (
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMysqlServers implements MysqlServerInterface
type FakeMysqlServers struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var mysqlserversResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "mysqlservers"}

var mysqlserversKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MysqlServer"}

// Get takes name of the mysqlServer, and returns the corresponding mysqlServer object, and an error if there is any.
func (c *FakeMysqlServers) Get(name string, options v1.GetOptions) (result *v1alpha1.MysqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mysqlserversResource, c.ns, name), &v1alpha1.MysqlServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlServer), err
}

// List takes label and field selectors, and returns the list of MysqlServers that match those selectors.
func (c *FakeMysqlServers) List(opts v1.ListOptions) (result *v1alpha1.MysqlServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mysqlserversResource, mysqlserversKind, c.ns, opts), &v1alpha1.MysqlServerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MysqlServerList{ListMeta: obj.(*v1alpha1.MysqlServerList).ListMeta}
	for _, item := range obj.(*v1alpha1.MysqlServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mysqlServers.
func (c *FakeMysqlServers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mysqlserversResource, c.ns, opts))

}

// Create takes the representation of a mysqlServer and creates it.  Returns the server's representation of the mysqlServer, and an error, if there is any.
func (c *FakeMysqlServers) Create(mysqlServer *v1alpha1.MysqlServer) (result *v1alpha1.MysqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mysqlserversResource, c.ns, mysqlServer), &v1alpha1.MysqlServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlServer), err
}

// Update takes the representation of a mysqlServer and updates it. Returns the server's representation of the mysqlServer, and an error, if there is any.
func (c *FakeMysqlServers) Update(mysqlServer *v1alpha1.MysqlServer) (result *v1alpha1.MysqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mysqlserversResource, c.ns, mysqlServer), &v1alpha1.MysqlServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMysqlServers) UpdateStatus(mysqlServer *v1alpha1.MysqlServer) (*v1alpha1.MysqlServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mysqlserversResource, "status", c.ns, mysqlServer), &v1alpha1.MysqlServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlServer), err
}

// Delete takes name of the mysqlServer and deletes it. Returns an error if one occurs.
func (c *FakeMysqlServers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mysqlserversResource, c.ns, name), &v1alpha1.MysqlServer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMysqlServers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mysqlserversResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MysqlServerList{})
	return err
}

// Patch applies the patch and returns the patched mysqlServer.
func (c *FakeMysqlServers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MysqlServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mysqlserversResource, c.ns, name, pt, data, subresources...), &v1alpha1.MysqlServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlServer), err
}
