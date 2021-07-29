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

	v1alpha1 "kubeform.dev/provider-ovh-api/apis/domain/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeZoneRedirections implements ZoneRedirectionInterface
type FakeZoneRedirections struct {
	Fake *FakeDomainV1alpha1
	ns   string
}

var zoneredirectionsResource = schema.GroupVersionResource{Group: "domain.ovh.kubeform.com", Version: "v1alpha1", Resource: "zoneredirections"}

var zoneredirectionsKind = schema.GroupVersionKind{Group: "domain.ovh.kubeform.com", Version: "v1alpha1", Kind: "ZoneRedirection"}

// Get takes name of the zoneRedirection, and returns the corresponding zoneRedirection object, and an error if there is any.
func (c *FakeZoneRedirections) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ZoneRedirection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(zoneredirectionsResource, c.ns, name), &v1alpha1.ZoneRedirection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZoneRedirection), err
}

// List takes label and field selectors, and returns the list of ZoneRedirections that match those selectors.
func (c *FakeZoneRedirections) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ZoneRedirectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(zoneredirectionsResource, zoneredirectionsKind, c.ns, opts), &v1alpha1.ZoneRedirectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ZoneRedirectionList{ListMeta: obj.(*v1alpha1.ZoneRedirectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ZoneRedirectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested zoneRedirections.
func (c *FakeZoneRedirections) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(zoneredirectionsResource, c.ns, opts))

}

// Create takes the representation of a zoneRedirection and creates it.  Returns the server's representation of the zoneRedirection, and an error, if there is any.
func (c *FakeZoneRedirections) Create(ctx context.Context, zoneRedirection *v1alpha1.ZoneRedirection, opts v1.CreateOptions) (result *v1alpha1.ZoneRedirection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(zoneredirectionsResource, c.ns, zoneRedirection), &v1alpha1.ZoneRedirection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZoneRedirection), err
}

// Update takes the representation of a zoneRedirection and updates it. Returns the server's representation of the zoneRedirection, and an error, if there is any.
func (c *FakeZoneRedirections) Update(ctx context.Context, zoneRedirection *v1alpha1.ZoneRedirection, opts v1.UpdateOptions) (result *v1alpha1.ZoneRedirection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(zoneredirectionsResource, c.ns, zoneRedirection), &v1alpha1.ZoneRedirection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZoneRedirection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeZoneRedirections) UpdateStatus(ctx context.Context, zoneRedirection *v1alpha1.ZoneRedirection, opts v1.UpdateOptions) (*v1alpha1.ZoneRedirection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(zoneredirectionsResource, "status", c.ns, zoneRedirection), &v1alpha1.ZoneRedirection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZoneRedirection), err
}

// Delete takes name of the zoneRedirection and deletes it. Returns an error if one occurs.
func (c *FakeZoneRedirections) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(zoneredirectionsResource, c.ns, name), &v1alpha1.ZoneRedirection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeZoneRedirections) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(zoneredirectionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ZoneRedirectionList{})
	return err
}

// Patch applies the patch and returns the patched zoneRedirection.
func (c *FakeZoneRedirections) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ZoneRedirection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(zoneredirectionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ZoneRedirection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ZoneRedirection), err
}
