package get

import (
	"encoding/json"
	"net/http"
	//	"time"
)

type superfastipInfo struct {
	Desc string
	Data [][]string
}

func SuperfastipApi() func() (*[]string, error) {
	/*官方地址 http://www.superfastip.com*/
	API_URL := "http://superfastip.com/welcome/getapi"
	UPDATE_TIME := 7200 //官方2小时更新一次 秒
	dtime := Delay(int64(UPDATE_TIME))
	return func() (*[]string, error) {
		dtime()
		info := []string{}
		resp, err := http.Get(API_URL)
		if err != nil {
			return &info, err
		}
		defer resp.Body.Close()
		sinfo := new(superfastipInfo)
		if err = json.NewDecoder(resp.Body).Decode(sinfo); err != nil {
			return &info, err
		}
		for _, i := range sinfo.Data {
			info = append(info, i[5]+"://"+i[1]+":"+i[2])
		}
		return &info, err
	}
}
