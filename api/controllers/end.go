package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pottava/spiraloop/api/generated/restapi/operations"
)

func postEnd(params operations.PostEndParams) middleware.Responder {
	return operations.NewPostEndCreated()
}
