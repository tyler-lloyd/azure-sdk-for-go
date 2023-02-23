//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WorkspaceManagedSQLServerSecurityAlertPolicyClient contains the methods for the WorkspaceManagedSQLServerSecurityAlertPolicy group.
// Don't use this type directly, use NewWorkspaceManagedSQLServerSecurityAlertPolicyClient() instead.
type WorkspaceManagedSQLServerSecurityAlertPolicyClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspaceManagedSQLServerSecurityAlertPolicyClient creates a new instance of WorkspaceManagedSQLServerSecurityAlertPolicyClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceManagedSQLServerSecurityAlertPolicyClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceManagedSQLServerSecurityAlertPolicyClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceManagedSQLServerSecurityAlertPolicyClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or Update a workspace managed sql server's threat detection policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - securityAlertPolicyName - The name of the security alert policy.
//   - parameters - The workspace managed sql server security alert policy.
//   - options - WorkspaceManagedSQLServerSecurityAlertPolicyClientBeginCreateOrUpdateOptions contains the optional parameters
//     for the WorkspaceManagedSQLServerSecurityAlertPolicyClient.BeginCreateOrUpdate method.
func (client *WorkspaceManagedSQLServerSecurityAlertPolicyClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, securityAlertPolicyName SecurityAlertPolicyNameAutoGenerated, parameters ServerSecurityAlertPolicy, options *WorkspaceManagedSQLServerSecurityAlertPolicyClientBeginCreateOrUpdateOptions) (*runtime.Poller[WorkspaceManagedSQLServerSecurityAlertPolicyClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, securityAlertPolicyName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[WorkspaceManagedSQLServerSecurityAlertPolicyClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[WorkspaceManagedSQLServerSecurityAlertPolicyClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or Update a workspace managed sql server's threat detection policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
func (client *WorkspaceManagedSQLServerSecurityAlertPolicyClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, securityAlertPolicyName SecurityAlertPolicyNameAutoGenerated, parameters ServerSecurityAlertPolicy, options *WorkspaceManagedSQLServerSecurityAlertPolicyClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, securityAlertPolicyName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceManagedSQLServerSecurityAlertPolicyClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, securityAlertPolicyName SecurityAlertPolicyNameAutoGenerated, parameters ServerSecurityAlertPolicy, options *WorkspaceManagedSQLServerSecurityAlertPolicyClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/securityAlertPolicies/{securityAlertPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if securityAlertPolicyName == "" {
		return nil, errors.New("parameter securityAlertPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityAlertPolicyName}", url.PathEscape(string(securityAlertPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Get a workspace managed sql server's security alert policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - securityAlertPolicyName - The name of the security alert policy.
//   - options - WorkspaceManagedSQLServerSecurityAlertPolicyClientGetOptions contains the optional parameters for the WorkspaceManagedSQLServerSecurityAlertPolicyClient.Get
//     method.
func (client *WorkspaceManagedSQLServerSecurityAlertPolicyClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, securityAlertPolicyName SecurityAlertPolicyNameAutoGenerated, options *WorkspaceManagedSQLServerSecurityAlertPolicyClientGetOptions) (WorkspaceManagedSQLServerSecurityAlertPolicyClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, securityAlertPolicyName, options)
	if err != nil {
		return WorkspaceManagedSQLServerSecurityAlertPolicyClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceManagedSQLServerSecurityAlertPolicyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceManagedSQLServerSecurityAlertPolicyClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagedSQLServerSecurityAlertPolicyClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, securityAlertPolicyName SecurityAlertPolicyNameAutoGenerated, options *WorkspaceManagedSQLServerSecurityAlertPolicyClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/securityAlertPolicies/{securityAlertPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if securityAlertPolicyName == "" {
		return nil, errors.New("parameter securityAlertPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityAlertPolicyName}", url.PathEscape(string(securityAlertPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceManagedSQLServerSecurityAlertPolicyClient) getHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerSecurityAlertPolicyClientGetResponse, error) {
	result := WorkspaceManagedSQLServerSecurityAlertPolicyClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerSecurityAlertPolicy); err != nil {
		return WorkspaceManagedSQLServerSecurityAlertPolicyClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get workspace managed sql server's threat detection policies.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - WorkspaceManagedSQLServerSecurityAlertPolicyClientListOptions contains the optional parameters for the WorkspaceManagedSQLServerSecurityAlertPolicyClient.NewListPager
//     method.
func (client *WorkspaceManagedSQLServerSecurityAlertPolicyClient) NewListPager(resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerSecurityAlertPolicyClientListOptions) *runtime.Pager[WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse]{
		More: func(page WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse) (WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WorkspaceManagedSQLServerSecurityAlertPolicyClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerSecurityAlertPolicyClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/securityAlertPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkspaceManagedSQLServerSecurityAlertPolicyClient) listHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse, error) {
	result := WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerSecurityAlertPolicyListResult); err != nil {
		return WorkspaceManagedSQLServerSecurityAlertPolicyClientListResponse{}, err
	}
	return result, nil
}