// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego-contrib/cwgo-pkg/config/apollo/apollo"
	cwclient "github.com/cloudwego-contrib/cwgo-pkg/config/apollo/client"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/cloudwego-contrib/cwgo-pkg/config/apollo/utils"
)

// Customed by user
type configLog struct{}

func (cl *configLog) Apply(opt *utils.Options) {
	fn := func(cp *apollo.ConfigParam) {
		klog.Infof("apollo config %v", cp)
	}
	opt.ApolloCustomFunctions = append(opt.ApolloCustomFunctions, fn)
}

func main() {
	klog.SetLevel(klog.LevelDebug)
	apolloClient, err := apollo.NewClient(apollo.Options{})
	if err != nil {
		panic(err)
	}
	cl := &configLog{}

	serviceName := "ServiceName" // your server-side service name
	clientName := "ClientName"   // your client-side service name
	client, err := echo.NewClient(
		serviceName,
		client.WithHostPorts("localhost:8899"),
		client.WithSuite(cwclient.NewSuite(serviceName, clientName, apolloClient, cl)),
	)
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &api.Request{Message: "my request"}
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			klog.Errorf("take request cwerror: %v", err)
		} else {
			klog.Infof("receive response %v", resp)
		}
		time.Sleep(time.Second * 10)
	}
}
