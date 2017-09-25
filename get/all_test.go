package get

import (
	"testing"
)

func TestAll(t *testing.T) {
	/*
		AllFun := map[string]func() func() (*[]string, error){
			"Shifengsoft": Shifengsoft,
			"Superfastip": Superfastip,
			"Get89ip":     Get89ip,
			"Coobobo":     Coobobo,
		}
	*/
	AllFun := map[string]func() func() (*[]string, error){"Coobobo": Coobobo}
	for funame, fun := range AllFun {
		fu := fun()
		data, err := fu()
		if err != nil {
			t.Error(funame, err)
		}
		if len(*data) == 0 {
			t.Error("! num 0:", funame)
		}
		t.Log(funame, "num:", len(*data))
	}
}
