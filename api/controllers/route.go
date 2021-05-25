// Package controllers defines application's routes.
package controllers

import (
	"github.com/pottava/spiraloop/api/generated/restapi/operations"
)

// Routes set API handlers
func Routes(api *operations.SpiraloopAPI) {
	api.PostStartHandler = operations.PostStartHandlerFunc(postStart)
	api.PostEndHandler = operations.PostEndHandlerFunc(postEnd)
}
