package proxy

import ()

type ProxyInfo struct{}

var (
	ProxyMap *SyncMap
)

func init() {
	ProxyMap = NewSyncMap()
}
