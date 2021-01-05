package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/meateam/listener-manager/permission"

	grpcPool "github.com/meateam/grpc-go-conn-pool/grpc"
	grpcPoolOptions "github.com/meateam/grpc-go-conn-pool/grpc/options"
	grpcPoolTypes "github.com/meateam/grpc-go-conn-pool/grpc/types"
	loggermiddleware "github.com/meateam/listener-manager/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.elastic.co/apm/module/apmgin"
	"go.elastic.co/apm/module/apmgrpc"
	"go.elastic.co/apm/module/apmhttp"
	"google.golang.org/grpc"
)

const (
	healthcheckRouter = "/api/healtcheck"
)

// NewManager creates new gin.Engine for the listener-manager server and sets it up.
func NewManager(logger *logrus.Logger) (*gin.Engine, []*grpc.ClientConn) {
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
				Logger:   logger,
				SkipPath: []string{healthcheckRouter},
			},
		),
	)

	apiRoutesGroup := r.Group("/api")

	// Health Check route.
	apiRoutesGroup.GET("/healthcheck", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Initiate services gRPC connections.
	permissionConn, err := initServiceConn(viper.GetString(configPermissionService))
	if err != nil {
		logger.Fatalf("couldn't setup service1 connection: %v", err)
	}

	// Initiate routers.
	pm := permission.NewManager(permissionConn, logger)

	// Authentication middleware on routes group.
	authRequiredRoutesGroup := apiRoutesGroup.Group("/")

	// Initiate client connection to file service.
	pm.Setup(authRequiredRoutesGroup)

	// Create a slice to manage connections and return it: []*grpc.ClientConn{conn1}
	return r, []*grpc.ClientConn{}
}

// initServiceConn creates a gRPC connection pool to url, returns the created connection pool
// and nil err on success. Returns non-nil error if any error occurred while
// creating the connection pool.
func initServiceConn(url string) (*grpcPoolTypes.ConnPool, error) {
	ctx := context.Background()
	connPool, err := grpcPool.DialPool(ctx,
		grpcPoolOptions.WithGRPCDialOption(grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor())),
		grpcPoolOptions.WithGRPCDialOption(grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(10<<20))),
		grpcPoolOptions.WithGRPCDialOption(grpc.WithInsecure()),
		grpcPoolOptions.WithEndpoint(url),
		grpcPoolOptions.WithGRPCConnectionPool(viper.GetInt(configPoolSize)),
	)
	if err != nil {
		return nil, err
	}
	return &connPool, nil
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

// reviveConns is the function for reviving the connections in the badConns connection pool
func reviveConns(badConns <-chan *grpcPoolTypes.ConnPool) {
	for {
		// Pull the pointer to the pool from the channel.
		// Will run when the channel isn't empty
		pool := <-badConns
		go func(pool *grpcPoolTypes.ConnPool) {
			// Get the target url
			target := (*pool).Conn().Target()
			var newPool *grpcPoolTypes.ConnPool
			err := fmt.Errorf("temp")
			for err != nil {
				time.Sleep(time.Second * time.Duration(2))
				err = nil
				newPool, err = initServiceConn(target)
			}
			(*pool).Close()
			// Replace the pointer to the new pool
			*pool = *newPool
		}(pool)
	}
}
