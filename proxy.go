package proxy

import (
	"./get"
	"log"
	"time"
)

func BuildProxy() {
	MaxCon := make(chan struct{}, 50000) //这里做并发量限制
	myip, err := TaobaoMyIp()
	if err != nil {
		log.Fatal("Get MyIP:", err)
	}
	for {
		data, err := get.GetProxy()
		log.Println("待验证：", len(*data))
		if err != nil {
			log.Fatal("Get ProxyBuilder:", err)
		}
		for _, proxy := range *data {
			time.Sleep(time.Millisecond * 10) //每个链接间隔
			MaxCon <- struct{}{}
			go func(proxy string) {
				defer func() { <-MaxCon }()
				if ip, err := VerifyProxy(proxy); err == nil {
					if ip != myip {
						ProxyMap.Store(proxy, ProxyInfo{})
					}
				}
			}(proxy)
		}
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
