# DataActionFormField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] [readonly] 
**Label** | Pointer to **string** | Human-readable label | [optional] [readonly] 
**Description** | Pointer to **string** | Description of field | [optional] [readonly] 
**Type** | Pointer to **string** | Type of field. | [optional] [readonly] 
**Default** | Pointer to **string** | Default value of the field. | [optional] [readonly] 
**OauthUrl** | Pointer to **string** | The URL for an oauth link, if type is &#39;oauth_link&#39;. | [optional] [readonly] 
**Interactive** | **bool** | Whether or not a field supports interactive forms. | [optional] [readonly] 
**Required** | **bool** | Whether or not the field is required. This is a user-interface hint. A user interface displaying this form should not submit it without a value for this field. The action server must also perform this validation. | [optional] [readonly] 
**Options** | Pointer to [**[]DataActionFormSelectOption**](DataActionFormSelectOption.md) | If the form type is &#39;select&#39;, a list of options to be selected from. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


