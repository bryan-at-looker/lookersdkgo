# Space

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique Name | 
**ParentId** | Pointer to **string** | Id of Parent. If the parent id is null, this is a root-level entry | [optional] 
**Id** | **string** | Unique Id | [optional] [readonly] 
**ContentMetadataId** | Pointer to **int64** | Id of content metadata | [optional] [readonly] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Time the space was created | [optional] [readonly] 
**CreatorId** | Pointer to **int64** | User Id of Creator | [optional] [readonly] 
**ChildCount** | Pointer to **int64** | Children Count | [optional] [readonly] 
**ExternalId** | Pointer to **string** | Embedder&#39;s Id if this space was autogenerated as an embedding shared space via &#39;external_group_id&#39; in an SSO embed login | [optional] [readonly] 
**IsEmbed** | **bool** | Space is an embed space | [optional] [readonly] 
**IsEmbedSharedRoot** | **bool** | Space is the root embed shared space | [optional] [readonly] 
**IsEmbedUsersRoot** | **bool** | Space is the root embed users space | [optional] [readonly] 
**IsPersonal** | **bool** | Space is a user&#39;s personal space | [optional] [readonly] 
**IsPersonalDescendant** | **bool** | Space is descendant of a user&#39;s personal space | [optional] [readonly] 
**IsSharedRoot** | **bool** | Space is the root shared space | [optional] [readonly] 
**IsUsersRoot** | **bool** | Space is the root user space | [optional] [readonly] 
**Can** | **map[string]bool** | Operations the current user is able to perform on this object | [optional] [readonly] 
**Dashboards** | Pointer to [**[]DashboardBase**](DashboardBase.md) | Dashboards | [optional] [readonly] 
**Looks** | Pointer to [**[]LookWithDashboards**](LookWithDashboards.md) | Looks | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


