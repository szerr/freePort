package get

import (
	"testing"
)

func TestAll(t *testing.T) {
	/*AllFun := map[string]func() func() (*[]string, error){
		"Shifengsoft":    Shifengsoft,
		"SuperfastipApi": SuperfastipApi,
		"Get89ip":        Get89ip,
		"Coobobo":        Coobobo,
		"Superfastip":    Superfastip,
	}*/
	AllFun := map[string]func() func() (*[]string, error){"Superfastip": Ip181}
	for funame, fun := range AllFun {
		fu := fun()
		data, err := fu()
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

func TestGetProxy(t *testing.T) {
	data, err := GetProxy()
	if err != nil {
		t.Error(err)
	}
	t.Log("All Proxy num: ", len(*data))
}
