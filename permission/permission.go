package permission

import (
	"fmt"

	"github.com/gin-gonic/gin"
	grpcPoolTypes "github.com/meateam/grpc-go-conn-pool/grpc/types"
	"github.com/meateam/listener-manager/factory"
	ppb "github.com/meateam/permission-service/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const ()

// Manager is a structure that handles permission requests.
type Manager struct {
	// PermissionClientFactory
	permissionClient factory.PermissionClientFactory

	logger *logrus.Logger
}

// NewManager creates a new Manager, and initializes clients with the given connection.
// Logger defaults to logrus.New().
func NewManager(
	permissionConn *grpcPoolTypes.ConnPool,
	logger *logrus.Logger,
) *Manager {
	// If no logger is given, use a default logger.
	if logger == nil {
		logger = logrus.New()
	}

	m := &Manager{logger: logger}

	m.permissionClient = func() ppb.PermissionClient {
		return ppb.NewPermissionClient((*permissionConn).Conn())
	}

	return m
}

// GetFilePermissions is a route function for retrieving permissions of a file
// File id is extracted from url params
func (m *Manager) GetFilePermissions(c *gin.Context) {

	permissionsRequest := &ppb.GetFilePermissionsRequest{FileID: "fakefileID"}
	permissionsResponse, err := m.permissionClient().GetFilePermissions(c, permissionsRequest)
	if err != nil && status.Code(err) != codes.NotFound {
		err = fmt.Errorf("error while GetFilePermissions: %v", err)
		return
	}

	fmt.Printf("%v", permissionsResponse)
}
