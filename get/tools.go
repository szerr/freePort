package get

import (
	"time"
)

func Sleep2Time(STime int64) int64 { //睡眠到某个时间戳，返回当前时间戳
	time.Sleep(time.Duration(STime-time.Now().Unix()) * time.Second)
	return time.Now().Unix()
}

func Delay(STime int64) func() { //延迟返回某个时间，单位是秒
	var ptime int64 = 0
	return func() {
		Sleep2Time(ptime)
		ptime = time.Now().Unix() + STime
	}
}
