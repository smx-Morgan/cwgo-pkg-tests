package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/cloudwego-contrib/cwgo-pkg/registry/consul/consulhertz"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	consulapi "github.com/hashicorp/consul/api"
)

var (
	wg      sync.WaitGroup
	localIP = "127.0.0.1"
)

func main() {
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	consulClient, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
		return
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		addr := net.JoinHostPort(localIP, "8888")
		r := consulhertz.NewConsulRegister(consulClient)
		h := server.Default(
			server.WithHostPorts(addr),
			server.WithRegistry(r, &registry.Info{
				ServiceName: "hertz.test.demo",
				Addr:        utils.NewNetAddr("tcp", addr),
				Weight:      10,
				Tags:        nil,
			}),
		)

		h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
			ctx.JSON(consts.StatusOK, utils.H{"ping": "pong1"})
		})
		h.Spin()
	}()
	go func() {
		defer wg.Done()
		addr := net.JoinHostPort(localIP, "8889")
		r := consulhertz.NewConsulRegister(consulClient)
		h := server.Default(
			server.WithHostPorts(addr),
			server.WithRegistry(r, &registry.Info{
				ServiceName: "hertz.test.demo",
				Addr:        utils.NewNetAddr("tcp", addr),
				Weight:      10,
				Tags:        nil,
			}),
		)

		h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
			ctx.JSON(consts.StatusOK, utils.H{"ping": "pong2"})
		})
		h.Spin()
	}()

	wg.Wait()
}
