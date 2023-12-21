package server

import (
	"go_crud/common/constants"
	con "go_crud/internal/server/controllers"
	middlewares "go_crud/internal/server/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var METHODS = constants.HttpMethodsE

func SetupRoutes(server *gin.Engine) {
	// TODO: clean this up, use router groups
	routes := Routes{
		Route{
			Method:      METHODS.GET,
			Path:        "/hc",
			Handler:     con.HealthCheck,
			Middlewares: nil,
		},
		Route{
			Method:      METHODS.GET,
			Path:        "/actuator/prometheus",
			Handler:     gin.WrapH(promhttp.Handler()),
			Middlewares: nil,
		},
		Route{
			Method:      METHODS.POST,
			Path:        "/auth/login",
			Handler:     con.AuthLoginHandler,
			Middlewares: []gin.HandlerFunc{middlewares.TrackLatencyMiddleware()},
		},
		Route{
			Method:      METHODS.POST,
			Path:        "/auth/logout",
			Handler:     con.AuthLogoutHandler,
			Middlewares: []gin.HandlerFunc{middlewares.TrackLatencyMiddleware()},
		},
		Route{
			Method:      METHODS.GET,
			Path:        "/user/:username",
			Handler:     con.GetUserHandler,
			Middlewares: []gin.HandlerFunc{middlewares.TrackLatencyMiddleware(), middlewares.ValidateUserMiddleware},
		},
		Route{
			Method:      METHODS.GET,
			Path:        "/me",
			Handler:     con.GetSelfUserHandler,
			Middlewares: []gin.HandlerFunc{middlewares.TrackLatencyMiddleware(), middlewares.ValidateUserMiddleware},
		},
		Route{
			Method:      METHODS.DELETE,
			Path:        "/user/:username",
			Handler:     con.DeleteUserHandler,
			Middlewares: []gin.HandlerFunc{middlewares.TrackLatencyMiddleware(), middlewares.ValidateUserMiddleware},
		},
		Route{
			Method:      METHODS.POST,
			Path:        "/user",
			Handler:     con.CreateUser,
			Middlewares: nil,
		},
		Route{
			Method:      METHODS.GET,
			Path:        "/post/:slug",
			Handler:     con.GetPostBySlugHandler,
			Middlewares: []gin.HandlerFunc{middlewares.TrackLatencyMiddleware(), middlewares.ValidateUserMiddleware},
		},
		Route{
			Method:      METHODS.POST,
			Path:        "/post/create",
			Handler:     con.CreatePostHandler,
			Middlewares: []gin.HandlerFunc{middlewares.TrackLatencyMiddleware(), middlewares.ValidateUserMiddleware},
		},
		Route{
			Method:      METHODS.GET,
			Path:        "/post/fetchBulk",
			Handler:     con.FetchPostsBulk,
			Middlewares: []gin.HandlerFunc{middlewares.TrackLatencyMiddleware(), middlewares.ValidateUserMiddleware},
		},
	}

	for _, route := range routes {
		allHandlers := make([]gin.HandlerFunc, 0)
		if route.Middlewares != nil {
			allHandlers = append(allHandlers, route.Middlewares...)
		}

		allHandlers = append(allHandlers, route.Handler)
		server.Handle(route.Method, route.Path, allHandlers...)
	}
}
