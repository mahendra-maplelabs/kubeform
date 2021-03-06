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

// FakePinpointBaiduChannels implements PinpointBaiduChannelInterface
type FakePinpointBaiduChannels struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var pinpointbaiduchannelsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "pinpointbaiduchannels"}

var pinpointbaiduchannelsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "PinpointBaiduChannel"}

// Get takes name of the pinpointBaiduChannel, and returns the corresponding pinpointBaiduChannel object, and an error if there is any.
func (c *FakePinpointBaiduChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.PinpointBaiduChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pinpointbaiduchannelsResource, c.ns, name), &v1alpha1.PinpointBaiduChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointBaiduChannel), err
}

// List takes label and field selectors, and returns the list of PinpointBaiduChannels that match those selectors.
func (c *FakePinpointBaiduChannels) List(opts v1.ListOptions) (result *v1alpha1.PinpointBaiduChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pinpointbaiduchannelsResource, pinpointbaiduchannelsKind, c.ns, opts), &v1alpha1.PinpointBaiduChannelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PinpointBaiduChannelList{ListMeta: obj.(*v1alpha1.PinpointBaiduChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.PinpointBaiduChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pinpointBaiduChannels.
func (c *FakePinpointBaiduChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pinpointbaiduchannelsResource, c.ns, opts))

}

// Create takes the representation of a pinpointBaiduChannel and creates it.  Returns the server's representation of the pinpointBaiduChannel, and an error, if there is any.
func (c *FakePinpointBaiduChannels) Create(pinpointBaiduChannel *v1alpha1.PinpointBaiduChannel) (result *v1alpha1.PinpointBaiduChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pinpointbaiduchannelsResource, c.ns, pinpointBaiduChannel), &v1alpha1.PinpointBaiduChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointBaiduChannel), err
}

// Update takes the representation of a pinpointBaiduChannel and updates it. Returns the server's representation of the pinpointBaiduChannel, and an error, if there is any.
func (c *FakePinpointBaiduChannels) Update(pinpointBaiduChannel *v1alpha1.PinpointBaiduChannel) (result *v1alpha1.PinpointBaiduChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pinpointbaiduchannelsResource, c.ns, pinpointBaiduChannel), &v1alpha1.PinpointBaiduChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointBaiduChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePinpointBaiduChannels) UpdateStatus(pinpointBaiduChannel *v1alpha1.PinpointBaiduChannel) (*v1alpha1.PinpointBaiduChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pinpointbaiduchannelsResource, "status", c.ns, pinpointBaiduChannel), &v1alpha1.PinpointBaiduChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointBaiduChannel), err
}

// Delete takes name of the pinpointBaiduChannel and deletes it. Returns an error if one occurs.
func (c *FakePinpointBaiduChannels) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pinpointbaiduchannelsResource, c.ns, name), &v1alpha1.PinpointBaiduChannel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePinpointBaiduChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pinpointbaiduchannelsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PinpointBaiduChannelList{})
	return err
}

// Patch applies the patch and returns the patched pinpointBaiduChannel.
func (c *FakePinpointBaiduChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PinpointBaiduChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pinpointbaiduchannelsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PinpointBaiduChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointBaiduChannel), err
}
