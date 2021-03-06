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

// FakeOpsworksApplications implements OpsworksApplicationInterface
type FakeOpsworksApplications struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var opsworksapplicationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "opsworksapplications"}

var opsworksapplicationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "OpsworksApplication"}

// Get takes name of the opsworksApplication, and returns the corresponding opsworksApplication object, and an error if there is any.
func (c *FakeOpsworksApplications) Get(name string, options v1.GetOptions) (result *v1alpha1.OpsworksApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(opsworksapplicationsResource, c.ns, name), &v1alpha1.OpsworksApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksApplication), err
}

// List takes label and field selectors, and returns the list of OpsworksApplications that match those selectors.
func (c *FakeOpsworksApplications) List(opts v1.ListOptions) (result *v1alpha1.OpsworksApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(opsworksapplicationsResource, opsworksapplicationsKind, c.ns, opts), &v1alpha1.OpsworksApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OpsworksApplicationList{ListMeta: obj.(*v1alpha1.OpsworksApplicationList).ListMeta}
	for _, item := range obj.(*v1alpha1.OpsworksApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested opsworksApplications.
func (c *FakeOpsworksApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(opsworksapplicationsResource, c.ns, opts))

}

// Create takes the representation of a opsworksApplication and creates it.  Returns the server's representation of the opsworksApplication, and an error, if there is any.
func (c *FakeOpsworksApplications) Create(opsworksApplication *v1alpha1.OpsworksApplication) (result *v1alpha1.OpsworksApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(opsworksapplicationsResource, c.ns, opsworksApplication), &v1alpha1.OpsworksApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksApplication), err
}

// Update takes the representation of a opsworksApplication and updates it. Returns the server's representation of the opsworksApplication, and an error, if there is any.
func (c *FakeOpsworksApplications) Update(opsworksApplication *v1alpha1.OpsworksApplication) (result *v1alpha1.OpsworksApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(opsworksapplicationsResource, c.ns, opsworksApplication), &v1alpha1.OpsworksApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksApplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpsworksApplications) UpdateStatus(opsworksApplication *v1alpha1.OpsworksApplication) (*v1alpha1.OpsworksApplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(opsworksapplicationsResource, "status", c.ns, opsworksApplication), &v1alpha1.OpsworksApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksApplication), err
}

// Delete takes name of the opsworksApplication and deletes it. Returns an error if one occurs.
func (c *FakeOpsworksApplications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(opsworksapplicationsResource, c.ns, name), &v1alpha1.OpsworksApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpsworksApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(opsworksapplicationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OpsworksApplicationList{})
	return err
}

// Patch applies the patch and returns the patched opsworksApplication.
func (c *FakeOpsworksApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(opsworksapplicationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OpsworksApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksApplication), err
}
