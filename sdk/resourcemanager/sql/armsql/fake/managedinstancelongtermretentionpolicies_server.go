//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"net/http"
	"net/url"
	"regexp"
)

// ManagedInstanceLongTermRetentionPoliciesServer is a fake server for instances of the armsql.ManagedInstanceLongTermRetentionPoliciesClient type.
type ManagedInstanceLongTermRetentionPoliciesServer struct {
	// BeginCreateOrUpdate is the fake for method ManagedInstanceLongTermRetentionPoliciesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName armsql.ManagedInstanceLongTermRetentionPolicyName, parameters armsql.ManagedInstanceLongTermRetentionPolicy, options *armsql.ManagedInstanceLongTermRetentionPoliciesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsql.ManagedInstanceLongTermRetentionPoliciesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagedInstanceLongTermRetentionPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, policyName armsql.ManagedInstanceLongTermRetentionPolicyName, options *armsql.ManagedInstanceLongTermRetentionPoliciesClientGetOptions) (resp azfake.Responder[armsql.ManagedInstanceLongTermRetentionPoliciesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDatabasePager is the fake for method ManagedInstanceLongTermRetentionPoliciesClient.NewListByDatabasePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDatabasePager func(resourceGroupName string, managedInstanceName string, databaseName string, options *armsql.ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseOptions) (resp azfake.PagerResponder[armsql.ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseResponse])
}

// NewManagedInstanceLongTermRetentionPoliciesServerTransport creates a new instance of ManagedInstanceLongTermRetentionPoliciesServerTransport with the provided implementation.
// The returned ManagedInstanceLongTermRetentionPoliciesServerTransport instance is connected to an instance of armsql.ManagedInstanceLongTermRetentionPoliciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedInstanceLongTermRetentionPoliciesServerTransport(srv *ManagedInstanceLongTermRetentionPoliciesServer) *ManagedInstanceLongTermRetentionPoliciesServerTransport {
	return &ManagedInstanceLongTermRetentionPoliciesServerTransport{
		srv:                    srv,
		beginCreateOrUpdate:    newTracker[azfake.PollerResponder[armsql.ManagedInstanceLongTermRetentionPoliciesClientCreateOrUpdateResponse]](),
		newListByDatabasePager: newTracker[azfake.PagerResponder[armsql.ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseResponse]](),
	}
}

// ManagedInstanceLongTermRetentionPoliciesServerTransport connects instances of armsql.ManagedInstanceLongTermRetentionPoliciesClient to instances of ManagedInstanceLongTermRetentionPoliciesServer.
// Don't use this type directly, use NewManagedInstanceLongTermRetentionPoliciesServerTransport instead.
type ManagedInstanceLongTermRetentionPoliciesServerTransport struct {
	srv                    *ManagedInstanceLongTermRetentionPoliciesServer
	beginCreateOrUpdate    *tracker[azfake.PollerResponder[armsql.ManagedInstanceLongTermRetentionPoliciesClientCreateOrUpdateResponse]]
	newListByDatabasePager *tracker[azfake.PagerResponder[armsql.ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseResponse]]
}

// Do implements the policy.Transporter interface for ManagedInstanceLongTermRetentionPoliciesServerTransport.
func (m *ManagedInstanceLongTermRetentionPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedInstanceLongTermRetentionPoliciesClient.BeginCreateOrUpdate":
		resp, err = m.dispatchBeginCreateOrUpdate(req)
	case "ManagedInstanceLongTermRetentionPoliciesClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagedInstanceLongTermRetentionPoliciesClient.NewListByDatabasePager":
		resp, err = m.dispatchNewListByDatabasePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedInstanceLongTermRetentionPoliciesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := m.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupLongTermRetentionPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.ManagedInstanceLongTermRetentionPolicy](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		policyNameParam, err := parseWithCast(matches[regex.SubexpIndex("policyName")], func(v string) (armsql.ManagedInstanceLongTermRetentionPolicyName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armsql.ManagedInstanceLongTermRetentionPolicyName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, policyNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		m.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		m.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (m *ManagedInstanceLongTermRetentionPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupLongTermRetentionPolicies/(?P<policyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	policyNameParam, err := parseWithCast(matches[regex.SubexpIndex("policyName")], func(v string) (armsql.ManagedInstanceLongTermRetentionPolicyName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsql.ManagedInstanceLongTermRetentionPolicyName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, policyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedInstanceLongTermRetentionPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedInstanceLongTermRetentionPoliciesServerTransport) dispatchNewListByDatabasePager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByDatabasePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDatabasePager not implemented")}
	}
	newListByDatabasePager := m.newListByDatabasePager.get(req)
	if newListByDatabasePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupLongTermRetentionPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListByDatabasePager(resourceGroupNameParam, managedInstanceNameParam, databaseNameParam, nil)
		newListByDatabasePager = &resp
		m.newListByDatabasePager.add(req, newListByDatabasePager)
		server.PagerResponderInjectNextLinks(newListByDatabasePager, req, func(page *armsql.ManagedInstanceLongTermRetentionPoliciesClientListByDatabaseResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDatabasePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByDatabasePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDatabasePager) {
		m.newListByDatabasePager.remove(req)
	}
	return resp, nil
}