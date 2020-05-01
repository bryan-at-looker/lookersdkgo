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

// WriteScheduledPlan struct for WriteScheduledPlan
type WriteScheduledPlan struct {
	// Name of this scheduled plan
	Name *string `json:"name,omitempty"`
	// User Id which owns this scheduled plan
	UserId *int64 `json:"user_id,omitempty"`
	// Whether schedule is run as recipient (only applicable for email recipients)
	RunAsRecipient bool `json:"run_as_recipient,omitempty"`
	// Whether the ScheduledPlan is enabled
	Enabled bool `json:"enabled,omitempty"`
	// Id of a look
	LookId *int64 `json:"look_id,omitempty"`
	// Id of a dashboard
	DashboardId *int64 `json:"dashboard_id,omitempty"`
	// Id of a LookML dashboard
	LookmlDashboardId *string `json:"lookml_dashboard_id,omitempty"`
	// Query string to run look or dashboard with
	FiltersString *string `json:"filters_string,omitempty"`
	// (DEPRECATED) Alias for filters_string field
	DashboardFilters *string `json:"dashboard_filters,omitempty"`
	// Delivery should occur if running the dashboard or look returns results
	RequireResults bool `json:"require_results,omitempty"`
	// Delivery should occur if the dashboard look does not return results
	RequireNoResults bool `json:"require_no_results,omitempty"`
	// Delivery should occur if data have changed since the last run
	RequireChange bool `json:"require_change,omitempty"`
	// Will run an unlimited query and send all results.
	SendAllResults bool `json:"send_all_results,omitempty"`
	// Vixie-Style crontab specification when to run
	Crontab *string `json:"crontab,omitempty"`
	// Name of a datagroup; if specified will run when datagroup triggered (can't be used with cron string)
	Datagroup *string `json:"datagroup,omitempty"`
	// Timezone for interpreting the specified crontab (default is Looker instance timezone)
	Timezone *string `json:"timezone,omitempty"`
	// Query id
	QueryId *string `json:"query_id,omitempty"`
	// Scheduled plan destinations
	ScheduledPlanDestination *[]ScheduledPlanDestination `json:"scheduled_plan_destination,omitempty"`
	// Whether the plan in question should only be run once (usually for testing)
	RunOnce bool `json:"run_once,omitempty"`
	// Whether links back to Looker should be included in this ScheduledPlan
	IncludeLinks bool `json:"include_links,omitempty"`
	// The size of paper a PDF should be rendered for
	PdfPaperSize *string `json:"pdf_paper_size,omitempty"`
	// Whether the paper should be landscape
	PdfLandscape bool `json:"pdf_landscape,omitempty"`
	// Whether this schedule is in an embed context or not
	Embed bool `json:"embed,omitempty"`
	// Color scheme of the dashboard if applicable
	ColorTheme *string `json:"color_theme,omitempty"`
	// Whether or not to expand table vis to full length
	LongTables bool `json:"long_tables,omitempty"`
}
