package client

import (
	"crypto/tls"

	"github.com/apache/thrift/lib/go/thrift"
)

// ThriftClient ...
type ThriftClient struct {
	transport        thrift.TTransport
	ProtocolFactory  thrift.TProtocolFactory
	TransportFactory thrift.TTransportFactory
	client           thrift.TClient
	Address          string
}

// NewThriftClient ...
func NewThriftClient(addr string, transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, secure bool) (*ThriftClient, error) {
	var transport thrift.TTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		return nil, err
	}
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return nil, err
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	tsClient := thrift.NewTStandardClient(iprot, oprot)
	return &ThriftClient{
		transport:        transport,
		ProtocolFactory:  protocolFactory,
		TransportFactory: transportFactory,
		Address:          addr,
		client:           tsClient,
	}, nil
}

// Open ...
func (t *ThriftClient) Open() error {
	return t.transport.Open()
}

// Close ...
func (t *ThriftClient) Close() error {
	return t.transport.Close()
}

// GetClient ...
func (t *ThriftClient) GetClient() thrift.TClient {
	return t.client
}

// GetTransport ...
func (t *ThriftClient) GetTransport() thrift.TTransport {
	return t.transport
}
