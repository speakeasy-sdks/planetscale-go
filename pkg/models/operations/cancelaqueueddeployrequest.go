// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CancelAQueuedDeployRequestRequest struct {
	// The name of the deploy request's database
	Database string `pathParam:"style=simple,explode=false,name=database"`
	// The number of the deploy request
	Number string `pathParam:"style=simple,explode=false,name=number"`
	// The name of the deploy request's organization
	Organization string `pathParam:"style=simple,explode=false,name=organization"`
}

type CancelAQueuedDeployRequestResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Returns the deploy request whose deployment was canceled
	CancelAQueuedDeployRequest200ApplicationJSONObject map[string]map[string]interface{}
}
