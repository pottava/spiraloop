package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pottava/spiraloop/api/generated/restapi/operations"
)

func postStart(params operations.PostStartParams) middleware.Responder {
	return operations.NewPostStartCreated()
}
