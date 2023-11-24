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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"
	"net/http"
	"reflect"
	"regexp"
)

// MarketplaceAgreementsServer is a fake server for instances of the armdatadog.MarketplaceAgreementsClient type.
type MarketplaceAgreementsServer struct {
	// CreateOrUpdate is the fake for method MarketplaceAgreementsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, options *armdatadog.MarketplaceAgreementsClientCreateOrUpdateOptions) (resp azfake.Responder[armdatadog.MarketplaceAgreementsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method MarketplaceAgreementsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armdatadog.MarketplaceAgreementsClientListOptions) (resp azfake.PagerResponder[armdatadog.MarketplaceAgreementsClientListResponse])
}

// NewMarketplaceAgreementsServerTransport creates a new instance of MarketplaceAgreementsServerTransport with the provided implementation.
// The returned MarketplaceAgreementsServerTransport instance is connected to an instance of armdatadog.MarketplaceAgreementsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewMarketplaceAgreementsServerTransport(srv *MarketplaceAgreementsServer) *MarketplaceAgreementsServerTransport {
	return &MarketplaceAgreementsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armdatadog.MarketplaceAgreementsClientListResponse]](),
	}
}

// MarketplaceAgreementsServerTransport connects instances of armdatadog.MarketplaceAgreementsClient to instances of MarketplaceAgreementsServer.
// Don't use this type directly, use NewMarketplaceAgreementsServerTransport instead.
type MarketplaceAgreementsServerTransport struct {
	srv          *MarketplaceAgreementsServer
	newListPager *tracker[azfake.PagerResponder[armdatadog.MarketplaceAgreementsClientListResponse]]
}

// Do implements the policy.Transporter interface for MarketplaceAgreementsServerTransport.
func (m *MarketplaceAgreementsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "MarketplaceAgreementsClient.CreateOrUpdate":
		resp, err = m.dispatchCreateOrUpdate(req)
	case "MarketplaceAgreementsClient.NewListPager":
		resp, err = m.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *MarketplaceAgreementsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Datadog/agreements/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatadog.AgreementResource](req)
	if err != nil {
		return nil, err
	}
	var options *armdatadog.MarketplaceAgreementsClientCreateOrUpdateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armdatadog.MarketplaceAgreementsClientCreateOrUpdateOptions{
			Body: &body,
		}
	}
	respr, errRespr := m.srv.CreateOrUpdate(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AgreementResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *MarketplaceAgreementsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Datadog/agreements`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := m.srv.NewListPager(nil)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdatadog.MarketplaceAgreementsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		m.newListPager.remove(req)
	}
	return resp, nil
}