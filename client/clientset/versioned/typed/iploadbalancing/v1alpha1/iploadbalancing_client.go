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
	v1alpha1 "kubeform.dev/provider-ovh-api/apis/iploadbalancing/v1alpha1"
	"kubeform.dev/provider-ovh-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type IploadbalancingV1alpha1Interface interface {
	RESTClient() rest.Interface
	HttpFarmsGetter
	HttpFarmServersGetter
	HttpFrontendsGetter
	HttpRoutesGetter
	HttpRouteRulesGetter
	IploadbalancingsGetter
	RefreshesGetter
	TcpFarmsGetter
	TcpFarmServersGetter
	TcpFrontendsGetter
	TcpRoutesGetter
	TcpRouteRulesGetter
	VrackNetworksGetter
}

// IploadbalancingV1alpha1Client is used to interact with features provided by the iploadbalancing.ovh.kubeform.com group.
type IploadbalancingV1alpha1Client struct {
	restClient rest.Interface
}

func (c *IploadbalancingV1alpha1Client) HttpFarms(namespace string) HttpFarmInterface {
	return newHttpFarms(c, namespace)
}

func (c *IploadbalancingV1alpha1Client) HttpFarmServers(namespace string) HttpFarmServerInterface {
	return newHttpFarmServers(c, namespace)
}

func (c *IploadbalancingV1alpha1Client) HttpFrontends(namespace string) HttpFrontendInterface {
	return newHttpFrontends(c, namespace)
}

func (c *IploadbalancingV1alpha1Client) HttpRoutes(namespace string) HttpRouteInterface {
	return newHttpRoutes(c, namespace)
}

func (c *IploadbalancingV1alpha1Client) HttpRouteRules(namespace string) HttpRouteRuleInterface {
	return newHttpRouteRules(c, namespace)
}

func (c *IploadbalancingV1alpha1Client) Iploadbalancings(namespace string) IploadbalancingInterface {
	return newIploadbalancings(c, namespace)
}

func (c *IploadbalancingV1alpha1Client) Refreshes(namespace string) RefreshInterface {
	return newRefreshes(c, namespace)
}

func (c *IploadbalancingV1alpha1Client) TcpFarms(namespace string) TcpFarmInterface {
	return newTcpFarms(c, namespace)
}

func (c *IploadbalancingV1alpha1Client) TcpFarmServers(namespace string) TcpFarmServerInterface {
	return newTcpFarmServers(c, namespace)
}

func (c *IploadbalancingV1alpha1Client) TcpFrontends(namespace string) TcpFrontendInterface {
	return newTcpFrontends(c, namespace)
}

func (c *IploadbalancingV1alpha1Client) TcpRoutes(namespace string) TcpRouteInterface {
	return newTcpRoutes(c, namespace)
}

func (c *IploadbalancingV1alpha1Client) TcpRouteRules(namespace string) TcpRouteRuleInterface {
	return newTcpRouteRules(c, namespace)
}

func (c *IploadbalancingV1alpha1Client) VrackNetworks(namespace string) VrackNetworkInterface {
	return newVrackNetworks(c, namespace)
}

// NewForConfig creates a new IploadbalancingV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*IploadbalancingV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &IploadbalancingV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new IploadbalancingV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *IploadbalancingV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new IploadbalancingV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *IploadbalancingV1alpha1Client {
	return &IploadbalancingV1alpha1Client{c}
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
func (c *IploadbalancingV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
