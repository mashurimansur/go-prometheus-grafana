package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "my_app_request_total",
			Help: "Total number of processed requests",
		},
		[]string{"path", "status"},
	)

	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "my_request_error_total",
			Help: "Total number of error requests",
		},
		[]string{"path", "status"},
	)
)

func PrometheusInit() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(ErrorCount)
}

func TrackMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		c.Next()

		status := c.Writer.Status()
		RequestCount.WithLabelValues(path, fmt.Sprint(status)).Inc()
		if status >= 400 {
			ErrorCount.WithLabelValues(path, fmt.Sprint(status)).Inc()
		}
	}
}
