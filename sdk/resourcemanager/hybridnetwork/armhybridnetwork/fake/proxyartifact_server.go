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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ProxyArtifactServer is a fake server for instances of the armhybridnetwork.ProxyArtifactClient type.
type ProxyArtifactServer struct {
	// NewGetPager is the fake for method ProxyArtifactClient.NewGetPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetPager func(resourceGroupName string, publisherName string, artifactStoreName string, artifactName string, options *armhybridnetwork.ProxyArtifactClientGetOptions) (resp azfake.PagerResponder[armhybridnetwork.ProxyArtifactClientGetResponse])

	// NewListPager is the fake for method ProxyArtifactClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, publisherName string, artifactStoreName string, options *armhybridnetwork.ProxyArtifactClientListOptions) (resp azfake.PagerResponder[armhybridnetwork.ProxyArtifactClientListResponse])

	// BeginUpdateState is the fake for method ProxyArtifactClient.BeginUpdateState
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateState func(ctx context.Context, resourceGroupName string, publisherName string, artifactStoreName string, artifactName string, artifactVersionName string, parameters armhybridnetwork.ArtifactChangeState, options *armhybridnetwork.ProxyArtifactClientBeginUpdateStateOptions) (resp azfake.PollerResponder[armhybridnetwork.ProxyArtifactClientUpdateStateResponse], errResp azfake.ErrorResponder)
}

// NewProxyArtifactServerTransport creates a new instance of ProxyArtifactServerTransport with the provided implementation.
// The returned ProxyArtifactServerTransport instance is connected to an instance of armhybridnetwork.ProxyArtifactClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProxyArtifactServerTransport(srv *ProxyArtifactServer) *ProxyArtifactServerTransport {
	return &ProxyArtifactServerTransport{
		srv:              srv,
		newGetPager:      newTracker[azfake.PagerResponder[armhybridnetwork.ProxyArtifactClientGetResponse]](),
		newListPager:     newTracker[azfake.PagerResponder[armhybridnetwork.ProxyArtifactClientListResponse]](),
		beginUpdateState: newTracker[azfake.PollerResponder[armhybridnetwork.ProxyArtifactClientUpdateStateResponse]](),
	}
}

// ProxyArtifactServerTransport connects instances of armhybridnetwork.ProxyArtifactClient to instances of ProxyArtifactServer.
// Don't use this type directly, use NewProxyArtifactServerTransport instead.
type ProxyArtifactServerTransport struct {
	srv              *ProxyArtifactServer
	newGetPager      *tracker[azfake.PagerResponder[armhybridnetwork.ProxyArtifactClientGetResponse]]
	newListPager     *tracker[azfake.PagerResponder[armhybridnetwork.ProxyArtifactClientListResponse]]
	beginUpdateState *tracker[azfake.PollerResponder[armhybridnetwork.ProxyArtifactClientUpdateStateResponse]]
}

// Do implements the policy.Transporter interface for ProxyArtifactServerTransport.
func (p *ProxyArtifactServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProxyArtifactClient.NewGetPager":
		resp, err = p.dispatchNewGetPager(req)
	case "ProxyArtifactClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "ProxyArtifactClient.BeginUpdateState":
		resp, err = p.dispatchBeginUpdateState(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProxyArtifactServerTransport) dispatchNewGetPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetPager not implemented")}
	}
	newGetPager := p.newGetPager.get(req)
	if newGetPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactStores/(?P<artifactStoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactVersions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		artifactStoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactStoreName")])
		if err != nil {
			return nil, err
		}
		artifactNameParam, err := url.QueryUnescape(qp.Get("artifactName"))
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewGetPager(resourceGroupNameParam, publisherNameParam, artifactStoreNameParam, artifactNameParam, nil)
		newGetPager = &resp
		p.newGetPager.add(req, newGetPager)
		server.PagerResponderInjectNextLinks(newGetPager, req, func(page *armhybridnetwork.ProxyArtifactClientGetResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newGetPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetPager) {
		p.newGetPager.remove(req)
	}
	return resp, nil
}

func (p *ProxyArtifactServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactStores/(?P<artifactStoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifacts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		artifactStoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactStoreName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, publisherNameParam, artifactStoreNameParam, nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armhybridnetwork.ProxyArtifactClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}

func (p *ProxyArtifactServerTransport) dispatchBeginUpdateState(req *http.Request) (*http.Response, error) {
	if p.srv.BeginUpdateState == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateState not implemented")}
	}
	beginUpdateState := p.beginUpdateState.get(req)
	if beginUpdateState == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridNetwork/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactStores/(?P<artifactStoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifactVersions/(?P<artifactVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[armhybridnetwork.ArtifactChangeState](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		publisherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
		if err != nil {
			return nil, err
		}
		artifactStoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactStoreName")])
		if err != nil {
			return nil, err
		}
		artifactNameParam, err := url.QueryUnescape(qp.Get("artifactName"))
		if err != nil {
			return nil, err
		}
		artifactVersionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("artifactVersionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginUpdateState(req.Context(), resourceGroupNameParam, publisherNameParam, artifactStoreNameParam, artifactNameParam, artifactVersionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateState = &respr
		p.beginUpdateState.add(req, beginUpdateState)
	}

	resp, err := server.PollerResponderNext(beginUpdateState, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginUpdateState.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateState) {
		p.beginUpdateState.remove(req)
	}

	return resp, nil
}