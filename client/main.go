package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/duyledat197/test-thrift/base/client"
	"github.com/duyledat197/test-thrift/thr/tutorial"
)

func main() {
	ctx := context.Background()
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	addr := "localhost:12345"

	tc, err := client.NewThriftClient(addr, transportFactory, protocolFactory, false)
	if err != nil {
		panic(err)
	}
	if err := tc.Open(); err != nil {
		panic(err)
	}
	defer tc.Close()
	client := tutorial.NewCalculatorClient(tc.GetClient())
	req := &tutorial.AddRequest{
		A: 1,
		B: 2,
	}
	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.GetResult_())
}
