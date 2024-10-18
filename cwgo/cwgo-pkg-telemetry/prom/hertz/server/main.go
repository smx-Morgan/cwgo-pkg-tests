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
	"fmt"
	"net/http"

	"github.com/cloudwego-contrib/cwgo-pkg/telemetry/instrumentation/otelhertz"
	"github.com/cloudwego-contrib/cwgo-pkg/telemetry/provider/promprovider"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/prometheus/client_golang/prometheus"

	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
)

func main() {
	hlog.SetLogger(hertzlogrus.NewLogger())
	hlog.SetLevel(hlog.LevelDebug)

	serviceName := "demo_hertz_server"

	registry := prometheus.NewRegistry()
	provider := promprovider.NewPromProvider(
		promprovider.WithServiceName(serviceName),
		promprovider.WithRegistry(registry),
		promprovider.WithHttpServer(),
	)
	provider.Serve(":9090", "/metrics")

	tracer := otelhertz.NewServerTracer()
	h := server.Default(server.WithTracer(tracer), server.WithHostPorts(":39888"))
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		hlog.CtxDebugf(c, "message received successfully")

		ctx.JSON(consts.StatusOK, "pong")
	})

	promServerResp, err := http.Get("http://localhost:9090/metrics-demo")
	if err != nil {
		return
	}
	if promServerResp.StatusCode == http.StatusOK {
		fmt.Print("status is 200\n")
	}

	h.Spin()
}
