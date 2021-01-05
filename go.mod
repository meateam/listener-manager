module github.com/meateam/gateway-template

go 1.12

require (
	github.com/gin-contrib/cors v0.0.0-20190424000812-bd1331c62cae
	github.com/gin-gonic/gin v1.4.0
	github.com/meateam/api-gateway v0.0.0-20191216140045-ecbd5d0ef5e1
	github.com/meateam/elasticsearch-logger v1.1.3-0.20190901111807-4e8b84fb9fda
	github.com/meateam/grpc-go-conn-pool v0.0.0-20201221202625-350108d14ffa
	github.com/meateam/permission-service v0.0.0-20191029101002-980dd2c31d08
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.6.0
	go.elastic.co/apm v1.6.0
	go.elastic.co/apm/module/apmgin v1.5.0
	go.elastic.co/apm/module/apmgrpc v1.6.0
	go.elastic.co/apm/module/apmhttp v1.6.0
	google.golang.org/grpc v1.34.0
	google.golang.org/grpc/examples v0.0.0-20201226181154-53788aa5dcb4 // indirect
)

replace github.com/meateam/gateway-middleware/logger => ./logger

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
