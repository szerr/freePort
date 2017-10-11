package proxy_test

import (
	"grab/proxy"
	"log"
	"net/http"
	"time"
)

//抓取和验证免费http代理，提供远程或本地客户端
func Example_client() {
	next := proxy.ProxyClient("", 3) //使用本地的代理数据
	//next := proxy.ProxyClient("http://127.0.0.1:8082", 3) //推荐从服务器获取代理数据
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
