# GitBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Name** | Pointer to **string** | The short name on the local. Updating &#x60;name&#x60; results in &#x60;git checkout &lt;new_name&gt;&#x60; | [optional] 
**Remote** | Pointer to **string** | The name of the remote | [optional] [readonly] 
**RemoteName** | Pointer to **string** | The short name on the remote | [optional] [readonly] 
**Error** | Pointer to **string** | Name of error | [optional] [readonly] 
**Message** | Pointer to **string** | Message describing an error if present | [optional] [readonly] 
**OwnerName** | Pointer to **string** | Name of the owner of a personal branch | [optional] [readonly] 
**Readonly** | **bool** | Whether or not this branch is readonly | [optional] [readonly] 
**Personal** | **bool** | Whether or not this branch is a personal branch - readonly for all developers except the owner | [optional] [readonly] 
**IsLocal** | **bool** | Whether or not a local ref exists for the branch | [optional] [readonly] 
**IsRemote** | **bool** | Whether or not a remote ref exists for the branch | [optional] [readonly] 
**IsProduction** | **bool** | Whether or not this is the production branch | [optional] [readonly] 
**AheadCount** | Pointer to **int64** | Number of commits the local branch is ahead of the remote | [optional] [readonly] 
**BehindCount** | Pointer to **int64** | Number of commits the local branch is behind the remote | [optional] [readonly] 
**CommitAt** | Pointer to **int64** | UNIX timestamp at which this branch was last committed. | [optional] [readonly] 
**Ref** | Pointer to **string** | The resolved ref of this branch. Updating &#x60;ref&#x60; results in &#x60;git reset --hard &lt;new_ref&gt;&#x60;&#x60;. | [optional] 
**RemoteRef** | Pointer to **string** | The resolved ref of this branch remote. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


