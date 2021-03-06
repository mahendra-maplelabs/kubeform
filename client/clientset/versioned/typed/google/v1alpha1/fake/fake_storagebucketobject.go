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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStorageBucketObjects implements StorageBucketObjectInterface
type FakeStorageBucketObjects struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var storagebucketobjectsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "storagebucketobjects"}

var storagebucketobjectsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "StorageBucketObject"}

// Get takes name of the storageBucketObject, and returns the corresponding storageBucketObject object, and an error if there is any.
func (c *FakeStorageBucketObjects) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageBucketObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagebucketobjectsResource, c.ns, name), &v1alpha1.StorageBucketObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketObject), err
}

// List takes label and field selectors, and returns the list of StorageBucketObjects that match those selectors.
func (c *FakeStorageBucketObjects) List(opts v1.ListOptions) (result *v1alpha1.StorageBucketObjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagebucketobjectsResource, storagebucketobjectsKind, c.ns, opts), &v1alpha1.StorageBucketObjectList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageBucketObjectList{ListMeta: obj.(*v1alpha1.StorageBucketObjectList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageBucketObjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageBucketObjects.
func (c *FakeStorageBucketObjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagebucketobjectsResource, c.ns, opts))

}

// Create takes the representation of a storageBucketObject and creates it.  Returns the server's representation of the storageBucketObject, and an error, if there is any.
func (c *FakeStorageBucketObjects) Create(storageBucketObject *v1alpha1.StorageBucketObject) (result *v1alpha1.StorageBucketObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagebucketobjectsResource, c.ns, storageBucketObject), &v1alpha1.StorageBucketObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketObject), err
}

// Update takes the representation of a storageBucketObject and updates it. Returns the server's representation of the storageBucketObject, and an error, if there is any.
func (c *FakeStorageBucketObjects) Update(storageBucketObject *v1alpha1.StorageBucketObject) (result *v1alpha1.StorageBucketObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagebucketobjectsResource, c.ns, storageBucketObject), &v1alpha1.StorageBucketObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketObject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageBucketObjects) UpdateStatus(storageBucketObject *v1alpha1.StorageBucketObject) (*v1alpha1.StorageBucketObject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagebucketobjectsResource, "status", c.ns, storageBucketObject), &v1alpha1.StorageBucketObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketObject), err
}

// Delete takes name of the storageBucketObject and deletes it. Returns an error if one occurs.
func (c *FakeStorageBucketObjects) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagebucketobjectsResource, c.ns, name), &v1alpha1.StorageBucketObject{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageBucketObjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagebucketobjectsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageBucketObjectList{})
	return err
}

// Patch applies the patch and returns the patched storageBucketObject.
func (c *FakeStorageBucketObjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageBucketObject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagebucketobjectsResource, c.ns, name, pt, data, subresources...), &v1alpha1.StorageBucketObject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageBucketObject), err
}
