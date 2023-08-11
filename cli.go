package thriftutil

import (
	"time"

	"github.com/apache/thrift/lib/go/thrift"
)

type Connection struct {
	client *thrift.TStandardClient
	socket thrift.TTransport
}

func (c *Connection) Client() *thrift.TStandardClient {
	return c.client
}

func (c *Connection) Close() error {
	return c.socket.Close()
}

func ConnectServer(addr string, connTimeout time.Duration, dataTimeout time.Duration) (*Connection, error) {
	addr = ThriftAddress(addr)
	var transport thrift.TTransport
	var err error
	clientConfig := &thrift.TConfiguration{
		ConnectTimeout: connTimeout,
		SocketTimeout:  dataTimeout,
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(clientConfig)
	transportFactory := thrift.NewTFramedTransportFactoryConf(thrift.NewTBufferedTransportFactory(8192), clientConfig)
	transport = thrift.NewTSocketConf(addr, clientConfig)
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return nil, err
	}
	if err := transport.Open(); err != nil {
		return nil, err
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	return &Connection{
		client: thrift.NewTStandardClient(iprot, oprot),
		socket: transport,
	}, nil
}
