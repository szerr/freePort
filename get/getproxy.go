package get

import (
	"errors"
	"log"
)

type t struct{}

func GetProxy() (*[]string, error) {
	mli := make(map[string]t)
	li := &[]string{}
	for funame, fun := range ProxyBuilder {
		li, err := fun()
		if err != nil {
			log.Println("Waring:", funame, err)
		}
		log.Println("GetProxy:", funame, len(*li))
		for _, i := range *li {
			mli[i] = t{}
		}
	}
	for k, _ := range mli {
		*li = append(*li, k)
	}
	if len(*li) > 0 {
		return li, nil
	}
	return li, errors.New("proxy num 0!")
}
