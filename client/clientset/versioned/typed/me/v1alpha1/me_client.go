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

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-ovh-api/apis/me/v1alpha1"
	"kubeform.dev/provider-ovh-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type MeV1alpha1Interface interface {
	RESTClient() rest.Interface
	IdentityUsersGetter
	InstallationTemplatesGetter
	InstallationTemplatePartitionSchemesGetter
	InstallationTemplatePartitionSchemeHardwareRaidsGetter
	InstallationTemplatePartitionSchemePartitionsGetter
	IpxeScriptsGetter
	SshKeysGetter
}

// MeV1alpha1Client is used to interact with features provided by the me.ovh.kubeform.com group.
type MeV1alpha1Client struct {
	restClient rest.Interface
}

func (c *MeV1alpha1Client) IdentityUsers(namespace string) IdentityUserInterface {
	return newIdentityUsers(c, namespace)
}

func (c *MeV1alpha1Client) InstallationTemplates(namespace string) InstallationTemplateInterface {
	return newInstallationTemplates(c, namespace)
}

func (c *MeV1alpha1Client) InstallationTemplatePartitionSchemes(namespace string) InstallationTemplatePartitionSchemeInterface {
	return newInstallationTemplatePartitionSchemes(c, namespace)
}

func (c *MeV1alpha1Client) InstallationTemplatePartitionSchemeHardwareRaids(namespace string) InstallationTemplatePartitionSchemeHardwareRaidInterface {
	return newInstallationTemplatePartitionSchemeHardwareRaids(c, namespace)
}

func (c *MeV1alpha1Client) InstallationTemplatePartitionSchemePartitions(namespace string) InstallationTemplatePartitionSchemePartitionInterface {
	return newInstallationTemplatePartitionSchemePartitions(c, namespace)
}

func (c *MeV1alpha1Client) IpxeScripts(namespace string) IpxeScriptInterface {
	return newIpxeScripts(c, namespace)
}

func (c *MeV1alpha1Client) SshKeys(namespace string) SshKeyInterface {
	return newSshKeys(c, namespace)
}

// NewForConfig creates a new MeV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*MeV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MeV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new MeV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MeV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MeV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *MeV1alpha1Client {
	return &MeV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MeV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
