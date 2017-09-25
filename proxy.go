package proxy

import (
	"./get"
	"log"
	"time"
)

func BuildProxy() {
	MaxCon := make(chan struct{}, 50000) //这里做并发量限制
	for _, fun := range get.ProxyBuilder {
		go func(fun func() (*[]string, error)) {
			for {
				data, err := fun()
				log.Println("待验证：", len(*data))
				if err != nil {
					log.Println("Get ProxyBuilder:", err)
				}
				for i, proxy := range *data {
					time.Sleep(time.Millisecond * 20) //每个链接间隔
					MaxCon <- struct{}{}
					go func(proxy string) {
						defer func() { <-MaxCon }()
						if _, err := VerifyProxy(proxy); err != nil {
							log.Println("VerifyProxy:", i, proxy, ":", err)
						} else {
							ProxyMap.Store(proxy, ProxyInfo{})
							//*ProxyLi = append(*ProxyLi, proxy)
						}
					}(proxy)
				}
			}
		}(fun)
	}
}
