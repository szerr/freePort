package get

import (
	"testing"
)

func TestAll(t *testing.T) {
	//ProxyBuilder := map[string]func() func() (*[]string, error){"Superfastip": Ip181}
	for funame, fun := range ProxyBuilder {
		data, err := fun()
		if err != nil {
			t.Error(funame, err)
		}
		if len(*data) == 0 {
			t.Error("! num 0:", funame)
		} else {
			t.Log(funame, "num:", len(*data))
		}
	}
}
