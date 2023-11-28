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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ReplicationRecoveryPlansServer is a fake server for instances of the armrecoveryservicessiterecovery.ReplicationRecoveryPlansClient type.
type ReplicationRecoveryPlansServer struct {
	// BeginCreate is the fake for method ReplicationRecoveryPlansClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceName string, resourceGroupName string, recoveryPlanName string, input armrecoveryservicessiterecovery.CreateRecoveryPlanInput, options *armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientBeginCreateOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ReplicationRecoveryPlansClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceName string, resourceGroupName string, recoveryPlanName string, options *armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientBeginDeleteOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginFailoverCancel is the fake for method ReplicationRecoveryPlansClient.BeginFailoverCancel
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginFailoverCancel func(ctx context.Context, resourceName string, resourceGroupName string, recoveryPlanName string, options *armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientBeginFailoverCancelOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientFailoverCancelResponse], errResp azfake.ErrorResponder)

	// BeginFailoverCommit is the fake for method ReplicationRecoveryPlansClient.BeginFailoverCommit
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginFailoverCommit func(ctx context.Context, resourceName string, resourceGroupName string, recoveryPlanName string, options *armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientBeginFailoverCommitOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientFailoverCommitResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ReplicationRecoveryPlansClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceName string, resourceGroupName string, recoveryPlanName string, options *armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientGetOptions) (resp azfake.Responder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ReplicationRecoveryPlansClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceName string, resourceGroupName string, options *armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientListOptions) (resp azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientListResponse])

	// BeginPlannedFailover is the fake for method ReplicationRecoveryPlansClient.BeginPlannedFailover
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPlannedFailover func(ctx context.Context, resourceName string, resourceGroupName string, recoveryPlanName string, input armrecoveryservicessiterecovery.RecoveryPlanPlannedFailoverInput, options *armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientBeginPlannedFailoverOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientPlannedFailoverResponse], errResp azfake.ErrorResponder)

	// BeginReprotect is the fake for method ReplicationRecoveryPlansClient.BeginReprotect
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginReprotect func(ctx context.Context, resourceName string, resourceGroupName string, recoveryPlanName string, options *armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientBeginReprotectOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientReprotectResponse], errResp azfake.ErrorResponder)

	// BeginTestFailover is the fake for method ReplicationRecoveryPlansClient.BeginTestFailover
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginTestFailover func(ctx context.Context, resourceName string, resourceGroupName string, recoveryPlanName string, input armrecoveryservicessiterecovery.RecoveryPlanTestFailoverInput, options *armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientBeginTestFailoverOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientTestFailoverResponse], errResp azfake.ErrorResponder)

	// BeginTestFailoverCleanup is the fake for method ReplicationRecoveryPlansClient.BeginTestFailoverCleanup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginTestFailoverCleanup func(ctx context.Context, resourceName string, resourceGroupName string, recoveryPlanName string, input armrecoveryservicessiterecovery.RecoveryPlanTestFailoverCleanupInput, options *armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientBeginTestFailoverCleanupOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientTestFailoverCleanupResponse], errResp azfake.ErrorResponder)

	// BeginUnplannedFailover is the fake for method ReplicationRecoveryPlansClient.BeginUnplannedFailover
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUnplannedFailover func(ctx context.Context, resourceName string, resourceGroupName string, recoveryPlanName string, input armrecoveryservicessiterecovery.RecoveryPlanUnplannedFailoverInput, options *armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientBeginUnplannedFailoverOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientUnplannedFailoverResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method ReplicationRecoveryPlansClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceName string, resourceGroupName string, recoveryPlanName string, input armrecoveryservicessiterecovery.UpdateRecoveryPlanInput, options *armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientBeginUpdateOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewReplicationRecoveryPlansServerTransport creates a new instance of ReplicationRecoveryPlansServerTransport with the provided implementation.
// The returned ReplicationRecoveryPlansServerTransport instance is connected to an instance of armrecoveryservicessiterecovery.ReplicationRecoveryPlansClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReplicationRecoveryPlansServerTransport(srv *ReplicationRecoveryPlansServer) *ReplicationRecoveryPlansServerTransport {
	return &ReplicationRecoveryPlansServerTransport{
		srv:                      srv,
		beginCreate:              newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientCreateResponse]](),
		beginDelete:              newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientDeleteResponse]](),
		beginFailoverCancel:      newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientFailoverCancelResponse]](),
		beginFailoverCommit:      newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientFailoverCommitResponse]](),
		newListPager:             newTracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientListResponse]](),
		beginPlannedFailover:     newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientPlannedFailoverResponse]](),
		beginReprotect:           newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientReprotectResponse]](),
		beginTestFailover:        newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientTestFailoverResponse]](),
		beginTestFailoverCleanup: newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientTestFailoverCleanupResponse]](),
		beginUnplannedFailover:   newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientUnplannedFailoverResponse]](),
		beginUpdate:              newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientUpdateResponse]](),
	}
}

