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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDmsReplicationSubnetGroups implements DmsReplicationSubnetGroupInterface
type FakeDmsReplicationSubnetGroups struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dmsreplicationsubnetgroupsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dmsreplicationsubnetgroups"}

var dmsreplicationsubnetgroupsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DmsReplicationSubnetGroup"}

// Get takes name of the dmsReplicationSubnetGroup, and returns the corresponding dmsReplicationSubnetGroup object, and an error if there is any.
func (c *FakeDmsReplicationSubnetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.DmsReplicationSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dmsreplicationsubnetgroupsResource, c.ns, name), &v1alpha1.DmsReplicationSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsReplicationSubnetGroup), err
}

// List takes label and field selectors, and returns the list of DmsReplicationSubnetGroups that match those selectors.
func (c *FakeDmsReplicationSubnetGroups) List(opts v1.ListOptions) (result *v1alpha1.DmsReplicationSubnetGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dmsreplicationsubnetgroupsResource, dmsreplicationsubnetgroupsKind, c.ns, opts), &v1alpha1.DmsReplicationSubnetGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DmsReplicationSubnetGroupList{ListMeta: obj.(*v1alpha1.DmsReplicationSubnetGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.DmsReplicationSubnetGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dmsReplicationSubnetGroups.
func (c *FakeDmsReplicationSubnetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dmsreplicationsubnetgroupsResource, c.ns, opts))

}

// Create takes the representation of a dmsReplicationSubnetGroup and creates it.  Returns the server's representation of the dmsReplicationSubnetGroup, and an error, if there is any.
func (c *FakeDmsReplicationSubnetGroups) Create(dmsReplicationSubnetGroup *v1alpha1.DmsReplicationSubnetGroup) (result *v1alpha1.DmsReplicationSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dmsreplicationsubnetgroupsResource, c.ns, dmsReplicationSubnetGroup), &v1alpha1.DmsReplicationSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsReplicationSubnetGroup), err
}

// Update takes the representation of a dmsReplicationSubnetGroup and updates it. Returns the server's representation of the dmsReplicationSubnetGroup, and an error, if there is any.
func (c *FakeDmsReplicationSubnetGroups) Update(dmsReplicationSubnetGroup *v1alpha1.DmsReplicationSubnetGroup) (result *v1alpha1.DmsReplicationSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dmsreplicationsubnetgroupsResource, c.ns, dmsReplicationSubnetGroup), &v1alpha1.DmsReplicationSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsReplicationSubnetGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDmsReplicationSubnetGroups) UpdateStatus(dmsReplicationSubnetGroup *v1alpha1.DmsReplicationSubnetGroup) (*v1alpha1.DmsReplicationSubnetGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dmsreplicationsubnetgroupsResource, "status", c.ns, dmsReplicationSubnetGroup), &v1alpha1.DmsReplicationSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsReplicationSubnetGroup), err
}

// Delete takes name of the dmsReplicationSubnetGroup and deletes it. Returns an error if one occurs.
func (c *FakeDmsReplicationSubnetGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dmsreplicationsubnetgroupsResource, c.ns, name), &v1alpha1.DmsReplicationSubnetGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDmsReplicationSubnetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dmsreplicationsubnetgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DmsReplicationSubnetGroupList{})
	return err
}

// Patch applies the patch and returns the patched dmsReplicationSubnetGroup.
func (c *FakeDmsReplicationSubnetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DmsReplicationSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dmsreplicationsubnetgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DmsReplicationSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DmsReplicationSubnetGroup), err
}
