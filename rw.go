package thriftutil

import (
	"context"

	"github.com/apache/thrift/lib/go/thrift"
)

type ThriftMsg interface {
	Read(ctx context.Context, iprot thrift.TProtocol) error
	Write(ctx context.Context, oprot thrift.TProtocol) error
}

func Marshal(msg ThriftMsg) ([]byte, error) {
	memBuf := thrift.NewTMemoryBuffer()
	binProto := thrift.NewTBinaryProtocolConf(memBuf, &thrift.TConfiguration{})
	err := msg.Write(context.TODO(), binProto)
	if err != nil {
		return nil, err
	}
	return memBuf.Bytes(), nil
}

func UnMarshal(msg ThriftMsg, buf []byte) error {
	memBuf := thrift.NewTMemoryBuffer()
	memBuf.Write(buf)
	binProto := thrift.NewTBinaryProtocolConf(memBuf, &thrift.TConfiguration{})
	err := msg.Read(context.TODO(), binProto)
	if err != nil {
		return err
	}
	return nil
}

func Clone(src ThriftMsg, dst ThriftMsg) error {
	buf, err := Marshal(src)
	if err == nil {
		err = UnMarshal(dst, buf)
	}
	return err
}
