# IntegrationParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the parameter. | [optional] 
**Label** | Pointer to **string** | Label of the parameter. | [optional] [readonly] 
**Description** | Pointer to **string** | Short description of the parameter. | [optional] [readonly] 
**Required** | **bool** | Whether the parameter is required to be set to use the destination. If unspecified, this defaults to false. | [optional] [readonly] 
**HasValue** | **bool** | Whether the parameter has a value set. | [optional] [readonly] 
**Value** | Pointer to **string** | The current value of the parameter. Always null if the value is sensitive. When writing, null values will be ignored. Set the value to an empty string to clear it. | [optional] 
**UserAttributeName** | Pointer to **string** | When present, the param&#39;s value comes from this user attribute instead of the &#39;value&#39; parameter. Set to null to use the &#39;value&#39;. | [optional] 
**Sensitive** | **bool** | Whether the parameter contains sensitive data like API credentials. If unspecified, this defaults to true. | [optional] [readonly] 
**PerUser** | **bool** | When true, this parameter must be assigned to a user attribute in the admin panel (instead of a constant value), and that value may be updated by the user as part of the integration flow. | [optional] [readonly] 
**DelegateOauthUrl** | Pointer to **string** | When present, the param represents the oauth url the user will be taken to. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

