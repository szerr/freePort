package proxy

import (
	"./get"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"
)

type IpInfo struct {
	Data struct {
		Ip string
	}
}

func BuildProxy() {
	MaxCon := make(chan struct{}, 10000) //这里做并发量限制
	for {
		for _, fun := range get.ProxyBuilder {
			data, err := fun()
			if err != nil {
				log.Println(err)
			}
			log.Println("待验证：", len(*data))
			for i, proxy := range *data {
				time.Sleep(time.Millisecond * 20) //每个链接间隔
				MaxCon <- struct{}{}
				go func(proxy string) {
					defer func() { <-MaxCon }()
					if ip, err := VerifyProxy(proxy); err != nil {
						log.Println("error:", i, proxy, ":", err)
					} else {
						log.Println("ip:", i, proxy, ":", ip)
						ProxyMap.Store(proxy, ProxyInfo{})
						//*ProxyLi = append(*ProxyLi, proxy)
					}
				}(proxy)
			}
		}
	}
}

func VerifyProxy(ProxyUrl string) (string, error) {
	resp, err := GetByProxy("http://ip.taobao.com/service/getIpInfo2.php?ip=myip", ProxyUrl)
	for i := 2; i > 0 && err != nil; i-- {
		resp, err = GetByProxy("http://ip.taobao.com/service/getIpInfo2.php?ip=myip", ProxyUrl)
	}
	//resp, err := GetByProxy("http://www.baidu.com", ProxyUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	info := new(IpInfo)
	if err := json.NewDecoder(resp.Body).Decode(info); err != nil {
		return "", err
	}
	return info.Data.Ip, nil
}

func GetByProxy(url_addr, proxy_addr string) (*http.Response, error) {
	request, _ := http.NewRequest("GET", url_addr, nil)
	proxy, err := url.Parse(proxy_addr)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	return client.Do(request)
}
