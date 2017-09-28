package proxy

import (
	"testing"
)

func TestGetAllProxyFromServer(t *testing.T) {
	data, err := getAllProxyFromServer("http://127.0.0.1:8082")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(len(data))
	}
}

func TestRangeProxy(t *testing.T) {
	next := RangeProxy(func() ([]string, error) {
		return getAllProxyFromServer("http://127.0.0.1:8082")
	}, 1)
	var err error
	var data string
	_, err = next()
	if err != nil {
		t.Error(err)
	}
}
