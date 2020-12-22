package jaeger

import (
	"github.com/apache/thrift/lib/go/thrift"

	"github.com/jaegertracing/jaeger/cmd/agent/app/servers/thriftudp"
	"github.com/jaegertracing/jaeger/thrift-gen/agent"
)

// NewJaegerThriftUDPClient creates a new jaeger agent client that works like Jaeger client
func NewJaegerThriftUDPClient(hostPort string, protocolFactory thrift.TProtocolFactory) (thrift.TClient, error) {
	clientTransport, err := thriftudp.NewTUDPClientTransport(hostPort, "")
	if err != nil {
		return nil, err
	}

	client := agent.NewAgentClientFactory(clientTransport, protocolFactory)
	return client.Client_(), nil
}
