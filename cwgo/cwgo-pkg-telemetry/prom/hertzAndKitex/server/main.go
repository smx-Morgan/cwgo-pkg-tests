package main

import (
	"context"
	"cwgo-test/cwgo-pkg-telemetry/thrift/kitex_gen/api"
	"cwgo-test/cwgo-pkg-telemetry/thrift/kitex_gen/api/hello"
	"net"

	"github.com/cloudwego-contrib/cwgo-pkg/telemetry/instrumentation/otelhertz"
	"github.com/cloudwego-contrib/cwgo-pkg/telemetry/instrumentation/otelkitex"
	"github.com/cloudwego-contrib/cwgo-pkg/telemetry/provider/promprovider"
	"github.com/cloudwego/hertz/pkg/app"
	hertzserver "github.com/cloudwego/hertz/pkg/app/server"
	kitexserver "github.com/cloudwego/kitex/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func main() {
	serviceName := "demo_hertz_and_kitex_server"

	provider := promprovider.NewPromProvider(
		promprovider.WithServiceName(serviceName),
		promprovider.WithHttpServer(),
		promprovider.WithRPCServer(),
	)

	provider.Serve(":9090", "/metrics")

	go startRPCServer()
	startHttpServer()
}

func startHttpServer() {
	tracer := otelhertz.NewServerTracer()
	h := hertzserver.Default(hertzserver.WithTracer(tracer), hertzserver.WithHostPorts(":39888"))
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		hlog.CtxDebugf(c, "message received successfully")
		ctx.JSON(consts.StatusOK, "pong")
	})

	h.Spin()
}

func startRPCServer() {
	serviceName := "demo_hertz_and_kitex_server"
	addr, err := net.ResolveTCPAddr("tcp", ":8181")
	if err != nil {
		panic(err)
	}

	svr := hello.NewServer(
		new(EchoImpl),
		kitexserver.WithServiceAddr(addr),
		kitexserver.WithSuite(otelkitex.NewServerSuite()),
		kitexserver.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	if err := svr.Run(); err != nil {
		klog.Fatalf("server stopped with error:", err)
	}
}

var _ api.Hello = &EchoImpl{}

type EchoImpl struct{}

// Echo implements the Echo interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	klog.CtxDebugf(ctx, "echo called: %s", req.GetMessage())
	return &api.Response{Message: req.Message}, nil
}
