package get

var ProxyBuilder map[string]func() (*[]string, error)

func init() {
	ProxyBuilder = map[string]func() (*[]string, error){
		"Shifengsoft": Shifengsoft(),
		//"SuperfastipApi": SuperfastipApi(),
		"Get89ip":     Get89ip(),
		"Coobobo":     Coobobo(),
		"Superfastip": Superfastip(),
	}
}
