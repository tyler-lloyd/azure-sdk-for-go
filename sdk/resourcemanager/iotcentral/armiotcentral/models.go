//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armiotcentral

import "time"

// App - The IoT Central application.
type App struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// REQUIRED; A valid instance SKU.
	SKU *AppSKUInfo `json:"sku,omitempty"`

	// The managed identities for the IoT Central application.
	Identity *SystemAssignedServiceIdentity `json:"identity,omitempty"`

	// The common properties of an IoT Central application.
	Properties *AppProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AppAvailabilityInfo - The properties indicating whether a given IoT Central application name or subdomain is available.
type AppAvailabilityInfo struct {
	// READ-ONLY; The detailed reason message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The value which indicates whether the provided name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty" azure:"ro"`

	// READ-ONLY; The reason for unavailability.
	Reason *string `json:"reason,omitempty" azure:"ro"`
}

// AppListResult - A list of IoT Central Applications with a next link.
type AppListResult struct {
	// The link used to get the next page of IoT Central Applications.
	NextLink *string `json:"nextLink,omitempty"`

	// A list of IoT Central Applications.
	Value []*App `json:"value,omitempty"`
}

// AppPatch - The description of the IoT Central application.
type AppPatch struct {
	// The managed identities for the IoT Central application.
	Identity *SystemAssignedServiceIdentity `json:"identity,omitempty"`

	// The common properties of an IoT Central application.
	Properties *AppProperties `json:"properties,omitempty"`

	// A valid instance SKU.
	SKU *AppSKUInfo `json:"sku,omitempty"`

	// Instance tags
	Tags map[string]*string `json:"tags,omitempty"`
}

// AppProperties - The properties of an IoT Central application.
type AppProperties struct {
	// The display name of the application.
	DisplayName *string `json:"displayName,omitempty"`

	// Network Rule Set Properties of this IoT Central application.
	NetworkRuleSets *NetworkRuleSets `json:"networkRuleSets,omitempty"`

	// Whether requests from the public network are allowed.
	PublicNetworkAccess *PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// The subdomain of the application.
	Subdomain *string `json:"subdomain,omitempty"`

	// The ID of the application template, which is a blueprint that defines the characteristics and behaviors of an application.
	// Optional; if not specified, defaults to a blank blueprint and allows the
	// application to be defined from scratch.
	Template *string `json:"template,omitempty"`

	// READ-ONLY; The ID of the application.
	ApplicationID *string `json:"applicationId,omitempty" azure:"ro"`

	// READ-ONLY; Private endpoint connections created on this IoT Central application.
	PrivateEndpointConnections []*PrivateEndpointConnection `json:"privateEndpointConnections,omitempty" azure:"ro"`

	// READ-ONLY; The provisioning state of the application.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The current state of the application.
	State *AppState `json:"state,omitempty" azure:"ro"`
}

// AppSKUInfo - Information about the SKU of the IoT Central application.
type AppSKUInfo struct {
	// REQUIRED; The name of the SKU.
	Name *AppSKU `json:"name,omitempty"`
}

// AppTemplate - IoT Central Application Template.
type AppTemplate struct {
	// READ-ONLY; The description of the template.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The industry of the template.
	Industry *string `json:"industry,omitempty" azure:"ro"`

	// READ-ONLY; A list of locations that support the template.
	Locations []*AppTemplateLocations `json:"locations,omitempty" azure:"ro"`

	// READ-ONLY; The ID of the template.
	ManifestID *string `json:"manifestId,omitempty" azure:"ro"`

	// READ-ONLY; The version of the template.
	ManifestVersion *string `json:"manifestVersion,omitempty" azure:"ro"`

	// READ-ONLY; The name of the template.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The order of the template in the templates list.
	Order *float32 `json:"order,omitempty" azure:"ro"`

	// READ-ONLY; The title of the template.
	Title *string `json:"title,omitempty" azure:"ro"`
}

// AppTemplateLocations - IoT Central Application Template Locations.
type AppTemplateLocations struct {
	// READ-ONLY; The display name of the location.
	DisplayName *string `json:"displayName,omitempty" azure:"ro"`

	// READ-ONLY; The ID of the location.
	ID *string `json:"id,omitempty" azure:"ro"`
}

// AppTemplatesResult - A list of IoT Central Application Templates with a next link.
type AppTemplatesResult struct {
	// The link used to get the next page of IoT Central application templates.
	NextLink *string `json:"nextLink,omitempty"`

	// READ-ONLY; A list of IoT Central Application Templates.
	Value []*AppTemplate `json:"value,omitempty" azure:"ro"`
}

