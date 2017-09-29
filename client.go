package proxy

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"
)

func getAllProxyFromServer(url string) ([]string, error) {
	data := []string{}
	resp, err := (&http.Client{Timeout: 3}).Get(url + "/proxy")
	for i := 2; i > 0 && err != nil; i-- {
		resp, err = (&http.Client{Timeout: time.Second * 3}).Get(url + "/proxy")
	}
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func RangeProxy(getAllProxy func() ([]string, error), delayTime int) func() (string, error) {
	/*需要一个获取代理数据的回调函数
	delayTime为每次循环获取的最小间隔时间，防止同一ip调用间隔过短
	*/
	dtime := Delay()
	var data []string
	lend := 0
	var err error
	return func() (string, error) {
		if lend == 0 {
			dtime(delayTime)
			data, err = getAllProxy()
			if err != nil {
				return "", err
			}
			lend = len(data)
			if lend == 0 {
				return "", errors.New("RangeProxy: getAllProxy return null data")
			}
		}
		lend -= 1
		return data[lend], nil
	}
}

func PackGetAllProxy() ([]string, error) {
	return GetAllProxy(), nil
}

func PackGetAllProxyFromServer(server_addr string) func() ([]string, error) {
	return func() ([]string, error) { return getAllProxyFromServer(server_addr) }
}

func ProxyClient(server_addr string) func(client *http.Client) error {
	var next func() (string, error)
	if server_addr == "" {
		go BuildProxy()
		time.Sleep(time.Second * 5)
		for len(GetAllProxy()) > 0 { //等待代理获取和测试
			time.Sleep(time.Second)
		}
		next = RangeProxy(PackGetAllProxy, 2)
	} else {
		next = RangeProxy(PackGetAllProxyFromServer(server_addr), 2)
	}
	return func(client *http.Client) error {
		proxy, err := next()
		if err != nil {
			return err
		}
		proxy_url, err := url.Parse(proxy)
		if err != nil {
			return err
		}
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxy_url),
		}
		return err
	}
}
