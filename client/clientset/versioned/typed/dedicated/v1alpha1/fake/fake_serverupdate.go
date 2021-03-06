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

	v1alpha1 "kubeform.dev/provider-ovh-api/apis/dedicated/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServerUpdates implements ServerUpdateInterface
type FakeServerUpdates struct {
	Fake *FakeDedicatedV1alpha1
	ns   string
}

var serverupdatesResource = schema.GroupVersionResource{Group: "dedicated.ovh.kubeform.com", Version: "v1alpha1", Resource: "serverupdates"}

var serverupdatesKind = schema.GroupVersionKind{Group: "dedicated.ovh.kubeform.com", Version: "v1alpha1", Kind: "ServerUpdate"}

// Get takes name of the serverUpdate, and returns the corresponding serverUpdate object, and an error if there is any.
func (c *FakeServerUpdates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServerUpdate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serverupdatesResource, c.ns, name), &v1alpha1.ServerUpdate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerUpdate), err
}

// List takes label and field selectors, and returns the list of ServerUpdates that match those selectors.
func (c *FakeServerUpdates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServerUpdateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serverupdatesResource, serverupdatesKind, c.ns, opts), &v1alpha1.ServerUpdateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServerUpdateList{ListMeta: obj.(*v1alpha1.ServerUpdateList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServerUpdateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serverUpdates.
func (c *FakeServerUpdates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serverupdatesResource, c.ns, opts))

}

// Create takes the representation of a serverUpdate and creates it.  Returns the server's representation of the serverUpdate, and an error, if there is any.
func (c *FakeServerUpdates) Create(ctx context.Context, serverUpdate *v1alpha1.ServerUpdate, opts v1.CreateOptions) (result *v1alpha1.ServerUpdate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serverupdatesResource, c.ns, serverUpdate), &v1alpha1.ServerUpdate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerUpdate), err
}

// Update takes the representation of a serverUpdate and updates it. Returns the server's representation of the serverUpdate, and an error, if there is any.
func (c *FakeServerUpdates) Update(ctx context.Context, serverUpdate *v1alpha1.ServerUpdate, opts v1.UpdateOptions) (result *v1alpha1.ServerUpdate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serverupdatesResource, c.ns, serverUpdate), &v1alpha1.ServerUpdate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerUpdate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServerUpdates) UpdateStatus(ctx context.Context, serverUpdate *v1alpha1.ServerUpdate, opts v1.UpdateOptions) (*v1alpha1.ServerUpdate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serverupdatesResource, "status", c.ns, serverUpdate), &v1alpha1.ServerUpdate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerUpdate), err
}

// Delete takes name of the serverUpdate and deletes it. Returns an error if one occurs.
func (c *FakeServerUpdates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serverupdatesResource, c.ns, name), &v1alpha1.ServerUpdate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServerUpdates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serverupdatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServerUpdateList{})
	return err
}

// Patch applies the patch and returns the patched serverUpdate.
func (c *FakeServerUpdates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServerUpdate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serverupdatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServerUpdate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServerUpdate), err
}
