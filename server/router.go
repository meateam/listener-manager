package server

import (
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	loggermiddleware "github.com/meateam/api-gateway/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.elastic.co/apm/module/apmgin"
	"go.elastic.co/apm/module/apmhttp"
	"google.golang.org/grpc"
)

const (
	healthcheckRouter = "/api/healtcheck"
	uploadRouteRegexp = "/api/upload.+"
)

// NewRouter creates new gin.Engine for the gateway-template server and sets it up.
func NewRouter(logger *logrus.Logger) (*gin.Engine, []*grpc.ClientConn) {
	// If no logger is given, use a default logger.
	if logger == nil {
		logger = logrus.New()
	}

	gin.DisableConsoleColor()
	r := gin.New()

	// Setup logging, metrics, cors middlewares.
	r.Use(
		// Ignore logging healthcheck routes.
		gin.LoggerWithWriter(gin.DefaultWriter, healthcheckRouter),
		gin.Recovery(),
		apmgin.Middleware(r),
		cors.New(corsRouterConfig()),
		// Elasticsearch logger middleware.
		loggermiddleware.SetLogger(
			&loggermiddleware.Config{
				Logger:             logger,
				SkipPath:           []string{healthcheckRouter},
				SkipBodyPathRegexp: regexp.MustCompile(uploadRouteRegexp),
			},
		),
	)

	apiRoutesGroup := r.Group("/api")

	// Health Check route.
	apiRoutesGroup.GET("/healthcheck", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Frontend configuration route.
	apiRoutesGroup.GET("/config", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"chromeDownloadLink": viper.GetString(configDownloadChromeURL),
				"apmServerUrl":       viper.GetString(configExternalApmURL),
				"environment":        os.Getenv("ELASTIC_APM_ENVIRONMENT"),
				"authUrl":            viper.GetString(configAuthURL),
				"supportLink":        viper.GetString(configSupportLink),
			},
		)
	})

	// Initiate services gRPC connections.
	/*
		conn1, err := initServiceConn(viper.GetString(config1Service))
		if err != nil {
			logger.Fatalf("couldn't setup service1 connection: %v", err)
		}

		// Initiate routers.
		r1 := router1.NewRouter(conn1, logger)

		authRequiredRoutesGroup := apiRoutesGroup.Group("/path/", middlewareFunction)

		// Initiate client connection to file service.
		r1.Setup(authRequiredRoutesGroup)
	*/

	// Create a slice to manage connections and return it: []*grpc.ClientConn{conn1}
	return r, []*grpc.ClientConn{}
}

// corsRouterConfig configures cors policy for cors.New gin middleware.
func corsRouterConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AddExposeHeaders("x-uploadid")
	corsConfig.AllowAllOrigins = false
	corsConfig.AllowWildcard = true
	corsConfig.AllowOrigins = strings.Split(viper.GetString(configAllowOrigins), ",")
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders(
		"x-content-length",
		"authorization",
		"cache-control",
		"x-requested-with",
		"content-disposition",
		"content-range",
		apmhttp.TraceparentHeader,
	)

	return corsConfig
}
