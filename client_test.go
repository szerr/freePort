package proxy

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestGetAllProxyFromServer(t *testing.T) {
	return
	data, err := GetAllProxyFromServer("http://127.0.0.1:8082")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(len(data))
	}
}

func TestRangeProxy(t *testing.T) {
	return
	next := RangeProxy(func() ([]string, error) {
		return GetAllProxyFromServer("http://127.0.0.1:8082")
	}, 1)
	_, err := next()
	if err != nil {
		t.Error(err)
	}
}

func TestProxyClient(t *testing.T) {
	next := ProxyClient("http://127.0.0.1:8082", 1)
	var client *http.Client
	if err := next(client); err != nil {
		t.Error(err)
	} else {
		resp, err := client.Get("http://ip.cn")
		if err != nil {
			t.Error(err)
		} else {
			defer resp.Body.Close()
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			} else {
				log.Println(string(data))
			}
		}
	}
}
