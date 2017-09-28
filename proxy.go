package proxy

import (
	"./get"
	"errors"
	"log"
	"sync"
	"time"
)

func BuildProxy() {
	MaxCon := make(chan struct{}, 50000) //这里做并发量限制
	myip, err := TaobaoMyIp()
	if err != nil {
		log.Fatal("Get MyIP:", err)
	}
	for {
		proxyM := new(sync.Map)
		data, err := get.GetProxy()
		if err != nil {
			log.Fatal("Get ProxyBuilder:", err)
		}
		*data = append(*data, GetAllProxy()...) //已经保存的代理也重新验证一次
		log.Println("GetProxy: all", len(*data))
		for _, proxy := range *data {
			time.Sleep(time.Millisecond * 10) //每个连接间隔
			MaxCon <- struct{}{}
			go func(proxy string) {
				if ip, err := VerifyProxy(proxy); err == nil {
					if ip != myip {
						proxyM.Store(proxy, ProxyInfo{})
					}
				}
			}(proxy)
		}
		ProxyMap = proxyM
		log.Println("Sleep")
		time.Sleep(time.Second * 60 * 10) //获取代理间隔
	}
}

func GetAllProxy() []string {
	li := []string{}
	ProxyMap.Range(func(key, value interface{}) bool {
		li = append(li, key.(string))
		return true
	})
	return li
}

func DeleteProxy(key string) {
	ProxyMap.Delete(key)
}
