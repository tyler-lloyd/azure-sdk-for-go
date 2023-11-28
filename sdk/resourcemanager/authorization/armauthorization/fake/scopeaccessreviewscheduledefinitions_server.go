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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ScopeAccessReviewScheduleDefinitionsServer is a fake server for instances of the armauthorization.ScopeAccessReviewScheduleDefinitionsClient type.
type ScopeAccessReviewScheduleDefinitionsServer struct {
	// CreateOrUpdateByID is the fake for method ScopeAccessReviewScheduleDefinitionsClient.CreateOrUpdateByID
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdateByID func(ctx context.Context, scope string, scheduleDefinitionID string, properties armauthorization.AccessReviewScheduleDefinitionProperties, options *armauthorization.ScopeAccessReviewScheduleDefinitionsClientCreateOrUpdateByIDOptions) (resp azfake.Responder[armauthorization.ScopeAccessReviewScheduleDefinitionsClientCreateOrUpdateByIDResponse], errResp azfake.ErrorResponder)

	// DeleteByID is the fake for method ScopeAccessReviewScheduleDefinitionsClient.DeleteByID
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteByID func(ctx context.Context, scope string, scheduleDefinitionID string, options *armauthorization.ScopeAccessReviewScheduleDefinitionsClientDeleteByIDOptions) (resp azfake.Responder[armauthorization.ScopeAccessReviewScheduleDefinitionsClientDeleteByIDResponse], errResp azfake.ErrorResponder)

	// GetByID is the fake for method ScopeAccessReviewScheduleDefinitionsClient.GetByID
	// HTTP status codes to indicate success: http.StatusOK
	GetByID func(ctx context.Context, scope string, scheduleDefinitionID string, options *armauthorization.ScopeAccessReviewScheduleDefinitionsClientGetByIDOptions) (resp azfake.Responder[armauthorization.ScopeAccessReviewScheduleDefinitionsClientGetByIDResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ScopeAccessReviewScheduleDefinitionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(scope string, options *armauthorization.ScopeAccessReviewScheduleDefinitionsClientListOptions) (resp azfake.PagerResponder[armauthorization.ScopeAccessReviewScheduleDefinitionsClientListResponse])

	// Stop is the fake for method ScopeAccessReviewScheduleDefinitionsClient.Stop
	// HTTP status codes to indicate success: http.StatusNoContent
	Stop func(ctx context.Context, scope string, scheduleDefinitionID string, options *armauthorization.ScopeAccessReviewScheduleDefinitionsClientStopOptions) (resp azfake.Responder[armauthorization.ScopeAccessReviewScheduleDefinitionsClientStopResponse], errResp azfake.ErrorResponder)
}

// NewScopeAccessReviewScheduleDefinitionsServerTransport creates a new instance of ScopeAccessReviewScheduleDefinitionsServerTransport with the provided implementation.
// The returned ScopeAccessReviewScheduleDefinitionsServerTransport instance is connected to an instance of armauthorization.ScopeAccessReviewScheduleDefinitionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewScopeAccessReviewScheduleDefinitionsServerTransport(srv *ScopeAccessReviewScheduleDefinitionsServer) *ScopeAccessReviewScheduleDefinitionsServerTransport {
	return &ScopeAccessReviewScheduleDefinitionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armauthorization.ScopeAccessReviewScheduleDefinitionsClientListResponse]](),
	}
}

// ScopeAccessReviewScheduleDefinitionsServerTransport connects instances of armauthorization.ScopeAccessReviewScheduleDefinitionsClient to instances of ScopeAccessReviewScheduleDefinitionsServer.
// Don't use this type directly, use NewScopeAccessReviewScheduleDefinitionsServerTransport instead.
type ScopeAccessReviewScheduleDefinitionsServerTransport struct {
	srv          *ScopeAccessReviewScheduleDefinitionsServer
	newListPager *tracker[azfake.PagerResponder[armauthorization.ScopeAccessReviewScheduleDefinitionsClientListResponse]]
}

// Do implements the policy.Transporter interface for ScopeAccessReviewScheduleDefinitionsServerTransport.
func (s *ScopeAccessReviewScheduleDefinitionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ScopeAccessReviewScheduleDefinitionsClient.CreateOrUpdateByID":
		resp, err = s.dispatchCreateOrUpdateByID(req)
	case "ScopeAccessReviewScheduleDefinitionsClient.DeleteByID":
		resp, err = s.dispatchDeleteByID(req)
	case "ScopeAccessReviewScheduleDefinitionsClient.GetByID":
		resp, err = s.dispatchGetByID(req)
	case "ScopeAccessReviewScheduleDefinitionsClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	case "ScopeAccessReviewScheduleDefinitionsClient.Stop":
		resp, err = s.dispatchStop(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ScopeAccessReviewScheduleDefinitionsServerTransport) dispatchCreateOrUpdateByID(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrUpdateByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateByID not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armauthorization.AccessReviewScheduleDefinitionProperties](req)
	if err != nil {
		return nil, err
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	scheduleDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduleDefinitionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateOrUpdateByID(req.Context(), scopeParam, scheduleDefinitionIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessReviewScheduleDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScopeAccessReviewScheduleDefinitionsServerTransport) dispatchDeleteByID(req *http.Request) (*http.Response, error) {
	if s.srv.DeleteByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteByID not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	scheduleDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduleDefinitionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.DeleteByID(req.Context(), scopeParam, scheduleDefinitionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScopeAccessReviewScheduleDefinitionsServerTransport) dispatchGetByID(req *http.Request) (*http.Response, error) {
	if s.srv.GetByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByID not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	scheduleDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduleDefinitionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetByID(req.Context(), scopeParam, scheduleDefinitionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessReviewScheduleDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScopeAccessReviewScheduleDefinitionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armauthorization.ScopeAccessReviewScheduleDefinitionsClientListOptions
		if filterParam != nil {
			options = &armauthorization.ScopeAccessReviewScheduleDefinitionsClientListOptions{
				Filter: filterParam,
			}
		}
		resp := s.srv.NewListPager(scopeParam, options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armauthorization.ScopeAccessReviewScheduleDefinitionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

func (s *ScopeAccessReviewScheduleDefinitionsServerTransport) dispatchStop(req *http.Request) (*http.Response, error) {
	if s.srv.Stop == nil {
		return nil, &nonRetriableError{errors.New("fake for method Stop not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/accessReviewScheduleDefinitions/(?P<scheduleDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	scheduleDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("scheduleDefinitionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Stop(req.Context(), scopeParam, scheduleDefinitionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}