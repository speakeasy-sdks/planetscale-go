// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetABranchSchemaRequest struct {
	// The name of the database the branch belongs to
	Database string `pathParam:"style=simple,explode=false,name=database"`
	// The name of the branch
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// The name of the organization the branch belongs to
	Organization string `pathParam:"style=simple,explode=false,name=organization"`
	// If provided, the schema for this keyspace is returned
	Keyspace *string `queryParam:"style=form,explode=true,name=keyspace"`
}

type GetABranchSchemaResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Gets the schema for the branch
	GetABranchSchema200ApplicationJSONObject map[string]map[string]interface{}
}
