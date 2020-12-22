package srv

import (
	"crypto/tls"

	"github.com/apache/thrift/lib/go/thrift"
)

// ThriftServer ...
type ThriftServer struct {
	Address          string
	TransportFactory thrift.TTransportFactory
	ProtocolFactory  thrift.TProtocolFactory
	Transport        thrift.TServerTransport
	Server           *thrift.TSimpleServer
	Processor        thrift.TProcessor
}

// NewThriftServer ...
func NewThriftServer(
	addr string,
	transportFactory thrift.TTransportFactory,
	protocolFactory thrift.TProtocolFactory,
	processor thrift.TProcessor,
	secure bool,
) (*ThriftServer, error) {
	var transport thrift.TServerTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		if cert, err := tls.LoadX509KeyPair("server.crt", "server.key"); err == nil {
			cfg.Certificates = append(cfg.Certificates, cert)
		} else {
			return nil, err
		}
		transport, err = thrift.NewTSSLServerSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTServerSocket(addr)
	}

	if err != nil {
		return nil, err
	}
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	return &ThriftServer{
		Address:          addr,
		TransportFactory: transportFactory,
		ProtocolFactory:  protocolFactory,
		Transport:        transport,
		Processor:        processor,
		Server:           server,
	}, nil
}

// Start ...
func (t *ThriftServer) Start() error {
	return t.Server.Serve()
}
