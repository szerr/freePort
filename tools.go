package proxy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func VerifyProxy(ProxyUrl string) (string, error) {
	//return ICIp(ProxyUrl)
	return TaobaoIp(ProxyUrl)
}

func ICIp(ProxyUrl string) (string, error) {
	resp, err := GetByProxy("http://www.icanhazip.com/", ProxyUrl)
	for i := 2; i > 0 && err != nil; i-- {
		resp, err = GetByProxy("http://www.icanhazip.com/", ProxyUrl)
	}
	//resp, err := GetByProxy("http://www.baidu.com", ProxyUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
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
		Timeout: time.Second * 3,
	}
	return client.Do(request)
}

type IpInfo struct {
	Data struct {
		Ip string
	}
}

func TaobaoIp(ProxyUrl string) (string, error) {
	resp, err := GetByProxy("http://ip.taobao.com/service/getIpInfo2.php?ip=myip", ProxyUrl)
	for i := 1; i > 0 && err != nil; i-- {
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

func TaobaoMyIp() (string, error) {
	resp, err := http.Get("http://ip.taobao.com/service/getIpInfo2.php?ip=myip")
	for i := 2; i > 0 && err != nil; i-- {
		resp, err = http.Get("http://ip.taobao.com/service/getIpInfo2.php?ip=myip")
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

func Sleep2Time(STime int64) int64 { //睡眠到某个时间戳，返回当前时间戳
	time.Sleep(time.Duration(STime-time.Now().Unix()) * time.Second)
	return time.Now().Unix()
}

func Delay(STime int64) func() { //延迟返回某个时间，单位是秒
	var ptime int64 = 0
	return func() {
		Sleep2Time(ptime)
		ptime = time.Now().Unix() + STime
	}
}
