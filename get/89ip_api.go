package get

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func Get89ip() func() (*[]string, error) {
	//流年免费HTTP代理，
	API_URL := "http://www.89ip.cn/apijk/?&tqsl=4000&sxa=&sxb=&tta=&ports=&ktip=&cf=1" //
	dtime := Delay(int64(600))
	return func() (*[]string, error) {
		dtime()
		info := &[]string{}
		doc, err := goquery.NewDocument(API_URL)
		if err != nil {
			return info, err
		}
		reg := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+:\d+`)
		doc.Find("body").Each(func(i int, s *goquery.Selection) {
			data, err := s.Html()
			if err == nil {
				for _, addr := range reg.FindAllString(data, -1) {
					*info = append(*info, "http://"+addr)
				}
			}
		})
		return info, err
	}
}