// ReplicationRecoveryPlansServerTransport connects instances of armrecoveryservicessiterecovery.ReplicationRecoveryPlansClient to instances of ReplicationRecoveryPlansServer.
// Don't use this type directly, use NewReplicationRecoveryPlansServerTransport instead.
type ReplicationRecoveryPlansServerTransport struct {
	srv                      *ReplicationRecoveryPlansServer
	beginCreate              *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientCreateResponse]]
	beginDelete              *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientDeleteResponse]]
	beginFailoverCancel      *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientFailoverCancelResponse]]
	beginFailoverCommit      *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientFailoverCommitResponse]]
	newListPager             *tracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientListResponse]]
	beginPlannedFailover     *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientPlannedFailoverResponse]]
	beginReprotect           *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientReprotectResponse]]
	beginTestFailover        *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientTestFailoverResponse]]
	beginTestFailoverCleanup *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientTestFailoverCleanupResponse]]
	beginUnplannedFailover   *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientUnplannedFailoverResponse]]
	beginUpdate              *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ReplicationRecoveryPlansServerTransport.
func (r *ReplicationRecoveryPlansServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ReplicationRecoveryPlansClient.BeginCreate":
		resp, err = r.dispatchBeginCreate(req)
	case "ReplicationRecoveryPlansClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "ReplicationRecoveryPlansClient.BeginFailoverCancel":
		resp, err = r.dispatchBeginFailoverCancel(req)
	case "ReplicationRecoveryPlansClient.BeginFailoverCommit":
		resp, err = r.dispatchBeginFailoverCommit(req)
	case "ReplicationRecoveryPlansClient.Get":
		resp, err = r.dispatchGet(req)
	case "ReplicationRecoveryPlansClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	case "ReplicationRecoveryPlansClient.BeginPlannedFailover":
		resp, err = r.dispatchBeginPlannedFailover(req)
	case "ReplicationRecoveryPlansClient.BeginReprotect":
		resp, err = r.dispatchBeginReprotect(req)
	case "ReplicationRecoveryPlansClient.BeginTestFailover":
		resp, err = r.dispatchBeginTestFailover(req)
	case "ReplicationRecoveryPlansClient.BeginTestFailoverCleanup":
		resp, err = r.dispatchBeginTestFailoverCleanup(req)
	case "ReplicationRecoveryPlansClient.BeginUnplannedFailover":
		resp, err = r.dispatchBeginUnplannedFailover(req)
	case "ReplicationRecoveryPlansClient.BeginUpdate":
		resp, err = r.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReplicationRecoveryPlansServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := r.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryPlans/(?P<recoveryPlanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicessiterecovery.CreateRecoveryPlanInput](req)
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		recoveryPlanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPlanName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreate(req.Context(), resourceNameParam, resourceGroupNameParam, recoveryPlanNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		r.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		r.beginCreate.remove(req)
	}

	return resp, nil
}

func (r *ReplicationRecoveryPlansServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryPlans/(?P<recoveryPlanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		recoveryPlanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPlanName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceNameParam, resourceGroupNameParam, recoveryPlanNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		r.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		r.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		r.beginDelete.remove(req)
	}

	return resp, nil
}

func (r *ReplicationRecoveryPlansServerTransport) dispatchBeginFailoverCancel(req *http.Request) (*http.Response, error) {
	if r.srv.BeginFailoverCancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginFailoverCancel not implemented")}
	}
	beginFailoverCancel := r.beginFailoverCancel.get(req)
	if beginFailoverCancel == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryPlans/(?P<recoveryPlanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/failoverCancel`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		recoveryPlanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPlanName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginFailoverCancel(req.Context(), resourceNameParam, resourceGroupNameParam, recoveryPlanNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginFailoverCancel = &respr
		r.beginFailoverCancel.add(req, beginFailoverCancel)
	}

	resp, err := server.PollerResponderNext(beginFailoverCancel, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginFailoverCancel.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginFailoverCancel) {
		r.beginFailoverCancel.remove(req)
	}

	return resp, nil
}

func (r *ReplicationRecoveryPlansServerTransport) dispatchBeginFailoverCommit(req *http.Request) (*http.Response, error) {
	if r.srv.BeginFailoverCommit == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginFailoverCommit not implemented")}
	}
	beginFailoverCommit := r.beginFailoverCommit.get(req)
	if beginFailoverCommit == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryPlans/(?P<recoveryPlanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/failoverCommit`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		recoveryPlanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPlanName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginFailoverCommit(req.Context(), resourceNameParam, resourceGroupNameParam, recoveryPlanNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginFailoverCommit = &respr
		r.beginFailoverCommit.add(req, beginFailoverCommit)
	}

	resp, err := server.PollerResponderNext(beginFailoverCommit, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginFailoverCommit.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginFailoverCommit) {
		r.beginFailoverCommit.remove(req)
	}

	return resp, nil
}

func (r *ReplicationRecoveryPlansServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryPlans/(?P<recoveryPlanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	recoveryPlanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPlanName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceNameParam, resourceGroupNameParam, recoveryPlanNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RecoveryPlan, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReplicationRecoveryPlansServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryPlans`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListPager(resourceNameParam, resourceGroupNameParam, nil)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armrecoveryservicessiterecovery.ReplicationRecoveryPlansClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

func (r *ReplicationRecoveryPlansServerTransport) dispatchBeginPlannedFailover(req *http.Request) (*http.Response, error) {
	if r.srv.BeginPlannedFailover == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPlannedFailover not implemented")}
	}
	beginPlannedFailover := r.beginPlannedFailover.get(req)
	if beginPlannedFailover == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryPlans/(?P<recoveryPlanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/plannedFailover`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicessiterecovery.RecoveryPlanPlannedFailoverInput](req)
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		recoveryPlanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPlanName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginPlannedFailover(req.Context(), resourceNameParam, resourceGroupNameParam, recoveryPlanNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPlannedFailover = &respr
		r.beginPlannedFailover.add(req, beginPlannedFailover)
	}

	resp, err := server.PollerResponderNext(beginPlannedFailover, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginPlannedFailover.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPlannedFailover) {
		r.beginPlannedFailover.remove(req)
	}

	return resp, nil
}

func (r *ReplicationRecoveryPlansServerTransport) dispatchBeginReprotect(req *http.Request) (*http.Response, error) {
	if r.srv.BeginReprotect == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginReprotect not implemented")}
	}
	beginReprotect := r.beginReprotect.get(req)
	if beginReprotect == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryPlans/(?P<recoveryPlanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reProtect`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		recoveryPlanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPlanName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginReprotect(req.Context(), resourceNameParam, resourceGroupNameParam, recoveryPlanNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginReprotect = &respr
		r.beginReprotect.add(req, beginReprotect)
	}

	resp, err := server.PollerResponderNext(beginReprotect, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginReprotect.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginReprotect) {
		r.beginReprotect.remove(req)
	}

	return resp, nil
}

func (r *ReplicationRecoveryPlansServerTransport) dispatchBeginTestFailover(req *http.Request) (*http.Response, error) {
	if r.srv.BeginTestFailover == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginTestFailover not implemented")}
	}
	beginTestFailover := r.beginTestFailover.get(req)
	if beginTestFailover == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryPlans/(?P<recoveryPlanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/testFailover`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicessiterecovery.RecoveryPlanTestFailoverInput](req)
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		recoveryPlanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPlanName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginTestFailover(req.Context(), resourceNameParam, resourceGroupNameParam, recoveryPlanNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginTestFailover = &respr
		r.beginTestFailover.add(req, beginTestFailover)
	}

	resp, err := server.PollerResponderNext(beginTestFailover, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginTestFailover.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginTestFailover) {
		r.beginTestFailover.remove(req)
	}

	return resp, nil
}

func (r *ReplicationRecoveryPlansServerTransport) dispatchBeginTestFailoverCleanup(req *http.Request) (*http.Response, error) {
	if r.srv.BeginTestFailoverCleanup == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginTestFailoverCleanup not implemented")}
	}
	beginTestFailoverCleanup := r.beginTestFailoverCleanup.get(req)
	if beginTestFailoverCleanup == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryPlans/(?P<recoveryPlanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/testFailoverCleanup`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicessiterecovery.RecoveryPlanTestFailoverCleanupInput](req)
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		recoveryPlanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPlanName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginTestFailoverCleanup(req.Context(), resourceNameParam, resourceGroupNameParam, recoveryPlanNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginTestFailoverCleanup = &respr
		r.beginTestFailoverCleanup.add(req, beginTestFailoverCleanup)
	}

	resp, err := server.PollerResponderNext(beginTestFailoverCleanup, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginTestFailoverCleanup.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginTestFailoverCleanup) {
		r.beginTestFailoverCleanup.remove(req)
	}

	return resp, nil
}

func (r *ReplicationRecoveryPlansServerTransport) dispatchBeginUnplannedFailover(req *http.Request) (*http.Response, error) {
	if r.srv.BeginUnplannedFailover == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUnplannedFailover not implemented")}
	}
	beginUnplannedFailover := r.beginUnplannedFailover.get(req)
	if beginUnplannedFailover == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryPlans/(?P<recoveryPlanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/unplannedFailover`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicessiterecovery.RecoveryPlanUnplannedFailoverInput](req)
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		recoveryPlanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPlanName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginUnplannedFailover(req.Context(), resourceNameParam, resourceGroupNameParam, recoveryPlanNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUnplannedFailover = &respr
		r.beginUnplannedFailover.add(req, beginUnplannedFailover)
	}

	resp, err := server.PollerResponderNext(beginUnplannedFailover, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginUnplannedFailover.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUnplannedFailover) {
		r.beginUnplannedFailover.remove(req)
	}

	return resp, nil
}

func (r *ReplicationRecoveryPlansServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := r.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationRecoveryPlans/(?P<recoveryPlanName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicessiterecovery.UpdateRecoveryPlanInput](req)
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		recoveryPlanNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPlanName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginUpdate(req.Context(), resourceNameParam, resourceGroupNameParam, recoveryPlanNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		r.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		r.beginUpdate.remove(req)
	}

	return resp, nil
}