# CredentialsLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | Timestamp for the creation of this credential | [optional] [readonly] 
**Email** | Pointer to **string** | EMail address | [optional] [readonly] 
**IsDisabled** | **bool** | Has this credential been disabled? | [optional] [readonly] 
**LdapDn** | Pointer to **string** | LDAP Distinguished name for this user (as-of the last login) | [optional] [readonly] 
**LdapId** | Pointer to **string** | LDAP Unique ID for this user | [optional] [readonly] 
**LoggedInAt** | Pointer to **string** | Timestamp for most recent login using credential | [optional] [readonly] 
**Type** | Pointer to **string** | Short name for the type of this kind of credential | [optional] [readonly] 
**Url** | Pointer to **string** | Link to get this item | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


