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

	v1alpha1 "kubeform.dev/provider-ovh-api/apis/me/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInstallationTemplatePartitionSchemes implements InstallationTemplatePartitionSchemeInterface
type FakeInstallationTemplatePartitionSchemes struct {
	Fake *FakeMeV1alpha1
	ns   string
}

var installationtemplatepartitionschemesResource = schema.GroupVersionResource{Group: "me.ovh.kubeform.com", Version: "v1alpha1", Resource: "installationtemplatepartitionschemes"}

var installationtemplatepartitionschemesKind = schema.GroupVersionKind{Group: "me.ovh.kubeform.com", Version: "v1alpha1", Kind: "InstallationTemplatePartitionScheme"}

// Get takes name of the installationTemplatePartitionScheme, and returns the corresponding installationTemplatePartitionScheme object, and an error if there is any.
func (c *FakeInstallationTemplatePartitionSchemes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InstallationTemplatePartitionScheme, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(installationtemplatepartitionschemesResource, c.ns, name), &v1alpha1.InstallationTemplatePartitionScheme{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstallationTemplatePartitionScheme), err
}

// List takes label and field selectors, and returns the list of InstallationTemplatePartitionSchemes that match those selectors.
func (c *FakeInstallationTemplatePartitionSchemes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InstallationTemplatePartitionSchemeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(installationtemplatepartitionschemesResource, installationtemplatepartitionschemesKind, c.ns, opts), &v1alpha1.InstallationTemplatePartitionSchemeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InstallationTemplatePartitionSchemeList{ListMeta: obj.(*v1alpha1.InstallationTemplatePartitionSchemeList).ListMeta}
	for _, item := range obj.(*v1alpha1.InstallationTemplatePartitionSchemeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested installationTemplatePartitionSchemes.
func (c *FakeInstallationTemplatePartitionSchemes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(installationtemplatepartitionschemesResource, c.ns, opts))

}

// Create takes the representation of a installationTemplatePartitionScheme and creates it.  Returns the server's representation of the installationTemplatePartitionScheme, and an error, if there is any.
func (c *FakeInstallationTemplatePartitionSchemes) Create(ctx context.Context, installationTemplatePartitionScheme *v1alpha1.InstallationTemplatePartitionScheme, opts v1.CreateOptions) (result *v1alpha1.InstallationTemplatePartitionScheme, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(installationtemplatepartitionschemesResource, c.ns, installationTemplatePartitionScheme), &v1alpha1.InstallationTemplatePartitionScheme{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstallationTemplatePartitionScheme), err
}

// Update takes the representation of a installationTemplatePartitionScheme and updates it. Returns the server's representation of the installationTemplatePartitionScheme, and an error, if there is any.
func (c *FakeInstallationTemplatePartitionSchemes) Update(ctx context.Context, installationTemplatePartitionScheme *v1alpha1.InstallationTemplatePartitionScheme, opts v1.UpdateOptions) (result *v1alpha1.InstallationTemplatePartitionScheme, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(installationtemplatepartitionschemesResource, c.ns, installationTemplatePartitionScheme), &v1alpha1.InstallationTemplatePartitionScheme{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstallationTemplatePartitionScheme), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInstallationTemplatePartitionSchemes) UpdateStatus(ctx context.Context, installationTemplatePartitionScheme *v1alpha1.InstallationTemplatePartitionScheme, opts v1.UpdateOptions) (*v1alpha1.InstallationTemplatePartitionScheme, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(installationtemplatepartitionschemesResource, "status", c.ns, installationTemplatePartitionScheme), &v1alpha1.InstallationTemplatePartitionScheme{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstallationTemplatePartitionScheme), err
}

// Delete takes name of the installationTemplatePartitionScheme and deletes it. Returns an error if one occurs.
func (c *FakeInstallationTemplatePartitionSchemes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(installationtemplatepartitionschemesResource, c.ns, name), &v1alpha1.InstallationTemplatePartitionScheme{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInstallationTemplatePartitionSchemes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(installationtemplatepartitionschemesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InstallationTemplatePartitionSchemeList{})
	return err
}

// Patch applies the patch and returns the patched installationTemplatePartitionScheme.
func (c *FakeInstallationTemplatePartitionSchemes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InstallationTemplatePartitionScheme, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(installationtemplatepartitionschemesResource, c.ns, name, pt, data, subresources...), &v1alpha1.InstallationTemplatePartitionScheme{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstallationTemplatePartitionScheme), err
}
