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

// FakeCloudprojects implements CloudprojectInterface
type FakeCloudprojects struct {
	Fake *FakeVrackV1alpha1
	ns   string
}

var cloudprojectsResource = schema.GroupVersionResource{Group: "vrack.ovh.kubeform.com", Version: "v1alpha1", Resource: "cloudprojects"}

var cloudprojectsKind = schema.GroupVersionKind{Group: "vrack.ovh.kubeform.com", Version: "v1alpha1", Kind: "Cloudproject"}

// Get takes name of the cloudproject, and returns the corresponding cloudproject object, and an error if there is any.
func (c *FakeCloudprojects) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Cloudproject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudprojectsResource, c.ns, name), &v1alpha1.Cloudproject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cloudproject), err
}

// List takes label and field selectors, and returns the list of Cloudprojects that match those selectors.
func (c *FakeCloudprojects) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloudprojectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudprojectsResource, cloudprojectsKind, c.ns, opts), &v1alpha1.CloudprojectList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudprojectList{ListMeta: obj.(*v1alpha1.CloudprojectList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudprojectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudprojects.
func (c *FakeCloudprojects) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudprojectsResource, c.ns, opts))

}

// Create takes the representation of a cloudproject and creates it.  Returns the server's representation of the cloudproject, and an error, if there is any.
func (c *FakeCloudprojects) Create(ctx context.Context, cloudproject *v1alpha1.Cloudproject, opts v1.CreateOptions) (result *v1alpha1.Cloudproject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudprojectsResource, c.ns, cloudproject), &v1alpha1.Cloudproject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cloudproject), err
}

// Update takes the representation of a cloudproject and updates it. Returns the server's representation of the cloudproject, and an error, if there is any.
func (c *FakeCloudprojects) Update(ctx context.Context, cloudproject *v1alpha1.Cloudproject, opts v1.UpdateOptions) (result *v1alpha1.Cloudproject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudprojectsResource, c.ns, cloudproject), &v1alpha1.Cloudproject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cloudproject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudprojects) UpdateStatus(ctx context.Context, cloudproject *v1alpha1.Cloudproject, opts v1.UpdateOptions) (*v1alpha1.Cloudproject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudprojectsResource, "status", c.ns, cloudproject), &v1alpha1.Cloudproject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cloudproject), err
}

// Delete takes name of the cloudproject and deletes it. Returns an error if one occurs.
func (c *FakeCloudprojects) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudprojectsResource, c.ns, name), &v1alpha1.Cloudproject{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudprojects) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudprojectsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudprojectList{})
	return err
}

// Patch applies the patch and returns the patched cloudproject.
func (c *FakeCloudprojects) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Cloudproject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudprojectsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Cloudproject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cloudproject), err
}
