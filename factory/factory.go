package factory

import (
	ppb "github.com/meateam/permission-service/proto"
)

// PermissionClientFactory is a factory for the Permission GRPC client
type PermissionClientFactory = func() ppb.PermissionClient
