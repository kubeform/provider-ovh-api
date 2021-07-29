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

// FakeVracks implements VrackInterface
type FakeVracks struct {
	Fake *FakeVrackV1alpha1
	ns   string
}

var vracksResource = schema.GroupVersionResource{Group: "vrack.ovh.kubeform.com", Version: "v1alpha1", Resource: "vracks"}

var vracksKind = schema.GroupVersionKind{Group: "vrack.ovh.kubeform.com", Version: "v1alpha1", Kind: "Vrack"}

// Get takes name of the vrack, and returns the corresponding vrack object, and an error if there is any.
func (c *FakeVracks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Vrack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vracksResource, c.ns, name), &v1alpha1.Vrack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vrack), err
}

// List takes label and field selectors, and returns the list of Vracks that match those selectors.
func (c *FakeVracks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VrackList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vracksResource, vracksKind, c.ns, opts), &v1alpha1.VrackList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VrackList{ListMeta: obj.(*v1alpha1.VrackList).ListMeta}
	for _, item := range obj.(*v1alpha1.VrackList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vracks.
func (c *FakeVracks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vracksResource, c.ns, opts))

}

// Create takes the representation of a vrack and creates it.  Returns the server's representation of the vrack, and an error, if there is any.
func (c *FakeVracks) Create(ctx context.Context, vrack *v1alpha1.Vrack, opts v1.CreateOptions) (result *v1alpha1.Vrack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vracksResource, c.ns, vrack), &v1alpha1.Vrack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vrack), err
}

// Update takes the representation of a vrack and updates it. Returns the server's representation of the vrack, and an error, if there is any.
func (c *FakeVracks) Update(ctx context.Context, vrack *v1alpha1.Vrack, opts v1.UpdateOptions) (result *v1alpha1.Vrack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vracksResource, c.ns, vrack), &v1alpha1.Vrack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vrack), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVracks) UpdateStatus(ctx context.Context, vrack *v1alpha1.Vrack, opts v1.UpdateOptions) (*v1alpha1.Vrack, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vracksResource, "status", c.ns, vrack), &v1alpha1.Vrack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vrack), err
}

// Delete takes name of the vrack and deletes it. Returns an error if one occurs.
func (c *FakeVracks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vracksResource, c.ns, name), &v1alpha1.Vrack{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVracks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vracksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VrackList{})
	return err
}

// Patch applies the patch and returns the patched vrack.
func (c *FakeVracks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Vrack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vracksResource, c.ns, name, pt, data, subresources...), &v1alpha1.Vrack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Vrack), err
}
