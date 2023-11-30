package thriftutil

import (
	"net"
	"strconv"
	"strings"
)

func ThriftAddress(addr string) string {
	if len(addr) < 7 {
		return addr
	}
	if addr[0:7] == "http://" {
		return addr[7:]
	}
	if addr[0:7] == "tcp4://" {
		return addr[7:]
	}
	if addr[0:7] == "tcp6://" {
		return addr[7:]
	}
	return addr
}

func checkAddress(prep, addr string) bool {
	pos := strings.LastIndex(addr, ":")
	if pos == -1 {
		return false
	}
	prepStr := ""
	ipStr := addr[0:pos]
	portStr := addr[pos+1:]
	if strings.Contains(ipStr, ".") {
		prepStr = "tcp4://"
	} else if strings.Contains(ipStr, ":") {
		prepStr = "tcp6://"
		if ipStr[0] == '[' {
			ipLen := len(ipStr)
			if ipStr[ipLen-1] == ']' {
				ipStr = ipStr[1 : ipLen-2]
			} else {
				return false
			}
		}
	}
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	if prep != "" && prepStr != prep {
		return false
	}
	port, err := strconv.ParseInt(portStr, 10, 16)
	if err != nil {
		return false
	}
	if port < 0 {
		return false
	}
	return true
}

func IsThriftAddress(addr string) bool {
	if len(addr) < 7 {
		return checkAddress("", addr)
	}
	if addr[0:7] == "http://" {
		return checkAddress("", addr[7:])
	}
	if addr[0:7] == "tcp4://" {
		return checkAddress("tcp4://", addr[7:])
	}
	if addr[0:7] == "tcp6://" {
		return checkAddress("tcp6://", addr[7:])
	}
	return checkAddress("", addr)
}
