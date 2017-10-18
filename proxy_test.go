package proxy

import (
	"log"
	"testing"
	"time"
)

func TestBuildProxy(t *testing.T) {
	BuildProxy()
	for i := 0; i < 580; i++ {
		log.Println(i, len(*ProxyLi))
		time.Sleep(time.Second)
	}
}
