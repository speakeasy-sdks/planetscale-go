// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ListBranchPasswordsRequest struct {
	// The name of the branch the password belongs to
	Branch string `pathParam:"style=simple,explode=false,name=branch"`
	// The name of the database the password belongs to
	Database string `pathParam:"style=simple,explode=false,name=database"`
	// The name of the organization the password belongs to
	Organization string `pathParam:"style=simple,explode=false,name=organization"`
	// If provided, specifies the page offset of returned results
	Page *float64 `queryParam:"style=form,explode=true,name=page"`
	// If provided, specifies the number of returned results
	PerPage *float64 `queryParam:"style=form,explode=true,name=per_page"`
	// A read-only region of the database branch. If present, the password results will be filtered to only those in the region
	ReadOnlyRegionID *string `queryParam:"style=form,explode=true,name=read_only_region_id"`
}

type ListBranchPasswordsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Gets the passwords for the database branch
	ListBranchPasswords200ApplicationJSONObject map[string]map[string]interface{}
}
