package get

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

func Coobobo() func() (*[]string, error) {
	/*这里只抓国内的*/
	URL := "http://www.coobobo.com/free-http-proxy/"
	PAGE_NUM := 10
	dtime := Delay(300)
	return func() (*[]string, error) {
		dtime()
		info := &[]string{}
		for page := 1; page <= PAGE_NUM; page += 1 {
			//doc, err := goquery.NewDocument(URL + strconv.Itoa(page))
			doc, err := TimeOutDoc(URL+strconv.Itoa(page), 3)
			if err != nil {
				return info, err
			}
			doc.Find(".table-bordered tbody tr").Each(func(i int, s *goquery.Selection) {
				data := s.Text()
				li := strings.Split(data, "\n")
				if strings.Contains(li[9], "中国") {
					ipinfo := strings.Split(li[3], `"`)
					addr := "http://" + ipinfo[1] + "." + ipinfo[3] + ipinfo[5] + ":" + strings.Replace(li[6], " ", "", -1)
					*info = append(*info, addr)
				}
			})
		}
		PAGE_NUM = 2
		return info, nil
	}
}
