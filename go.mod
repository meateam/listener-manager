module github.com/meateam/listener-manager

go 1.12

require (
	github.com/containerd/containerd v1.4.3 // indirect
	github.com/containerd/continuity v0.0.0-20201208142359-180525291bb7 // indirect
	github.com/docker/docker v20.10.2+incompatible // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gin-contrib/cors v0.0.0-20190424000812-bd1331c62cae
	github.com/gin-gonic/gin v1.4.0
	github.com/meateam/elasticsearch-logger v1.1.3-0.20190901111807-4e8b84fb9fda
	github.com/meateam/gateway-template v0.0.0-20191222161709-1df00476e2d6
	github.com/meateam/grpc-go-conn-pool v0.0.0-20201221202625-350108d14ffa
	github.com/meateam/permission-service v0.0.0-20191029101002-980dd2c31d08
	github.com/moby/sys/mount v0.2.0 // indirect
	github.com/moby/sys/symlink v0.1.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v0.1.1 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.6.0
	github.com/streadway/amqp v1.0.0
	go.elastic.co/apm v1.6.0
	go.elastic.co/apm/module/apmgin v1.5.0
	go.elastic.co/apm/module/apmgrpc v1.6.0
	go.elastic.co/apm/module/apmhttp v1.6.0
	google.golang.org/grpc v1.34.0
	google.golang.org/grpc/examples v0.0.0-20201226181154-53788aa5dcb4 // indirect
)

replace github.com/meateam/gateway-middleware/logger => ./logger

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
