package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego-contrib/cwgo-pkg/registry/zookeeper/zookeeperkitex/resolver"
	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
)

func main() {
	r, err := resolver.NewZookeeperResolver([]string{"127.0.0.1:2181"}, 40*time.Second)
	if err != nil {
		panic(err)
	}

	client, err := echo.NewClient("Hello", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		resp, err := client.Echo(ctx, &api.Request{Message: "Hello"})
		cancel()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
