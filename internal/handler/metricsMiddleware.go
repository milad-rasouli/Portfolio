package handler

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type MetricsMiddleware struct {
	Logger            *zap.Logger
	errCounterAboutMe prometheus.Counter
	errCounterBlog    prometheus.Counter
	errCounterHome    prometheus.Counter
	errCounterContact prometheus.Counter
	errCounterUser    prometheus.Counter
	errCounterAdmin   prometheus.Counter
	reqCounterAboutMe prometheus.Counter
	reqCounterBlog    prometheus.Counter
	reqCounterHome    prometheus.Counter
	reqCounterContact prometheus.Counter
	reqCounterUser    prometheus.Counter
	reqCounterAdmin   prometheus.Counter
	reqLatencyAboutMe prometheus.Histogram
	reqLatencyBlog    prometheus.Histogram
	reqLatencyHome    prometheus.Histogram
	reqLatencyContact prometheus.Histogram
	reqLatencyUser    prometheus.Histogram
	reqLatencyAdmin   prometheus.Histogram
}

func NewMetricsMiddleware(logger *zap.Logger) *MetricsMiddleware {

	return &MetricsMiddleware{
		Logger: logger,
		errCounterAboutMe: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "about_me_error_counter",
			Name:      "about_me",
		}),
		errCounterBlog: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "blog_error_counter",
			Name:      "blog",
		}),
		errCounterHome: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "home_error_counter",
			Name:      "home",
		}),
		errCounterContact: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "contact_error_counter",
			Name:      "contact",
		}),
		errCounterUser: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "user_error_counter",
			Name:      "user",
		}),
		errCounterAdmin: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "admin_error_counter",
			Name:      "admin",
		}),
		reqCounterAboutMe: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "about_me_request_counter",
			Name:      "about_me",
		}),
		reqCounterBlog: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "blog_request_counter",
			Name:      "blog",
		}),
		reqCounterHome: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "home_request_counter",
			Name:      "home",
		}),
		reqCounterContact: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "contact_request_counter",
			Name:      "home",
		}),
		reqCounterUser: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "user_request_counter",
			Name:      "user",
		}),
		reqCounterAdmin: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "admin_request_counter",
			Name:      "admin",
		}),
		reqLatencyAboutMe: promauto.NewHistogram(prometheus.HistogramOpts{
			Namespace: "about_me_request_latency",
			Name:      "about_me",
			Buckets:   []float64{0.1, 0.5, 1},
		}),
		reqLatencyBlog: promauto.NewHistogram(prometheus.HistogramOpts{
			Namespace: "blog_request_latency",
			Name:      "blog",
			Buckets:   []float64{0.1, 0.5, 1},
		}),
		reqLatencyHome: promauto.NewHistogram(prometheus.HistogramOpts{
			Namespace: "home_request_latency",
			Name:      "home",
			Buckets:   []float64{0.1, 0.5, 1},
		}),
		reqLatencyContact: promauto.NewHistogram(prometheus.HistogramOpts{
			Namespace: "contact_request_latency",
			Name:      "contact",
			Buckets:   []float64{0.1, 0.5, 1},
		}),
		reqLatencyUser: promauto.NewHistogram(prometheus.HistogramOpts{
			Namespace: "user_request_latency",
			Name:      "user",
			Buckets:   []float64{0.1, 0.5, 1},
		}),
		reqLatencyAdmin: promauto.NewHistogram(prometheus.HistogramOpts{
			Namespace: "admin_request_latency",
			Name:      "admin",
			Buckets:   []float64{0.1, 0.5, 1},
		}),
	}
}

func (m *MetricsMiddleware) GetMetrics(c fiber.Ctx) error {
	handler := adaptor.HTTPHandler(promhttp.Handler())
	err := handler(c)
	if err != nil {
		m.Logger.Error("Read metrics error", zap.Error(err))
	}
	return err
}

func (m *MetricsMiddleware) Middleware(c fiber.Ctx) error {

	start := time.Now()
	err := c.Next()
	duration := time.Since(start).Seconds()

	path := c.Path()
	switch {
	case strings.HasPrefix(path, "/about-me/") || path == "/about-me":
		m.reqCounterAboutMe.Inc()
		m.reqLatencyAboutMe.Observe(duration)
		if err != nil {
			m.errCounterAboutMe.Inc()
		}
		m.Logger.Info("metricsMiddleware: about-me")

	case strings.HasPrefix(path, "/blog/") || path == "/blog":
		m.reqCounterBlog.Inc()
		m.reqLatencyBlog.Observe(duration)
		if err != nil {
			m.errCounterBlog.Inc()
		}
		m.Logger.Info("metricsMiddleware: blog")

	case path == "/":
		m.reqCounterHome.Inc()
		m.reqLatencyHome.Observe(duration)
		if err != nil {
			m.errCounterHome.Inc()
		}
		m.Logger.Info("metricsMiddleware: home") //TODO: remove this line

	case strings.HasPrefix(path, "/contact/") || path == "/contact":
		m.reqCounterContact.Inc()
		m.reqLatencyContact.Observe(duration)
		if err != nil {
			m.errCounterContact.Inc()
		}
		m.Logger.Info("metricsMiddleware: contact")

	case strings.HasPrefix(path, "/user/") || path == "/user":
		m.reqCounterUser.Inc()
		m.reqLatencyUser.Observe(duration)
		if err != nil {
			m.errCounterUser.Inc()
		}
		m.Logger.Info("metricsMiddleware: user")

	case strings.HasPrefix(path, "/admin/") || path == "/admin":
		m.reqCounterAdmin.Inc()
		m.reqLatencyAdmin.Observe(duration)
		if err != nil {
			m.errCounterAdmin.Inc()
		}
	}
	return err

}
func (m *MetricsMiddleware) Register(metrics fiber.Router) {
	metrics.Get("/", m.GetMetrics) //TODO: use go http library and diffrent port from app
}
