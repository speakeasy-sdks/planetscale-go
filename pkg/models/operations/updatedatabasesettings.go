// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UpdateDatabaseSettingsRequestBody struct {
	// Whether or not data branching is allowed on the database
	AllowDataBranching *bool `json:"allow_data_branching,omitempty"`
	// Whether or not to copy migration data to new branches and in deploy requests.
	AutomaticMigrations *bool `json:"automatic_migrations,omitempty"`
	// The default branch of the database
	DefaultBranch *string `json:"default_branch,omitempty"`
	// Whether or not full queries should be collected from the database
	InsightsRawQueries *bool `json:"insights_raw_queries,omitempty"`
	// A migration framework to use on the database
	MigrationFramework *string `json:"migration_framework,omitempty"`
	// Name of table to use as migration table for the database
	MigrationTableName *string `json:"migration_table_name,omitempty"`
	// Notes on the database
	Notes *string `json:"notes,omitempty"`
	// Whether or not the web console can be used on the production branch of the database
	ProductionBranchWebConsole *bool `json:"production_branch_web_console,omitempty"`
	// Whether or not deploy requests must be approved by a database administrator other than the request creator
	RequireApprovalForDeploy *bool `json:"require_approval_for_deploy,omitempty"`
	// Whether or not to limit branch creation to the AWS us-east-1 region.
	RestrictBranchRegion *bool `json:"restrict_branch_region,omitempty"`
}

type UpdateDatabaseSettingsRequest struct {
	RequestBody *UpdateDatabaseSettingsRequestBody `request:"mediaType=application/json"`
	// The name of the database
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// The name of the organization the database belongs to
	Organization string `pathParam:"style=simple,explode=false,name=organization"`
}

type UpdateDatabaseSettingsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Returns the updated database
	UpdateDatabaseSettings200ApplicationJSONObject map[string]map[string]interface{}
}