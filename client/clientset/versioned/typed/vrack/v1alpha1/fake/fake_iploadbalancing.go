/*
Copyright AppsCode Inc. and Contributors

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
	"context"

	v1alpha1 "kubeform.dev/provider-ovh-api/apis/vrack/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIploadbalancings implements IploadbalancingInterface
type FakeIploadbalancings struct {
	Fake *FakeVrackV1alpha1
	ns   string
}

var iploadbalancingsResource = schema.GroupVersionResource{Group: "vrack.ovh.kubeform.com", Version: "v1alpha1", Resource: "iploadbalancings"}

var iploadbalancingsKind = schema.GroupVersionKind{Group: "vrack.ovh.kubeform.com", Version: "v1alpha1", Kind: "Iploadbalancing"}

// Get takes name of the iploadbalancing, and returns the corresponding iploadbalancing object, and an error if there is any.
func (c *FakeIploadbalancings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Iploadbalancing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iploadbalancingsResource, c.ns, name), &v1alpha1.Iploadbalancing{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Iploadbalancing), err
}

// List takes label and field selectors, and returns the list of Iploadbalancings that match those selectors.
func (c *FakeIploadbalancings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IploadbalancingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iploadbalancingsResource, iploadbalancingsKind, c.ns, opts), &v1alpha1.IploadbalancingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IploadbalancingList{ListMeta: obj.(*v1alpha1.IploadbalancingList).ListMeta}
	for _, item := range obj.(*v1alpha1.IploadbalancingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iploadbalancings.
func (c *FakeIploadbalancings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iploadbalancingsResource, c.ns, opts))

}

// Create takes the representation of a iploadbalancing and creates it.  Returns the server's representation of the iploadbalancing, and an error, if there is any.
func (c *FakeIploadbalancings) Create(ctx context.Context, iploadbalancing *v1alpha1.Iploadbalancing, opts v1.CreateOptions) (result *v1alpha1.Iploadbalancing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iploadbalancingsResource, c.ns, iploadbalancing), &v1alpha1.Iploadbalancing{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Iploadbalancing), err
}

// Update takes the representation of a iploadbalancing and updates it. Returns the server's representation of the iploadbalancing, and an error, if there is any.
func (c *FakeIploadbalancings) Update(ctx context.Context, iploadbalancing *v1alpha1.Iploadbalancing, opts v1.UpdateOptions) (result *v1alpha1.Iploadbalancing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iploadbalancingsResource, c.ns, iploadbalancing), &v1alpha1.Iploadbalancing{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Iploadbalancing), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIploadbalancings) UpdateStatus(ctx context.Context, iploadbalancing *v1alpha1.Iploadbalancing, opts v1.UpdateOptions) (*v1alpha1.Iploadbalancing, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iploadbalancingsResource, "status", c.ns, iploadbalancing), &v1alpha1.Iploadbalancing{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Iploadbalancing), err
}

// Delete takes name of the iploadbalancing and deletes it. Returns an error if one occurs.
func (c *FakeIploadbalancings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iploadbalancingsResource, c.ns, name), &v1alpha1.Iploadbalancing{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIploadbalancings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iploadbalancingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IploadbalancingList{})
	return err
}

// Patch applies the patch and returns the patched iploadbalancing.
func (c *FakeIploadbalancings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Iploadbalancing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iploadbalancingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Iploadbalancing{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Iploadbalancing), err
}
