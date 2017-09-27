package get

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

func Superfastip() func() (*[]string, error) {
	URL := "http://www.superfastip.com/welcome/getips/"
	PAGE_NUM := 20
	dtime := Delay(300)
	return func() (*[]string, error) {
		dtime()
		info := &[]string{}
		doc, err := goquery.NewDocument(URL)
		if err != nil {
			return info, err
		}
		doc.Find(".pagination").Each(func(i int, s *goquery.Selection) { //找页数
			data := strings.Replace(strings.Split(s.Text(), "\n")[5], " ", "", -1)
			PAGE_NUM, _ = strconv.Atoi(data)
		})
		for page := 1; page <= PAGE_NUM; page += 1 {
			doc, err := goquery.NewDocument(URL + strconv.Itoa(page))
			if err != nil {
				return info, err
			}
			doc.Find(".table-bordered tbody tr").Each(func(i int, s *goquery.Selection) {
				data := strings.Replace(s.Text(), " ", "", -1)
				li := strings.Split(data, "\n")
				addr := li[6] + "://" + li[2] + ":" + li[3]
				*info = append(*info, addr)
			})
		}
		PAGE_NUM = 1
		return info, nil
	}
}
