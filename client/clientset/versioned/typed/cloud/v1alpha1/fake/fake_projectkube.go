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

	v1alpha1 "kubeform.dev/provider-ovh-api/apis/cloud/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProjectKubes implements ProjectKubeInterface
type FakeProjectKubes struct {
	Fake *FakeCloudV1alpha1
	ns   string
}

var projectkubesResource = schema.GroupVersionResource{Group: "cloud.ovh.kubeform.com", Version: "v1alpha1", Resource: "projectkubes"}

var projectkubesKind = schema.GroupVersionKind{Group: "cloud.ovh.kubeform.com", Version: "v1alpha1", Kind: "ProjectKube"}

// Get takes name of the projectKube, and returns the corresponding projectKube object, and an error if there is any.
func (c *FakeProjectKubes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProjectKube, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(projectkubesResource, c.ns, name), &v1alpha1.ProjectKube{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectKube), err
}

// List takes label and field selectors, and returns the list of ProjectKubes that match those selectors.
func (c *FakeProjectKubes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProjectKubeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(projectkubesResource, projectkubesKind, c.ns, opts), &v1alpha1.ProjectKubeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProjectKubeList{ListMeta: obj.(*v1alpha1.ProjectKubeList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProjectKubeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projectKubes.
func (c *FakeProjectKubes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(projectkubesResource, c.ns, opts))

}

// Create takes the representation of a projectKube and creates it.  Returns the server's representation of the projectKube, and an error, if there is any.
func (c *FakeProjectKubes) Create(ctx context.Context, projectKube *v1alpha1.ProjectKube, opts v1.CreateOptions) (result *v1alpha1.ProjectKube, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(projectkubesResource, c.ns, projectKube), &v1alpha1.ProjectKube{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectKube), err
}

// Update takes the representation of a projectKube and updates it. Returns the server's representation of the projectKube, and an error, if there is any.
func (c *FakeProjectKubes) Update(ctx context.Context, projectKube *v1alpha1.ProjectKube, opts v1.UpdateOptions) (result *v1alpha1.ProjectKube, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(projectkubesResource, c.ns, projectKube), &v1alpha1.ProjectKube{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectKube), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProjectKubes) UpdateStatus(ctx context.Context, projectKube *v1alpha1.ProjectKube, opts v1.UpdateOptions) (*v1alpha1.ProjectKube, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(projectkubesResource, "status", c.ns, projectKube), &v1alpha1.ProjectKube{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectKube), err
}

// Delete takes name of the projectKube and deletes it. Returns an error if one occurs.
func (c *FakeProjectKubes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(projectkubesResource, c.ns, name), &v1alpha1.ProjectKube{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProjectKubes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(projectkubesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProjectKubeList{})
	return err
}

// Patch applies the patch and returns the patched projectKube.
func (c *FakeProjectKubes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProjectKube, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(projectkubesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ProjectKube{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProjectKube), err
}
