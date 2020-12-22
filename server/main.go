package main

import (
	"fmt"

	"github.com/duyledat197/test-thrift/base/srv"
	"github.com/duyledat197/test-thrift/server/handler"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/duyledat197/test-thrift/thr/tutorial"
)

func main() {
	addr := "127.0.0.1:12345"
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	hdl := handler.NewCaculatorHandler()
	processor := tutorial.NewCalculatorProcessor(hdl)
	thrs, err := srv.NewThriftServer(addr, transportFactory, protocolFactory, processor, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the simple server... on ", addr)
	thrs.Start()
}
