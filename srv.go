package thriftutil

import "github.com/apache/thrift/lib/go/thrift"

func StartServer(addr string, processor thrift.TProcessor) (server *thrift.TSimpleServer, err error) {
	addr = ThriftAddress(addr)
	serverConfig := &thrift.TConfiguration{}
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(serverConfig)
	transportFactory := thrift.NewTFramedTransportFactoryConf(thrift.NewTBufferedTransportFactory(8192), serverConfig)
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return nil, err
	}
	server = thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	err = server.Listen()
	if err != nil {
		return nil, err
	}
	go server.AcceptLoop()
	return server, nil
}

func StopServer(server *thrift.TSimpleServer) error {
	if server != nil {
		return server.Stop()
	}
	return nil
}
