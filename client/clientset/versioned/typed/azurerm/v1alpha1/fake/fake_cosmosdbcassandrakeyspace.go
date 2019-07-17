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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
)

// FakeCosmosdbCassandraKeyspaces implements CosmosdbCassandraKeyspaceInterface
type FakeCosmosdbCassandraKeyspaces struct {
	Fake *FakeAzurermV1alpha1
}

var cosmosdbcassandrakeyspacesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "cosmosdbcassandrakeyspaces"}

var cosmosdbcassandrakeyspacesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "CosmosdbCassandraKeyspace"}

// Get takes name of the cosmosdbCassandraKeyspace, and returns the corresponding cosmosdbCassandraKeyspace object, and an error if there is any.
func (c *FakeCosmosdbCassandraKeyspaces) Get(name string, options v1.GetOptions) (result *v1alpha1.CosmosdbCassandraKeyspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cosmosdbcassandrakeyspacesResource, name), &v1alpha1.CosmosdbCassandraKeyspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbCassandraKeyspace), err
}

// List takes label and field selectors, and returns the list of CosmosdbCassandraKeyspaces that match those selectors.
func (c *FakeCosmosdbCassandraKeyspaces) List(opts v1.ListOptions) (result *v1alpha1.CosmosdbCassandraKeyspaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cosmosdbcassandrakeyspacesResource, cosmosdbcassandrakeyspacesKind, opts), &v1alpha1.CosmosdbCassandraKeyspaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CosmosdbCassandraKeyspaceList{ListMeta: obj.(*v1alpha1.CosmosdbCassandraKeyspaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.CosmosdbCassandraKeyspaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cosmosdbCassandraKeyspaces.
func (c *FakeCosmosdbCassandraKeyspaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cosmosdbcassandrakeyspacesResource, opts))
}

// Create takes the representation of a cosmosdbCassandraKeyspace and creates it.  Returns the server's representation of the cosmosdbCassandraKeyspace, and an error, if there is any.
func (c *FakeCosmosdbCassandraKeyspaces) Create(cosmosdbCassandraKeyspace *v1alpha1.CosmosdbCassandraKeyspace) (result *v1alpha1.CosmosdbCassandraKeyspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cosmosdbcassandrakeyspacesResource, cosmosdbCassandraKeyspace), &v1alpha1.CosmosdbCassandraKeyspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbCassandraKeyspace), err
}

// Update takes the representation of a cosmosdbCassandraKeyspace and updates it. Returns the server's representation of the cosmosdbCassandraKeyspace, and an error, if there is any.
func (c *FakeCosmosdbCassandraKeyspaces) Update(cosmosdbCassandraKeyspace *v1alpha1.CosmosdbCassandraKeyspace) (result *v1alpha1.CosmosdbCassandraKeyspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cosmosdbcassandrakeyspacesResource, cosmosdbCassandraKeyspace), &v1alpha1.CosmosdbCassandraKeyspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbCassandraKeyspace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCosmosdbCassandraKeyspaces) UpdateStatus(cosmosdbCassandraKeyspace *v1alpha1.CosmosdbCassandraKeyspace) (*v1alpha1.CosmosdbCassandraKeyspace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(cosmosdbcassandrakeyspacesResource, "status", cosmosdbCassandraKeyspace), &v1alpha1.CosmosdbCassandraKeyspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbCassandraKeyspace), err
}

// Delete takes name of the cosmosdbCassandraKeyspace and deletes it. Returns an error if one occurs.
func (c *FakeCosmosdbCassandraKeyspaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(cosmosdbcassandrakeyspacesResource, name), &v1alpha1.CosmosdbCassandraKeyspace{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCosmosdbCassandraKeyspaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(cosmosdbcassandrakeyspacesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CosmosdbCassandraKeyspaceList{})
	return err
}

// Patch applies the patch and returns the patched cosmosdbCassandraKeyspace.
func (c *FakeCosmosdbCassandraKeyspaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CosmosdbCassandraKeyspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cosmosdbcassandrakeyspacesResource, name, pt, data, subresources...), &v1alpha1.CosmosdbCassandraKeyspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CosmosdbCassandraKeyspace), err
}