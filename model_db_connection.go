/*
 * Looker API 3.1 Reference
 *
 * ### Authorization  The Looker API uses Looker **API3** credentials for authorization and access control. Looker admins can create API3 credentials on Looker's **Admin/Users** page. Pass API3 credentials to the **_/login** endpoint to obtain a temporary access_token. Include that access_token in the Authorization header of Looker API requests. For details, see [Looker API Authorization](https://looker.com/docs/r/api/authorization)  ### Client SDKs  The Looker API is a RESTful system that should be usable by any programming language capable of making HTTPS requests. Client SDKs for a variety of programming languages can be generated from the Looker API's Swagger JSON metadata to streamline use of the Looker API in your applications. A client SDK for Ruby is available as an example. For more information, see [Looker API Client SDKs](https://looker.com/docs/r/api/client_sdks)  ### Try It Out!  The 'api-docs' page served by the Looker instance includes 'Try It Out!' buttons for each API method. After logging in with API3 credentials, you can use the \"Try It Out!\" buttons to call the API directly from the documentation page to interactively explore API features and responses.  Note! With great power comes great responsibility: The \"Try It Out!\" button makes API calls to your live Looker instance. Be especially careful with destructive API operations such as `delete_user` or similar. There is no \"undo\" for API operations.  ### Versioning  Future releases of Looker will expand this API release-by-release to securely expose more and more of the core power of Looker to API client applications. API endpoints marked as \"beta\" may receive breaking changes without warning (but we will try to avoid doing that). Stable (non-beta) API endpoints should not receive breaking changes in future releases. For more information, see [Looker API Versioning](https://looker.com/docs/r/api/versioning)  ### In This Release  The following are a few examples of noteworthy items that have changed between API 3.0 and API 3.1. For more comprehensive coverage of API changes, please see the release notes for your Looker release.  ### Examples of new things added in API 3.1 (compared to API 3.0):  * [Dashboard construction](#!/3.1/Dashboard/) APIs * [Themes](#!/3.1/Theme/) and [custom color collections](#!/3.1/ColorCollection) APIs * Create and run [SQL Runner](#!/3.1/Query/run_sql_query) queries * Create and run [merged results](#!/3.1/Query/create_merge_query) queries * Create and modify [dashboard filters](#!/3.1/Dashboard/create_dashboard_filter) * Create and modify [password requirements](#!/3.1/Auth/password_config)  ### Deprecated in API 3.0  The following functions and properties have been deprecated in API 3.0.  They continue to exist and work in API 3.0 for the next several Looker releases but they have not been carried forward to API 3.1:  * Dashboard Prefetch functions * User access_filter functions * User API 1.0 credentials functions * Space.is_root and Space.is_user_root properties. Use Space.is_shared_root and Space.is_users_root instead.  ### Semantic changes in API 3.1:  * [all_looks()](#!/3.1/Look/all_looks) no longer includes soft-deleted looks, matching [all_dashboards()](#!/3.1/Dashboard/all_dashboards) behavior. You can find soft-deleted looks using [search_looks()](#!/3.1/Look/search_looks) with the `deleted` param set to True. * [all_spaces()](#!/3.1/Space/all_spaces) no longer includes duplicate items * [search_users()](#!/3.1/User/search_users) no longer accepts Y,y,1,0,N,n for Boolean params, only \"true\" and \"false\". * For greater client and network compatibility, [render_task_results](#!/3.1/RenderTask/render_task_results) now returns HTTP status **202 Accepted** instead of HTTP status **102 Processing** * [all_running_queries()](#!/3.1/Query/all_running_queries) and [kill_query](#!/3.1/Query/kill_query) functions have moved into the [Query](#!/3.1/Query/) function group.   If you have application code which relies on the old behavior of the APIs above, you may continue using the API 3.0 functions in this Looker release. We strongly suggest you update your code to use API 3.1 analogs as soon as possible.  
 *
 * API version: 3.1.0
 * Contact: support@looker.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package looker
// DbConnection struct for DbConnection
type DbConnection struct {
	// Operations the current user is able to perform on this object
	Can map[string]bool `json:"can,omitempty"`
	// Name of the connection. Also used as the unique identifier
	Name string `json:"name,omitempty"`
	Dialect Dialect `json:"dialect,omitempty"`
	// SQL Runner snippets for this connection
	Snippets []Snippet `json:"snippets,omitempty"`
	// Host name/address of server
	Host *string `json:"host,omitempty"`
	// Port number on server
	Port *string `json:"port,omitempty"`
	// Username for server authentication
	Username *string `json:"username,omitempty"`
	// (Write-Only) Password for server authentication
	Password *string `json:"password,omitempty"`
	// Whether the connection uses OAuth for authentication.
	UsesOauth bool `json:"uses_oauth,omitempty"`
	// (Write-Only) Base64 encoded Certificate body for server authentication (when appropriate for dialect).
	Certificate *string `json:"certificate,omitempty"`
	// (Write-Only) Certificate keyfile type - .json or .p12
	FileType *string `json:"file_type,omitempty"`
	// Database name
	Database *string `json:"database,omitempty"`
	// Time zone of database
	DbTimezone *string `json:"db_timezone,omitempty"`
	// Timezone to use in queries
	QueryTimezone *string `json:"query_timezone,omitempty"`
	// Scheme name
	Schema *string `json:"schema,omitempty"`
	// Maximum number of concurrent connection to use
	MaxConnections *int64 `json:"max_connections,omitempty"`
	// Maximum size of query in GBs (BigQuery only, can be a user_attribute name)
	MaxBillingGigabytes *string `json:"max_billing_gigabytes,omitempty"`
	// Use SSL/TLS when connecting to server
	Ssl bool `json:"ssl,omitempty"`
	// Verify the SSL
	VerifySsl bool `json:"verify_ssl,omitempty"`
	// Name of temporary database (if used)
	TmpDbName *string `json:"tmp_db_name,omitempty"`
	// Additional params to add to JDBC connection string
	JdbcAdditionalParams *string `json:"jdbc_additional_params,omitempty"`
	// Connection Pool Timeout, in seconds
	PoolTimeout *int64 `json:"pool_timeout,omitempty"`
	// (Read/Write) SQL Dialect name
	DialectName *string `json:"dialect_name,omitempty"`
	// Creation date for this connection
	CreatedAt *string `json:"created_at,omitempty"`
	// Id of user who last modified this connection configuration
	UserId *string `json:"user_id,omitempty"`
	// Is this an example connection?
	Example bool `json:"example,omitempty"`
	// (Limited access feature) Are per user db credentials enabled. Enabling will remove previously set username and password
	UserDbCredentials *bool `json:"user_db_credentials,omitempty"`
	// Fields whose values map to user attribute names
	UserAttributeFields *[]string `json:"user_attribute_fields,omitempty"`
	// Cron string specifying when maintenance such as PDT trigger checks and drops should be performed
	MaintenanceCron *string `json:"maintenance_cron,omitempty"`
	// Unix timestamp at start of last completed PDT trigger check process
	LastRegenAt *string `json:"last_regen_at,omitempty"`
	// Unix timestamp at start of last completed PDT reap process
	LastReapAt *string `json:"last_reap_at,omitempty"`
	// Precache tables in the SQL Runner
	SqlRunnerPrecacheTables bool `json:"sql_runner_precache_tables,omitempty"`
	// SQL statements (semicolon separated) to issue after connecting to the database. Requires `custom_after_connect_statements` license feature
	AfterConnectStatements *string `json:"after_connect_statements,omitempty"`
	PdtContextOverride DbConnectionOverride `json:"pdt_context_override,omitempty"`
	// Is this connection created and managed by Looker
	Managed bool `json:"managed,omitempty"`
}