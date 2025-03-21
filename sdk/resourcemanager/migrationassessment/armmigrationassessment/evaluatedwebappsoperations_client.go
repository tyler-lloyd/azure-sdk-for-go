// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationassessment

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// EvaluatedWebAppsOperationsClient contains the methods for the EvaluatedWebAppsOperations group.
// Don't use this type directly, use NewEvaluatedWebAppsOperationsClient() instead.
type EvaluatedWebAppsOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEvaluatedWebAppsOperationsClient creates a new instance of EvaluatedWebAppsOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEvaluatedWebAppsOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EvaluatedWebAppsOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EvaluatedWebAppsOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a EvaluatedWebApp
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - businessCaseName - Business case ARM name
//   - evaluatedWebAppName - Business case Evaluated web App ARM name
//   - options - EvaluatedWebAppsOperationsClientGetOptions contains the optional parameters for the EvaluatedWebAppsOperationsClient.Get
//     method.
func (client *EvaluatedWebAppsOperationsClient) Get(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, evaluatedWebAppName string, options *EvaluatedWebAppsOperationsClientGetOptions) (EvaluatedWebAppsOperationsClientGetResponse, error) {
	var err error
	const operationName = "EvaluatedWebAppsOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, businessCaseName, evaluatedWebAppName, options)
	if err != nil {
		return EvaluatedWebAppsOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EvaluatedWebAppsOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EvaluatedWebAppsOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EvaluatedWebAppsOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, evaluatedWebAppName string, _ *EvaluatedWebAppsOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/businessCases/{businessCaseName}/evaluatedWebApps/{evaluatedWebAppName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if businessCaseName == "" {
		return nil, errors.New("parameter businessCaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{businessCaseName}", url.PathEscape(businessCaseName))
	if evaluatedWebAppName == "" {
		return nil, errors.New("parameter evaluatedWebAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{evaluatedWebAppName}", url.PathEscape(evaluatedWebAppName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EvaluatedWebAppsOperationsClient) getHandleResponse(resp *http.Response) (EvaluatedWebAppsOperationsClientGetResponse, error) {
	result := EvaluatedWebAppsOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EvaluatedWebApp); err != nil {
		return EvaluatedWebAppsOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByBusinessCasePager - List EvaluatedWebApp resources by BusinessCase
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - Assessment Project Name
//   - businessCaseName - Business case ARM name
//   - options - EvaluatedWebAppsOperationsClientListByBusinessCaseOptions contains the optional parameters for the EvaluatedWebAppsOperationsClient.NewListByBusinessCasePager
//     method.
func (client *EvaluatedWebAppsOperationsClient) NewListByBusinessCasePager(resourceGroupName string, projectName string, businessCaseName string, options *EvaluatedWebAppsOperationsClientListByBusinessCaseOptions) *runtime.Pager[EvaluatedWebAppsOperationsClientListByBusinessCaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[EvaluatedWebAppsOperationsClientListByBusinessCaseResponse]{
		More: func(page EvaluatedWebAppsOperationsClientListByBusinessCaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EvaluatedWebAppsOperationsClientListByBusinessCaseResponse) (EvaluatedWebAppsOperationsClientListByBusinessCaseResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EvaluatedWebAppsOperationsClient.NewListByBusinessCasePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByBusinessCaseCreateRequest(ctx, resourceGroupName, projectName, businessCaseName, options)
			}, nil)
			if err != nil {
				return EvaluatedWebAppsOperationsClientListByBusinessCaseResponse{}, err
			}
			return client.listByBusinessCaseHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByBusinessCaseCreateRequest creates the ListByBusinessCase request.
func (client *EvaluatedWebAppsOperationsClient) listByBusinessCaseCreateRequest(ctx context.Context, resourceGroupName string, projectName string, businessCaseName string, options *EvaluatedWebAppsOperationsClientListByBusinessCaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/assessmentProjects/{projectName}/businessCases/{businessCaseName}/evaluatedWebApps"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if businessCaseName == "" {
		return nil, errors.New("parameter businessCaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{businessCaseName}", url.PathEscape(businessCaseName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-01-01-preview")
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	if options != nil && options.TotalRecordCount != nil {
		reqQP.Set("totalRecordCount", strconv.FormatInt(int64(*options.TotalRecordCount), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByBusinessCaseHandleResponse handles the ListByBusinessCase response.
func (client *EvaluatedWebAppsOperationsClient) listByBusinessCaseHandleResponse(resp *http.Response) (EvaluatedWebAppsOperationsClientListByBusinessCaseResponse, error) {
	result := EvaluatedWebAppsOperationsClientListByBusinessCaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EvaluatedWebAppListResult); err != nil {
		return EvaluatedWebAppsOperationsClientListByBusinessCaseResponse{}, err
	}
	return result, nil
}
