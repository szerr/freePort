package proxy

import (
	"sync"
)

type ProxyInfo struct{}

var (
	ProxyMap *sync.Map
)

func init() {
	ProxyMap = new(sync.Map)
}
