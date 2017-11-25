package router

import (
	"github.com/FlorentinDUBOIS/bananas/service"
	"github.com/gin-gonic/gin"
)

// Route structure to describe a http route
type Route struct {
	Method  string
	Path    string
	Handler func(ctx *gin.Context)
}

// Routes is a slice of Route
var Routes = []Route{
	{
		Method:  "POST",
		Path:    "/messages",
		Handler: service.AnswerMessage,
	},
}
