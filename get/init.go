package get

var ProxyBuilder []func() (*[]string, error)

func init() {
	ProxyBuilder = []func() (*[]string, error){
		//		Superfastip(),
		//		Shifengsoft(),
		//		Get89ip(),
		Coobobo(),
	}
}
