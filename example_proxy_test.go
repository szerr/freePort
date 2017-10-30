package proxy_test

import (
	"grab/proxy"
	"log"
	"net/http"
	"time"
)

//服务示例
func Example_server() {
	go proxy.BuildProxy()
	go proxy.Server(":8082")
	llend := 0
	log.Println("Start...")
	for {
		lend := 0
		time.Sleep(time.Second)
		proxy.ProxyMap.Range(func(key interface{}, value interface{}) bool {
			lend += 1
			return true
		})
		if lend != llend {
			log.Println("lend:", lend)
			llend = lend
		}
	}

}

//使用远程代理
func Example_client() {
	next := proxy.ProxyClient("http://127.0.0.1:8082", 3) //推荐从服务器获取代理数据
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	//next会迭代代理信息，填到client里。每次迭代取最新的
	//下边会一直换代理请求
	for err := next(client); ; err = next(client) {
		if err != nil {
			log.Fatal(err)
		}
		client.Get("http://www.icanhazip.com/")
	}
}

//用本地服务获取ip
func Example_local() {
	go proxy.BuildProxy() //抓取/验证服务
	next := proxy.ProxyClient("", 6)
	client := &http.Client{Timeout: time.Second * 6} //可以在这里定义除代理外的其他client属性 }
	URL := "http://www.baidu.com"
	for err := next(client); err == nil; err = next(client) {
		resp, err := client.Get(URL)
		if err != nil {
			log.Println(err)
			continue
		} else {
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println(len(data))
		}
		time.Sleep(time.Second * 1)
	}
}