// AppsClientBeginCreateOrUpdateOptions contains the optional parameters for the AppsClient.BeginCreateOrUpdate method.
type AppsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AppsClientBeginDeleteOptions contains the optional parameters for the AppsClient.BeginDelete method.
type AppsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AppsClientBeginUpdateOptions contains the optional parameters for the AppsClient.BeginUpdate method.
type AppsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AppsClientCheckNameAvailabilityOptions contains the optional parameters for the AppsClient.CheckNameAvailability method.
type AppsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// AppsClientCheckSubdomainAvailabilityOptions contains the optional parameters for the AppsClient.CheckSubdomainAvailability
// method.
type AppsClientCheckSubdomainAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// AppsClientGetOptions contains the optional parameters for the AppsClient.Get method.
type AppsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AppsClientListByResourceGroupOptions contains the optional parameters for the AppsClient.ListByResourceGroup method.
type AppsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AppsClientListBySubscriptionOptions contains the optional parameters for the AppsClient.ListBySubscription method.
type AppsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// AppsClientListTemplatesOptions contains the optional parameters for the AppsClient.ListTemplates method.
type AppsClientListTemplatesOptions struct {
	// placeholder for future optional parameters
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// NetworkRuleSetIPRule - An object for an IP range that will be allowed access.
type NetworkRuleSetIPRule struct {
	// The readable name of the IP rule.
	FilterName *string `json:"filterName,omitempty"`

	// The CIDR block defining the IP range.
	IPMask *string `json:"ipMask,omitempty"`

	// READ-ONLY; The network action for the IP mask.
	Action *IPRuleAction `json:"action,omitempty" azure:"ro"`
}

// NetworkRuleSets - Network Rule Set Properties of this IoT Central application.
type NetworkRuleSets struct {
	// Whether these rules apply for device connectivity to IoT Hub and Device Provisioning service associated with this application.
	ApplyToDevices *bool `json:"applyToDevices,omitempty"`

	// Whether these rules apply for connectivity via IoT Central web portal and APIs.
	ApplyToIoTCentral *bool `json:"applyToIoTCentral,omitempty"`

	// The default network action to apply.
	DefaultAction *NetworkAction `json:"defaultAction,omitempty"`

	// List of IP rules.
	IPRules []*NetworkRuleSetIPRule `json:"ipRules,omitempty"`
}

// Operation - IoT Central REST API operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Operation name: {provider}/{resource}/{read | write | action | delete}
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation.
	Origin *string `json:"origin,omitempty" azure:"ro"`

	// READ-ONLY; Additional descriptions for the operation.
	Properties interface{} `json:"properties,omitempty" azure:"ro"`
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// READ-ONLY; Friendly description for the operation,
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; Name of the operation
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; Service provider: Microsoft IoT Central
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; Resource Type: IoT Central
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationInputs - Input values.
type OperationInputs struct {
	// REQUIRED; The name of the IoT Central application instance to check.
	Name *string `json:"name,omitempty"`

	// The type of the IoT Central resource to query.
	Type *string `json:"type,omitempty"`
}

// OperationListResult - A list of IoT Central operations. It contains a list of operations and a URL link to get the next
// set of results.
type OperationListResult struct {
	// The link used to get the next page of IoT Central description objects.
	NextLink *string `json:"nextLink,omitempty"`

	// READ-ONLY; A list of operations supported by the Microsoft.IoTCentral resource provider.
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpoint - The private endpoint resource.
type PrivateEndpoint struct {
	// READ-ONLY; The ARM identifier for private endpoint.
	ID *string `json:"id,omitempty" azure:"ro"`
}

// PrivateEndpointConnection - The private endpoint connection resource.
type PrivateEndpointConnection struct {
	// Resource properties.
	Properties *PrivateEndpointConnectionProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionListResult - List of private endpoint connections associated with the specified resource.
type PrivateEndpointConnectionListResult struct {
	// Array of private endpoint connections.
	Value []*PrivateEndpointConnection `json:"value,omitempty"`
}

// PrivateEndpointConnectionProperties - Properties of the private endpoint connection.
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState,omitempty"`

	// The private endpoint resource.
	PrivateEndpoint *PrivateEndpoint `json:"privateEndpoint,omitempty"`

	// READ-ONLY; The group ids for the private endpoint resource.
	GroupIDs []*string `json:"groupIds,omitempty" azure:"ro"`

	// READ-ONLY; The provisioning state of the private endpoint connection resource.
	ProvisioningState *PrivateEndpointConnectionProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// PrivateEndpointConnectionsClientBeginCreateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginCreate
// method.
type PrivateEndpointConnectionsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
type PrivateEndpointConnectionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListOptions contains the optional parameters for the PrivateEndpointConnectionsClient.List
// method.
type PrivateEndpointConnectionsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResource - A private link resource.
type PrivateLinkResource struct {
	// Resource properties.
	Properties *PrivateLinkResourceProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PrivateLinkResourceListResult - A list of private link resources.
type PrivateLinkResourceListResult struct {
	// Array of private link resources
	Value []*PrivateLinkResource `json:"value,omitempty"`
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// The private link resource private link DNS zone name.
	RequiredZoneNames []*string `json:"requiredZoneNames,omitempty"`

	// READ-ONLY; The private link resource group id.
	GroupID *string `json:"groupId,omitempty" azure:"ro"`

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string `json:"requiredMembers,omitempty" azure:"ro"`
}

// PrivateLinkServiceConnectionState - A collection of information about the state of the connection between service consumer
// and provider.
type PrivateLinkServiceConnectionState struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string `json:"actionsRequired,omitempty"`

	// The reason for approval/rejection of the connection.
	Description *string `json:"description,omitempty"`

	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus `json:"status,omitempty"`
}

// PrivateLinksClientGetOptions contains the optional parameters for the PrivateLinksClient.Get method.
type PrivateLinksClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinksClientListOptions contains the optional parameters for the PrivateLinksClient.List method.
type PrivateLinksClientListOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SystemAssignedServiceIdentity - Managed service identity (either system assigned, or none)
type SystemAssignedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (either system assigned, or none).
	Type *SystemAssignedServiceIdentityType `json:"type,omitempty"`

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}