/*
 * Looker API 3.1 Reference
 *
 * ### Authorization  The Looker API uses Looker **API3** credentials for authorization and access control. Looker admins can create API3 credentials on Looker's **Admin/Users** page. Pass API3 credentials to the **_/login** endpoint to obtain a temporary access_token. Include that access_token in the Authorization header of Looker API requests. For details, see [Looker API Authorization](https://looker.com/docs/r/api/authorization)  ### Client SDKs  The Looker API is a RESTful system that should be usable by any programming language capable of making HTTPS requests. Client SDKs for a variety of programming languages can be generated from the Looker API's Swagger JSON metadata to streamline use of the Looker API in your applications. A client SDK for Ruby is available as an example. For more information, see [Looker API Client SDKs](https://looker.com/docs/r/api/client_sdks)  ### Try It Out!  The 'api-docs' page served by the Looker instance includes 'Try It Out!' buttons for each API method. After logging in with API3 credentials, you can use the \"Try It Out!\" buttons to call the API directly from the documentation page to interactively explore API features and responses.  Note! With great power comes great responsibility: The \"Try It Out!\" button makes API calls to your live Looker instance. Be especially careful with destructive API operations such as `delete_user` or similar. There is no \"undo\" for API operations.  ### Versioning  Future releases of Looker will expand this API release-by-release to securely expose more and more of the core power of Looker to API client applications. API endpoints marked as \"beta\" may receive breaking changes without warning (but we will try to avoid doing that). Stable (non-beta) API endpoints should not receive breaking changes in future releases. For more information, see [Looker API Versioning](https://looker.com/docs/r/api/versioning)  ### In This Release  The following are a few examples of noteworthy items that have changed between API 3.0 and API 3.1. For more comprehensive coverage of API changes, please see the release notes for your Looker release.  ### Examples of new things added in API 3.1 (compared to API 3.0):  * [Dashboard construction](#!/3.1/Dashboard/) APIs * [Themes](#!/3.1/Theme/) and [custom color collections](#!/3.1/ColorCollection) APIs * Create and run [SQL Runner](#!/3.1/Query/run_sql_query) queries * Create and run [merged results](#!/3.1/Query/create_merge_query) queries * Create and modify [dashboard filters](#!/3.1/Dashboard/create_dashboard_filter) * Create and modify [password requirements](#!/3.1/Auth/password_config)  ### Deprecated in API 3.0  The following functions and properties have been deprecated in API 3.0.  They continue to exist and work in API 3.0 for the next several Looker releases but they have not been carried forward to API 3.1:  * Dashboard Prefetch functions * User access_filter functions * User API 1.0 credentials functions * Space.is_root and Space.is_user_root properties. Use Space.is_shared_root and Space.is_users_root instead.  ### Semantic changes in API 3.1:  * [all_looks()](#!/3.1/Look/all_looks) no longer includes soft-deleted looks, matching [all_dashboards()](#!/3.1/Dashboard/all_dashboards) behavior. You can find soft-deleted looks using [search_looks()](#!/3.1/Look/search_looks) with the `deleted` param set to True. * [all_spaces()](#!/3.1/Space/all_spaces) no longer includes duplicate items * [search_users()](#!/3.1/User/search_users) no longer accepts Y,y,1,0,N,n for Boolean params, only \"true\" and \"false\". * For greater client and network compatibility, [render_task_results](#!/3.1/RenderTask/render_task_results) now returns HTTP status **202 Accepted** instead of HTTP status **102 Processing** * [all_running_queries()](#!/3.1/Query/all_running_queries) and [kill_query](#!/3.1/Query/kill_query) functions have moved into the [Query](#!/3.1/Query/) function group.   If you have application code which relies on the old behavior of the APIs above, you may continue using the API 3.0 functions in this Looker release. We strongly suggest you update your code to use API 3.1 analogs as soon as possible.
 *
 * API version: 3.1.0
 * Contact: support@looker.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package lookersdkgo

// User struct for User
type User struct {
	// Operations the current user is able to perform on this object
	Can map[string]bool `json:"can,omitempty"`
	// URL for the avatar image (may be generic)
	AvatarUrl *string `json:"avatar_url,omitempty"`
	// URL for the avatar image (may be generic), does not specify size
	AvatarUrlWithoutSizing *string `json:"avatar_url_without_sizing,omitempty"`
	// API 3 credentials
	CredentialsApi3  *[]CredentialsApi3 `json:"credentials_api3,omitempty"`
	CredentialsEmail CredentialsEmail   `json:"credentials_email,omitempty"`
	// Embed credentials
	CredentialsEmbed        *[]CredentialsEmbed     `json:"credentials_embed,omitempty"`
	CredentialsGoogle       CredentialsGoogle       `json:"credentials_google,omitempty"`
	CredentialsLdap         CredentialsLdap         `json:"credentials_ldap,omitempty"`
	CredentialsLookerOpenid CredentialsLookerOpenid `json:"credentials_looker_openid,omitempty"`
	CredentialsOidc         CredentialsOidc         `json:"credentials_oidc,omitempty"`
	CredentialsSaml         CredentialsSaml         `json:"credentials_saml,omitempty"`
	CredentialsTotp         CredentialsTotp         `json:"credentials_totp,omitempty"`
	// Full name for display (available only if both first_name and last_name are set)
	DisplayName *string `json:"display_name,omitempty"`
	// EMail address
	Email *string `json:"email,omitempty"`
	// (Embed only) ID of user's group space based on the external_group_id optionally specified during embed user login
	EmbedGroupSpaceId *int64 `json:"embed_group_space_id,omitempty"`
	// First name
	FirstName *string `json:"first_name,omitempty"`
	// Array of ids of the groups for this user
	GroupIds *[]int64 `json:"group_ids,omitempty"`
	// ID string for user's home space
	HomeSpaceId *string `json:"home_space_id,omitempty"`
	// ID string for user's home folder
	HomeFolderId *string `json:"home_folder_id,omitempty"`
	// Unique Id
	Id int64 `json:"id,omitempty"`
	// Account has been disabled
	IsDisabled bool `json:"is_disabled,omitempty"`
	// Last name
	LastName *string `json:"last_name,omitempty"`
	// User's preferred locale. User locale takes precedence over Looker's system-wide default locale. Locale determines language of display strings and date and numeric formatting in API responses. Locale string must be a 2 letter language code or a combination of language code and region code: 'en' or 'en-US', for example.
	Locale *string `json:"locale,omitempty"`
	// Array of strings representing the Looker versions that this user has used (this only goes back as far as '3.54.0')
	LookerVersions *[]string `json:"looker_versions,omitempty"`
	// User's dev workspace has been checked for presence of applicable production projects
	ModelsDirValidated *bool `json:"models_dir_validated,omitempty"`
	// ID of user's personal space
	PersonalSpaceId *int64 `json:"personal_space_id,omitempty"`
	// ID of user's personal folder
	PersonalFolderId *int64 `json:"personal_folder_id,omitempty"`
	// User is identified as an employee of Looker
	PresumedLookerEmployee bool `json:"presumed_looker_employee,omitempty"`
	// Array of ids of the roles for this user
	RoleIds *[]int64 `json:"role_ids,omitempty"`
	// Active sessions
	Sessions *[]Session `json:"sessions,omitempty"`
	// Per user dictionary of undocumented state information owned by the Looker UI.
	UiState *map[string]string `json:"ui_state,omitempty"`
	// User is identified as an employee of Looker who has been verified via Looker corporate authentication
	VerifiedLookerEmployee bool `json:"verified_looker_employee,omitempty"`
	// User's roles are managed by an external directory like SAML or LDAP and can not be changed directly.
	RolesExternallyManaged bool `json:"roles_externally_managed,omitempty"`
	// Link to get this item
	Url *string `json:"url,omitempty"`
}
