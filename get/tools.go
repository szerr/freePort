package get

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
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

func TimeOutDoc(url string, timeout int) (*goquery.Document, error) {
	var doc *goquery.Document
	resp, err := (&http.Client{Timeout: time.Second * time.Duration(timeout)}).Get(url)
	if err != nil {
		return doc, err
	}
	doc, err = goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return doc, err
	}
	return doc, err
}
