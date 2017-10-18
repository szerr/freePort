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

for err := next(client) ; err == nil; err = next(client) {
	resp, err := cleint.Get(URL)
}

服务器：
        go proxy.BuildProxy()
        proxy.Server(":8082")


注意：即使通过服务验证，免费代理也无法百分百确认可用。如果超时或没得到想要的页面就换代理重试。
