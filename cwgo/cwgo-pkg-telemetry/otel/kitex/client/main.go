/*
 * Copyright 2024 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/cloudwego-contrib/cwgo-pkg/telemetry/instrumentation/otelkitex"
	"github.com/cloudwego-contrib/cwgo-pkg/telemetry/provider/otelprovider"
	"github.com/cloudwego/kitex-examples/hello/kitex_gen/api/hello"
	"github.com/cloudwego/kitex-examples/hello/kitex_gen/api"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.opentelemetry.io/otel"
)

func main() {
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelDebug)

	serviceName := "echo-client"

	p := otelprovider.NewOpenTelemetryProvider(
		otelprovider.WithServiceName(serviceName),
		// Support setting ExportEndpoint via environment variables: OTEL_EXPORTER_OTLP_ENDPOINT
		otelprovider.WithExportEndpoint(":4317"),
		otelprovider.WithInsecure(),
		otelprovider.WithRPCServer(),
	)
	defer p.Shutdown(context.Background())

	demoServerAddr, ok := os.LookupEnv("DEMO_SERVER_ENDPOINT")
	if !ok {
		demoServerAddr = "0.0.0.0:8181"
	}

	c, err := hello.NewClient(
		"echo",
		client.WithHostPorts(demoServerAddr),
		client.WithSuite(otelkitex.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
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

	klog.CtxInfof(ctx, "req:%v, res:%v", req, resp)
}
