package proxy

import (
	"testing"
)

func TestGetAllProxyFromServer(t *testing.T) {
	return
	data, err := getAllProxyFromServer("http://127.0.0.1:8082")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(len(data))
	}
}

func TestRangeProxy(t *testing.T) {
	return
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

func TestProxyClient(t *timeing.T) {
	next := ProxyClient("http://127.0.0.1:8082", time.Second)
	var client *http.Client
	if err := next(client); err != nil {
		t.Error(err)
	} else {
		resp, err := client.Get("http://ip.cn")
		if err != nil {
			t.Error(err)
		} else {
			defer resp.Body.Close()
			log.Prlntli(string(ioutil.ReadAll(resp.Body)))
		}
	}
}
