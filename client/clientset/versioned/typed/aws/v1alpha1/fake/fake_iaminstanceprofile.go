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

// FakeIamInstanceProfiles implements IamInstanceProfileInterface
type FakeIamInstanceProfiles struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var iaminstanceprofilesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iaminstanceprofiles"}

var iaminstanceprofilesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IamInstanceProfile"}

// Get takes name of the iamInstanceProfile, and returns the corresponding iamInstanceProfile object, and an error if there is any.
func (c *FakeIamInstanceProfiles) Get(name string, options v1.GetOptions) (result *v1alpha1.IamInstanceProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iaminstanceprofilesResource, c.ns, name), &v1alpha1.IamInstanceProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamInstanceProfile), err
}

// List takes label and field selectors, and returns the list of IamInstanceProfiles that match those selectors.
func (c *FakeIamInstanceProfiles) List(opts v1.ListOptions) (result *v1alpha1.IamInstanceProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iaminstanceprofilesResource, iaminstanceprofilesKind, c.ns, opts), &v1alpha1.IamInstanceProfileList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamInstanceProfileList{ListMeta: obj.(*v1alpha1.IamInstanceProfileList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamInstanceProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamInstanceProfiles.
func (c *FakeIamInstanceProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iaminstanceprofilesResource, c.ns, opts))

}

// Create takes the representation of a iamInstanceProfile and creates it.  Returns the server's representation of the iamInstanceProfile, and an error, if there is any.
func (c *FakeIamInstanceProfiles) Create(iamInstanceProfile *v1alpha1.IamInstanceProfile) (result *v1alpha1.IamInstanceProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iaminstanceprofilesResource, c.ns, iamInstanceProfile), &v1alpha1.IamInstanceProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamInstanceProfile), err
}

// Update takes the representation of a iamInstanceProfile and updates it. Returns the server's representation of the iamInstanceProfile, and an error, if there is any.
func (c *FakeIamInstanceProfiles) Update(iamInstanceProfile *v1alpha1.IamInstanceProfile) (result *v1alpha1.IamInstanceProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iaminstanceprofilesResource, c.ns, iamInstanceProfile), &v1alpha1.IamInstanceProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamInstanceProfile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamInstanceProfiles) UpdateStatus(iamInstanceProfile *v1alpha1.IamInstanceProfile) (*v1alpha1.IamInstanceProfile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iaminstanceprofilesResource, "status", c.ns, iamInstanceProfile), &v1alpha1.IamInstanceProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamInstanceProfile), err
}

// Delete takes name of the iamInstanceProfile and deletes it. Returns an error if one occurs.
func (c *FakeIamInstanceProfiles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iaminstanceprofilesResource, c.ns, name), &v1alpha1.IamInstanceProfile{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamInstanceProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iaminstanceprofilesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamInstanceProfileList{})
	return err
}

// Patch applies the patch and returns the patched iamInstanceProfile.
func (c *FakeIamInstanceProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamInstanceProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iaminstanceprofilesResource, c.ns, name, pt, data, subresources...), &v1alpha1.IamInstanceProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamInstanceProfile), err
}
