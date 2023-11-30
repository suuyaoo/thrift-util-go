package thriftutil

import "testing"

func TestThriftAddress(t *testing.T) {
	if !IsThriftAddress("tcp4://127.0.0.1:8888") {
		t.Error("tcp4://127.0.0.1:8888")
	}

	if !IsThriftAddress("tcp6://[::1]:8888") {
		t.Error("tcp6://[::1]:8888")
	}

	if !IsThriftAddress("http://127.0.0.1:8888") {
		t.Error("http://127.0.0.1:8888")
	}

	if !IsThriftAddress("http://[::1]:8888") {
		t.Error("http://[::1]:8888")
	}

	if !IsThriftAddress("127.0.0.1:8888") {
		t.Error("127.0.0.1:8888")
	}

	if !IsThriftAddress("[::1]:8888") {
		t.Error("[::1]:8888")
	}

	if IsThriftAddress("8888") {
		t.Error("8888")
	}

	if IsThriftAddress("abc") {
		t.Error("abc")
	}

	if IsThriftAddress("tcp6://[::1:9999") {
		t.Error("tcp6://[::1:9999")
	}

	if IsThriftAddress("tcp6://::1]:9999") {
		t.Error("tcp6://::1]:9999")
	}

	if IsThriftAddress("tcp6://127.0.0.1:9999") {
		t.Error("tcp6://127.0.0.1:9999")
	}

	if IsThriftAddress("tcp4://[::1]:8888") {
		t.Error("tcp4://[::1]:8888")
	}

	{
		addr, ok := ParseThriftAddress("tcp6://[::1]:8888")
		if !ok || addr != "tcp6://[::1]:8888" {
			t.Error("tcp6://[::1]:8888")
		}
	}

	{
		addr, ok := ParseThriftAddress("[::1]:8888")
		if !ok || addr != "tcp6://[::1]:8888" {
			t.Error("[::1]:8888")
		}
	}

	{
		addr, ok := ParseThriftAddress("http://[::1]:8888")
		if !ok || addr != "tcp6://[::1]:8888" {
			t.Error("http://[::1]:8888")
		}
	}

	{
		addr, ok := ParseThriftAddress("127.0.0.1:8888")
		if !ok || addr != "tcp4://127.0.0.1:8888" {
			t.Error("127.0.0.1:8888")
		}
	}

	{
		addr, ok := ParseThriftAddress("http://127.0.0.1:8888")
		if !ok || addr != "tcp4://127.0.0.1:8888" {
			t.Error("http://127.0.0.1:8888")
		}
	}
}
