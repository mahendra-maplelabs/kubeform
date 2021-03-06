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

// FakeVirtualNetworkPeerings implements VirtualNetworkPeeringInterface
type FakeVirtualNetworkPeerings struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var virtualnetworkpeeringsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "virtualnetworkpeerings"}

var virtualnetworkpeeringsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "VirtualNetworkPeering"}

// Get takes name of the virtualNetworkPeering, and returns the corresponding virtualNetworkPeering object, and an error if there is any.
func (c *FakeVirtualNetworkPeerings) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualNetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualnetworkpeeringsResource, c.ns, name), &v1alpha1.VirtualNetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualNetworkPeering), err
}

// List takes label and field selectors, and returns the list of VirtualNetworkPeerings that match those selectors.
func (c *FakeVirtualNetworkPeerings) List(opts v1.ListOptions) (result *v1alpha1.VirtualNetworkPeeringList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualnetworkpeeringsResource, virtualnetworkpeeringsKind, c.ns, opts), &v1alpha1.VirtualNetworkPeeringList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualNetworkPeeringList{ListMeta: obj.(*v1alpha1.VirtualNetworkPeeringList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualNetworkPeeringList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualNetworkPeerings.
func (c *FakeVirtualNetworkPeerings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualnetworkpeeringsResource, c.ns, opts))

}

// Create takes the representation of a virtualNetworkPeering and creates it.  Returns the server's representation of the virtualNetworkPeering, and an error, if there is any.
func (c *FakeVirtualNetworkPeerings) Create(virtualNetworkPeering *v1alpha1.VirtualNetworkPeering) (result *v1alpha1.VirtualNetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualnetworkpeeringsResource, c.ns, virtualNetworkPeering), &v1alpha1.VirtualNetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualNetworkPeering), err
}

// Update takes the representation of a virtualNetworkPeering and updates it. Returns the server's representation of the virtualNetworkPeering, and an error, if there is any.
func (c *FakeVirtualNetworkPeerings) Update(virtualNetworkPeering *v1alpha1.VirtualNetworkPeering) (result *v1alpha1.VirtualNetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualnetworkpeeringsResource, c.ns, virtualNetworkPeering), &v1alpha1.VirtualNetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualNetworkPeering), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualNetworkPeerings) UpdateStatus(virtualNetworkPeering *v1alpha1.VirtualNetworkPeering) (*v1alpha1.VirtualNetworkPeering, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualnetworkpeeringsResource, "status", c.ns, virtualNetworkPeering), &v1alpha1.VirtualNetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualNetworkPeering), err
}

// Delete takes name of the virtualNetworkPeering and deletes it. Returns an error if one occurs.
func (c *FakeVirtualNetworkPeerings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualnetworkpeeringsResource, c.ns, name), &v1alpha1.VirtualNetworkPeering{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualNetworkPeerings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualnetworkpeeringsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualNetworkPeeringList{})
	return err
}

// Patch applies the patch and returns the patched virtualNetworkPeering.
func (c *FakeVirtualNetworkPeerings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualNetworkPeering, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualnetworkpeeringsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VirtualNetworkPeering{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualNetworkPeering), err
}
