package thriftutil

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

func IsThriftAddress(addr string) bool {
	if len(addr) < 7 {
		return false
	}
	if addr[0:7] == "http://" {
		return true
	}
	if addr[0:7] == "tcp4://" {
		return true
	}
	if addr[0:7] == "tcp6://" {
		return true
	}
	return false
}
