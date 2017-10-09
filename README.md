抓取和验证免费代理，提供远程或本地客户端

用法：
远程：
next := proxy.ProxyClient(ProxyServerURL, 6)
本地:
go proxy.BuildProxy() //抓取/验证服务
next := proxy.ProxyClient("", 6)

然后
client := &http.Client{
	Timeout: time.Second * 6,
}

for err := next(client) ; == nil; err = next(client) {
	resp, err != cleint.Get(URL)
}
