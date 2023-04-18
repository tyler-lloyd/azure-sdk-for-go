//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewCacheRulesClient() *CacheRulesClient {
	subClient, _ := NewCacheRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConnectedRegistriesClient() *ConnectedRegistriesClient {
	subClient, _ := NewConnectedRegistriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCredentialSetsClient() *CredentialSetsClient {
	subClient, _ := NewCredentialSetsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewExportPipelinesClient() *ExportPipelinesClient {
	subClient, _ := NewExportPipelinesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegistriesClient() *RegistriesClient {
	subClient, _ := NewRegistriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewImportPipelinesClient() *ImportPipelinesClient {
	subClient, _ := NewImportPipelinesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPipelineRunsClient() *PipelineRunsClient {
	subClient, _ := NewPipelineRunsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	subClient, _ := NewPrivateEndpointConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReplicationsClient() *ReplicationsClient {
	subClient, _ := NewReplicationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewScopeMapsClient() *ScopeMapsClient {
	subClient, _ := NewScopeMapsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTokensClient() *TokensClient {
	subClient, _ := NewTokensClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWebhooksClient() *WebhooksClient {
	subClient, _ := NewWebhooksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAgentPoolsClient() *AgentPoolsClient {
	subClient, _ := NewAgentPoolsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRunsClient() *RunsClient {
	subClient, _ := NewRunsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTaskRunsClient() *TaskRunsClient {
	subClient, _ := NewTaskRunsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTasksClient() *TasksClient {
	subClient, _ := NewTasksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}