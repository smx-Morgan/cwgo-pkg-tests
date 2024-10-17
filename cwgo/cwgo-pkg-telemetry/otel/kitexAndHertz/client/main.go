package main

import (
	"context"
	"cwgo-test/cwgo-pkg-telemetry/thrift/kitex_gen/api"
	"cwgo-test/cwgo-pkg-telemetry/thrift/kitex_gen/api/hello"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/cloudwego-contrib/cwgo-pkg/telemetry/instrumentation/otelhertz"
	"github.com/cloudwego-contrib/cwgo-pkg/telemetry/instrumentation/otelkitex"
	"github.com/cloudwego-contrib/cwgo-pkg/telemetry/provider/otelprovider"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzclient "github.com/cloudwego/hertz/pkg/app/client"
	kitexclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"

	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
)

func main() {
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelDebug)

	hlog.SetLogger(hertzlogrus.NewLogger())
	hlog.SetLevel(hlog.LevelDebug)

	serviceName := "echo-client"

	p := otelprovider.NewOpenTelemetryProvider(
		otelprovider.WithServiceName(serviceName),
		otelprovider.WithHttpServer(),
		otelprovider.WithRPCServer(),
		otelprovider.WithInsecure(),
		otelprovider.WithExportEndpoint(":4317"),
	)

	defer p.Shutdown(context.Background())

	go pingHttpServer()
	go pingRPCServer()

	select {}
}

func pingHttpServer() {
	c, _ := hertzclient.NewClient()
	c.Use(otelhertz.ClientMiddleware())

	for {
		ctx, span := otel.Tracer("github.com/hertz-contrib/obs-opentelemetry").
			Start(context.Background(), "loop")

		_, b, err := c.Get(ctx, nil, "http://0.0.0.0:8888/ping?foo=bar")
		if err != nil {
			hlog.CtxErrorf(ctx, err.Error())
		}

		span.SetAttributes(attribute.String("msg", string(b)))

		hlog.CtxInfof(ctx, "hertz client %s", string(b))
		span.End()

		<-time.After(time.Second)
	}
}

func pingRPCServer() {
	serviceName := "echo-client"

	demoServerAddr, ok := os.LookupEnv("DEMO_SERVER_ENDPOINT")
	if !ok {
		demoServerAddr = "0.0.0.0:8181"
	}

	c, err := hello.NewClient(
		"echo",
		kitexclient.WithHostPorts(demoServerAddr),
		kitexclient.WithSuite(otelkitex.NewClientSuite()),
		kitexclient.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	if err != nil {
		klog.Fatal(err)
	}

	// Yields a constantly-changing number
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		call(c)
		<-time.After(time.Second)
	}
}

func call(c hello.Client) {
	ctx, span := otel.Tracer("client").Start(context.Background(), "root")
	defer span.End()

	randomInt := rand.Intn(1000)
	req := &api.Request{Message: "my request " + strconv.Itoa(randomInt)}

	resp, err := c.Echo(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "err %v", err)
	}

	klog.CtxInfof(ctx, "kitex client: req:%v, res:%v", req, resp)
}
