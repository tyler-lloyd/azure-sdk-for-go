//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
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

// ServerOperationsServer is a fake server for instances of the armsql.ServerOperationsClient type.
type ServerOperationsServer struct {
	// NewListByServerPager is the fake for method ServerOperationsClient.NewListByServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerPager func(resourceGroupName string, serverName string, options *armsql.ServerOperationsClientListByServerOptions) (resp azfake.PagerResponder[armsql.ServerOperationsClientListByServerResponse])
}

// NewServerOperationsServerTransport creates a new instance of ServerOperationsServerTransport with the provided implementation.
// The returned ServerOperationsServerTransport instance is connected to an instance of armsql.ServerOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerOperationsServerTransport(srv *ServerOperationsServer) *ServerOperationsServerTransport {
	return &ServerOperationsServerTransport{
		srv:                  srv,
		newListByServerPager: newTracker[azfake.PagerResponder[armsql.ServerOperationsClientListByServerResponse]](),
	}
}

// ServerOperationsServerTransport connects instances of armsql.ServerOperationsClient to instances of ServerOperationsServer.
// Don't use this type directly, use NewServerOperationsServerTransport instead.
type ServerOperationsServerTransport struct {
	srv                  *ServerOperationsServer
	newListByServerPager *tracker[azfake.PagerResponder[armsql.ServerOperationsClientListByServerResponse]]
}

// Do implements the policy.Transporter interface for ServerOperationsServerTransport.
func (s *ServerOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ServerOperationsClient.NewListByServerPager":
		resp, err = s.dispatchNewListByServerPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServerOperationsServerTransport) dispatchNewListByServerPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerPager not implemented")}
	}
	newListByServerPager := s.newListByServerPager.get(req)
	if newListByServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByServerPager(resourceGroupNameParam, serverNameParam, nil)
		newListByServerPager = &resp
		s.newListByServerPager.add(req, newListByServerPager)
		server.PagerResponderInjectNextLinks(newListByServerPager, req, func(page *armsql.ServerOperationsClientListByServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerPager) {
		s.newListByServerPager.remove(req)
	}
	return resp, nil
}