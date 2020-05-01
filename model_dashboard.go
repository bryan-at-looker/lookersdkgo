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

import (
	"time"
)

// Dashboard struct for Dashboard
type Dashboard struct {
	// Operations the current user is able to perform on this object
	Can map[string]bool `json:"can,omitempty"`
	// Content Favorite Id
	ContentFavoriteId *int64 `json:"content_favorite_id,omitempty"`
	// Id of content metadata
	ContentMetadataId *int64 `json:"content_metadata_id,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
	// Is Hidden
	Hidden bool `json:"hidden,omitempty"`
	// Unique Id
	Id    string    `json:"id,omitempty"`
	Model LookModel `json:"model,omitempty"`
	// Timezone in which the Dashboard will run by default.
	QueryTimezone *string `json:"query_timezone,omitempty"`
	// Is Read-only
	Readonly bool `json:"readonly,omitempty"`
	// Refresh Interval, as a time duration phrase like \"2 hours 30 minutes\". A number with no time units will be interpreted as whole seconds.
	RefreshInterval *string `json:"refresh_interval,omitempty"`
	// Refresh Interval in milliseconds
	RefreshIntervalToI *int64     `json:"refresh_interval_to_i,omitempty"`
	Space              SpaceBase  `json:"space,omitempty"`
	Folder             FolderBase `json:"folder,omitempty"`
	// Dashboard Title
	Title *string `json:"title,omitempty"`
	// Id of User
	UserId *int64 `json:"user_id,omitempty"`
	// Background color
	BackgroundColor *string `json:"background_color,omitempty"`
	// Time that the Dashboard was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Enables crossfiltering in dashboards - only available in dashboards-next (beta)
	CrossfilterEnabled bool `json:"crossfilter_enabled,omitempty"`
	// Elements
	DashboardElements *[]DashboardElement `json:"dashboard_elements,omitempty"`
	// Filters
	DashboardFilters *[]DashboardFilter `json:"dashboard_filters,omitempty"`
	// Layouts
	DashboardLayouts *[]DashboardLayout `json:"dashboard_layouts,omitempty"`
	// Whether or not a dashboard is 'soft' deleted.
	Deleted bool `json:"deleted,omitempty"`
	// Time that the Dashboard was 'soft' deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Id of User that 'soft' deleted the dashboard.
	DeleterId *int64 `json:"deleter_id,omitempty"`
	// Relative path of URI of LookML file to edit the dashboard (LookML dashboard only).
	EditUri *string `json:"edit_uri,omitempty"`
	// Number of times favorited
	FavoriteCount *int64 `json:"favorite_count,omitempty"`
	// Time the dashboard was last accessed
	LastAccessedAt *time.Time `json:"last_accessed_at,omitempty"`
	// Time last viewed in the Looker web UI
	LastViewedAt *time.Time `json:"last_viewed_at,omitempty"`
	// configuration option that governs how dashboard loading will happen.
	LoadConfiguration *string `json:"load_configuration,omitempty"`
	// Links this dashboard to a particular LookML dashboard such that calling a **sync** operation on that LookML dashboard will update this dashboard to match.
	LookmlLinkId *string `json:"lookml_link_id,omitempty"`
	// Show filters bar.  **Security Note:** This property only affects the *cosmetic* appearance of the dashboard, not a user's ability to access data. Hiding the filters bar does **NOT** prevent users from changing filters by other means. For information on how to set up secure data access control policies, see [Control User Access to Data](https://looker.com/docs/r/api/control-access)
	ShowFiltersBar bool `json:"show_filters_bar,omitempty"`
	// Show title
	ShowTitle bool `json:"show_title,omitempty"`
	// Content Metadata Slug
	Slug *string `json:"slug,omitempty"`
	// Id of Space
	SpaceId *string `json:"space_id,omitempty"`
	// Id of folder
	FolderId *string `json:"folder_id,omitempty"`
	// Color of text on text tiles
	TextTileTextColor *string `json:"text_tile_text_color,omitempty"`
	// Tile background color
	TileBackgroundColor *string `json:"tile_background_color,omitempty"`
	// Tile text color
	TileTextColor *string `json:"tile_text_color,omitempty"`
	// Title color
	TitleColor *string `json:"title_color,omitempty"`
	// Number of times viewed in the Looker web UI
	ViewCount  *int64              `json:"view_count,omitempty"`
	Appearance DashboardAppearance `json:"appearance,omitempty"`
	// The preferred route for viewing this dashboard (ie: dashboards or dashboards-next)
	PreferredViewer *string `json:"preferred_viewer,omitempty"`
}
