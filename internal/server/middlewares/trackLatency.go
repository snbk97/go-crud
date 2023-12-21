package server

import (
	"fmt"
	mx "go_crud/internal/server/metrics"
	"time"

	"github.com/gin-gonic/gin"
)

func TrackLatencyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)
		// TODO: log more generic paths, or make the paths more generic using query params, instead of path params
		mx.InboundMetricsHistogram.WithLabelValues("go_crud", c.Request.Method, c.Request.URL.Path, fmt.Sprint(c.Writer.Status())).Observe(duration.Seconds())

	}
}
