package get

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

func Ip181() func() (*[]string, error) {
	URL := "http://www.ip181.com/daili/"
	PAGE_NUM := 20
	dtime := Delay(300)
	analyze := func(doc *goquery.Document) *[]string {
		info := &[]string{}
		doc.Find(".table-hover tbody tr").Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}
			data := s.Text()
			li := strings.Split(data, "\n")
			addr := strings.Replace("http://"+li[1]+":"+li[2], " ", "", -1)
			*info = append(*info, addr)
		})
		return info
	}
	return func() (*[]string, error) {
		dtime()
		info := &[]string{}
		doc, err := goquery.NewDocument("http://www.ip181.com/")
		if err != nil {
			return info, err
		}
		*info = append(*info, *analyze(doc)...)
		for page := 1; page <= PAGE_NUM; page += 1 {
			doc, err := goquery.NewDocument(URL + strconv.Itoa(page) + ".html")
			if err != nil {
				return info, err
			}
			*info = append(*info, *analyze(doc)...)
		}
		return info, nil
	}
}
