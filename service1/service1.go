package service1

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Client is a mock interface which is supposed to be imported from the service1 pb.
// The client contains the methods implemented by the service.
type Client interface {
	Method1(ctx context.Context)
}

// Router is a structure that handles service1 related requests.
type Router struct {
	service1Client Client
	logger         *logrus.Logger
}

// NewRouter creates a new Router, and initializes clients of the service1 service
// with the given connection. If logger is non-nil then it will
// be set as-is, otherwise logger would default to logrus.New().
func NewRouter(
	service1Conn *grpc.ClientConn,
	logger *logrus.Logger,
) *Router {
	// If no logger is given, use a default logger.
	if logger == nil {
		logger = logrus.New()
	}

	r := &Router{logger: logger}

	// r.service1Client = qpb.NewQuotaServiceClient(quotaConn)

	return r
}
