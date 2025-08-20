package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "post_rest_app_requests_total",
			Help: "Total number of requests processed by the Post Rest web server.",
		},
		[]string{"path", "status"},
	)

	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "post_rest_app_requests_errors_total",
			Help: "Total number of error requests processed by the Post Rest web server.",
		},
		[]string{"path", "status"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "post_rest_app_requests_duration",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)
)

func PrometheusInit() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(ErrorCount)
	prometheus.MustRegister(RequestDuration)
}

func TrackMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		timer := prometheus.NewTimer(RequestDuration.WithLabelValues(path, method))
		c.Next()
		status := c.Writer.Status()
		RequestCount.WithLabelValues(path, http.StatusText(status)).Inc()
		if status >= 400 {
			ErrorCount.WithLabelValues(path, http.StatusText(status)).Inc()
		}
		timer.ObserveDuration()
	}
}
