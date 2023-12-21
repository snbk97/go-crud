package server

import (
	storage "go_crud/common/storage"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	g := gin.Default()

	setupSentry(g)
	setupCors(g)
	SetupRoutes(g)

	setupAdditionals()
	return g
}

func setupCors(server *gin.Engine) {
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func setupSentry(server *gin.Engine) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	server.Use(sentrygin.New(sentrygin.Options{Repanic: true}))
}

func setupAdditionals() {
	storage.InitCache()      // init cache
	storage.InitDb()         // init db connection
	storage.RegisterModels() // init and migrate models
	storage.ConnectModels()  // connect models to db, so that model can use db
}
