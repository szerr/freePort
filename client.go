package proxy

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

func getAllProxyFromServer(url string) ([]string, error) {
	data := []string{}
	resp, err := (&http.Client{Timeout: 3}).Get(url + "/proxy")
	for i := 2; i > 0 && err != nil; i-- {
		resp, err = (&http.Client{Timeout: time.Second * 3}).Get(url + "/proxy")
	}
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func RangeProxy(getAllProxy func() ([]string, error), delayTime int) func() (string, error) {
	/*需要一个获取代理数据的回调函数
	delayTime为每次循环获取的最小间隔时间，防止同一ip调用间隔过短
	*/
	dtime := Delay()
	var data []string
	i := 0
	var err error
	return func() (string, error) {
		if i == 0 {
			dtime(delayTime)
			data, err = getAllProxy()
			if err != nil {
				return "", err
			}
			i = len(data)
			if i == 0 {
				return "", errors.New("RangeProxy: getAllProxy return null data")
			}
		}
		i -= 1
		log.Println(len(data), i)
		return data[i], nil
	}
}
