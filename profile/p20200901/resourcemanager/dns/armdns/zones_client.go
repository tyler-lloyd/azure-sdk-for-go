//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdns

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/p20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ZonesClient contains the methods for the Zones group.
// Don't use this type directly, use NewZonesClient() instead.
type ZonesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewZonesClient creates a new instance of ZonesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewZonesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ZonesClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armdns.ZonesClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ZonesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a DNS zone. Does not modify DNS records within the zone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - zoneName - The name of the DNS zone (without a terminating dot).
//   - parameters - Parameters supplied to the CreateOrUpdate operation.
//   - options - ZonesClientCreateOrUpdateOptions contains the optional parameters for the ZonesClient.CreateOrUpdate method.
func (client *ZonesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, zoneName string, parameters Zone, options *ZonesClientCreateOrUpdateOptions) (ZonesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, zoneName, parameters, options)
	if err != nil {
		return ZonesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ZonesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ZonesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ZonesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, parameters Zone, options *ZonesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ZonesClient) createOrUpdateHandleResponse(resp *http.Response) (ZonesClientCreateOrUpdateResponse, error) {
	result := ZonesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Zone); err != nil {
		return ZonesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes a DNS zone. WARNING: All DNS records in the zone will also be deleted. This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - zoneName - The name of the DNS zone (without a terminating dot).
//   - options - ZonesClientBeginDeleteOptions contains the optional parameters for the ZonesClient.BeginDelete method.
func (client *ZonesClient) BeginDelete(ctx context.Context, resourceGroupName string, zoneName string, options *ZonesClientBeginDeleteOptions) (*runtime.Poller[ZonesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, zoneName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ZonesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ZonesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a DNS zone. WARNING: All DNS records in the zone will also be deleted. This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-04-01
func (client *ZonesClient) deleteOperation(ctx context.Context, resourceGroupName string, zoneName string, options *ZonesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, zoneName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ZonesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, options *ZonesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a DNS zone. Retrieves the zone properties, but not the record sets within the zone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - zoneName - The name of the DNS zone (without a terminating dot).
//   - options - ZonesClientGetOptions contains the optional parameters for the ZonesClient.Get method.
func (client *ZonesClient) Get(ctx context.Context, resourceGroupName string, zoneName string, options *ZonesClientGetOptions) (ZonesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, zoneName, options)
	if err != nil {
		return ZonesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ZonesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ZonesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ZonesClient) getCreateRequest(ctx context.Context, resourceGroupName string, zoneName string, options *ZonesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if zoneName == "" {
		return nil, errors.New("parameter zoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{zoneName}", url.PathEscape(zoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ZonesClient) getHandleResponse(resp *http.Response) (ZonesClientGetResponse, error) {
	result := ZonesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Zone); err != nil {
		return ZonesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the DNS zones in all resource groups in a subscription.
//
// Generated from API version 2016-04-01
//   - options - ZonesClientListOptions contains the optional parameters for the ZonesClient.NewListPager method.
func (client *ZonesClient) NewListPager(options *ZonesClientListOptions) *runtime.Pager[ZonesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ZonesClientListResponse]{
		More: func(page ZonesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ZonesClientListResponse) (ZonesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ZonesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ZonesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ZonesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ZonesClient) listCreateRequest(ctx context.Context, options *ZonesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnszones"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2016-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ZonesClient) listHandleResponse(resp *http.Response) (ZonesClientListResponse, error) {
	result := ZonesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ZoneListResult); err != nil {
		return ZonesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists the DNS zones within a resource group.
//
// Generated from API version 2016-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ZonesClientListByResourceGroupOptions contains the optional parameters for the ZonesClient.NewListByResourceGroupPager
//     method.
func (client *ZonesClient) NewListByResourceGroupPager(resourceGroupName string, options *ZonesClientListByResourceGroupOptions) *runtime.Pager[ZonesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ZonesClientListByResourceGroupResponse]{
		More: func(page ZonesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ZonesClientListByResourceGroupResponse) (ZonesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ZonesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ZonesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ZonesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ZonesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ZonesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2016-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ZonesClient) listByResourceGroupHandleResponse(resp *http.Response) (ZonesClientListByResourceGroupResponse, error) {
	result := ZonesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ZoneListResult); err != nil {
		return ZonesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}