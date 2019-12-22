package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// NewRouter creates new gin.Engine for the gateway-template server and sets it up.
func NewRouter(logger *logrus.Logger) (*gin.Engine, []*grpc.ClientConn) {
	return nil, nil
}
