// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ListOrganizationsRequest struct {
	// If provided, specifies the page offset of returned results
	Page *float64 `queryParam:"style=form,explode=true,name=page"`
	// If provided, specifies the number of returned results
	PerPage *float64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListOrganizationsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Gets the organizations for the current user
	ListOrganizations200ApplicationJSONObject map[string]map[string]interface{}
}
