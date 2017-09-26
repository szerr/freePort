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
		for i, proxy := range *data {
			time.Sleep(time.Millisecond * 10) //每个链接间隔
			MaxCon <- struct{}{}
			go func(proxy string) {
				defer func() { <-MaxCon }()
				if ip, err := VerifyProxy(proxy); err != nil {
					log.Println("VerifyProxy:", i, proxy, ":", err)
				} else {
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
