抓取和验证免费http代理，提供远程或本地客户端

用法：
远程：
next := proxy.ProxyClient(ProxyServerURL, 6)
本地:
go proxy.BuildProxy() //抓取/验证服务
next := proxy.ProxyClient("", 6)

然后
client := &http.Client{
	Timeout: time.Second * 6, //可以在这里定义除代理外的其他client属性
}

for err := next(client) ; == nil; err = next(client) {
	resp, err != cleint.Get(URL)
}

服务器：
        go proxy.BuildProxy()
        proxy.Server(":8082")

GetAllProxy() 可以获取到本地的代理数据 返回url的字符串列表
GetAllProxyFromServer(URL) 从服务器获取代理数据 返回url的字符串列表

注意：即使通过服务验证，免费代理也无法百分百确认可用。如果超时或没得到想要的页面就换代理重试。
