// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ListDeployRequestsRequest struct {
	// The name of the deploy request's database
	Database string `pathParam:"style=simple,explode=false,name=database"`
	// The name of the deploy request's organization
	Organization string `pathParam:"style=simple,explode=false,name=organization"`
	// If provided, specifies the page offset of returned results
	Page *float64 `queryParam:"style=form,explode=true,name=page"`
	// If provided, specifies the number of returned results
	PerPage *float64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListDeployRequestsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Returns a list of deploy requests for a database
	ListDeployRequests200ApplicationJSONObject map[string]map[string]interface{}
}
