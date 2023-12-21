package server

import "github.com/gin-gonic/gin"

type Route struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

type Routes []Route
