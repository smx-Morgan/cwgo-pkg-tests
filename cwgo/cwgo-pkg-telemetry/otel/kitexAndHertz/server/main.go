package main

import (
	"context"
	"cwgo-test/cwgo-pkg-telemetry/thrift/kitex_gen/api"
	"cwgo-test/cwgo-pkg-telemetry/thrift/kitex_gen/api/hello"
	"net"

	"github.com/cloudwego-contrib/cwgo-pkg/telemetry/instrumentation/otelhertz"
	"github.com/cloudwego-contrib/cwgo-pkg/telemetry/instrumentation/otelkitex"
	"github.com/cloudwego-contrib/cwgo-pkg/telemetry/provider/otelprovider"
	"github.com/cloudwego/hertz/pkg/app"
	hertzserver "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	kitexserver "github.com/cloudwego/kitex/server"

	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
)

func main() {
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelDebug)

	hlog.SetLogger(hertzlogrus.NewLogger())
	hlog.SetLevel(hlog.LevelDebug)

	serviceName := "demo-hertz-server"
	p := otelprovider.NewOpenTelemetryProvider(
		otelprovider.WithServiceName(serviceName),
		otelprovider.WithExportEndpoint("localhost:4317"),
		otelprovider.WithRPCServer(),
		otelprovider.WithHttpServer(),
		otelprovider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	go runRPCServer()
	runHttpServer()
}

func runHttpServer() {
	tracer, cfg := otelhertz.NewServerOption()
	h := hertzserver.Default(tracer)
	h.Use(otelhertz.ServerMiddleware(cfg))

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		req := &api.Request{Message: "my request"}

		hlog.CtxDebugf(c, "message received successfully: %s", req.Message)
		ctx.JSON(consts.StatusOK, "resp")
	})
	h.Spin()
}

func runRPCServer() {
	serviceName := "echo"

	addr, err := net.ResolveTCPAddr("tcp", ":8181")
	if err != nil {
		panic(err)
	}
	svr := hello.NewServer(
		new(HelloImpl),
		kitexserver.WithServiceAddr(addr),
		kitexserver.WithSuite(otelkitex.NewServerSuite()),
		kitexserver.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	if err := svr.Run(); err != nil {
		klog.Fatalf("server stopped with error:", err)
	}
}

var _ api.Hello = &HelloImpl{}

type HelloImpl struct{}

// Echo implements the Echo interface.
func (s *HelloImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	klog.CtxDebugf(ctx, "echo called: %s", req.GetMessage())
	// nowSec := time.Now().Second()
	// if nowSec%3 == 1 {
	// 	klog.CtxErrorf(ctx, "mock error with request message: %s", req.GetMessage())
	// 	return nil, errors.New("mock error")
	// }
	// if nowSec%3 == 2 {
	// 	klog.CtxErrorf(ctx, "mock panic with request message: %s", req.GetMessage())
	// 	panic("mock panic")
	// }
	return &api.Response{Message: req.Message}, nil
}
