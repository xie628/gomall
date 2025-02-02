package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/xie628/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/xie628/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/xie628/gomall/demo/demo_proto/middleware"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}

	c, err := echoservice.NewClient("demo_proto",
		client.WithResolver(r),
		client.WithShortConnection(), // 使用短链接
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithMiddleware(middleware.Middleware),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := metainfo.WithPersistentValue(context.Background(),
		"CLIENT_NAME",
		"demo_proto_client",
	)

	res, err := c.Echo(ctx, &pbapi.Request{Message: "hh"})
	var bizErr *kerrors.GRPCBizStatusError
	if err != nil {
		ok := errors.As(err, &bizErr)
		if ok {
			fmt.Printf("%#v", bizErr)
		}
		klog.Fatal(err)
	}
	fmt.Printf("%v", res)
}
