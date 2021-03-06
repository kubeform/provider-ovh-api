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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubeform.dev/provider-ovh-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// IdentityUsers returns a IdentityUserInformer.
	IdentityUsers() IdentityUserInformer
	// InstallationTemplates returns a InstallationTemplateInformer.
	InstallationTemplates() InstallationTemplateInformer
	// InstallationTemplatePartitionSchemes returns a InstallationTemplatePartitionSchemeInformer.
	InstallationTemplatePartitionSchemes() InstallationTemplatePartitionSchemeInformer
	// InstallationTemplatePartitionSchemeHardwareRaids returns a InstallationTemplatePartitionSchemeHardwareRaidInformer.
	InstallationTemplatePartitionSchemeHardwareRaids() InstallationTemplatePartitionSchemeHardwareRaidInformer
	// InstallationTemplatePartitionSchemePartitions returns a InstallationTemplatePartitionSchemePartitionInformer.
	InstallationTemplatePartitionSchemePartitions() InstallationTemplatePartitionSchemePartitionInformer
	// IpxeScripts returns a IpxeScriptInformer.
	IpxeScripts() IpxeScriptInformer
	// SshKeys returns a SshKeyInformer.
	SshKeys() SshKeyInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// IdentityUsers returns a IdentityUserInformer.
func (v *version) IdentityUsers() IdentityUserInformer {
	return &identityUserInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InstallationTemplates returns a InstallationTemplateInformer.
func (v *version) InstallationTemplates() InstallationTemplateInformer {
	return &installationTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InstallationTemplatePartitionSchemes returns a InstallationTemplatePartitionSchemeInformer.
func (v *version) InstallationTemplatePartitionSchemes() InstallationTemplatePartitionSchemeInformer {
	return &installationTemplatePartitionSchemeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InstallationTemplatePartitionSchemeHardwareRaids returns a InstallationTemplatePartitionSchemeHardwareRaidInformer.
func (v *version) InstallationTemplatePartitionSchemeHardwareRaids() InstallationTemplatePartitionSchemeHardwareRaidInformer {
	return &installationTemplatePartitionSchemeHardwareRaidInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InstallationTemplatePartitionSchemePartitions returns a InstallationTemplatePartitionSchemePartitionInformer.
func (v *version) InstallationTemplatePartitionSchemePartitions() InstallationTemplatePartitionSchemePartitionInformer {
	return &installationTemplatePartitionSchemePartitionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// IpxeScripts returns a IpxeScriptInformer.
func (v *version) IpxeScripts() IpxeScriptInformer {
	return &ipxeScriptInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SshKeys returns a SshKeyInformer.
func (v *version) SshKeys() SshKeyInformer {
	return &sshKeyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
