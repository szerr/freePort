package get

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func Shifengsoft() func() (*[]string, error) {
	//拾风HTTP代理api，有效ip有5000+ 这里提取5k。官方说每台电脑 每分钟内提取次数超过60次封IP。
	API_URL := `http://ip.shifengsoft.com/get.php?tqsl=6000&submit=%CC%E1++%C8%A1`
	dtime := Delay(int64(600))
	return func() (*[]string, error) {
		dtime()
		info := &[]string{}
		doc, err := goquery.NewDocument(API_URL)
		if err != nil {
			return info, err
		}
		reg := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+:\d.+`)
		doc.Find("body .mass").Each(func(i int, s *goquery.Selection) {
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
