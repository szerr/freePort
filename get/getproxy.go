package get

type t struct{}

func GetProxy() (*[]string, error) {
	mli := make(map[string]t)
	li := &[]string{}
	for _, fun := range ProxyBuilder {
		li, err := fun()
		if err != nil {
			return li, err
		}
		for _, i := range *li {
			mli[i] = t{}
		}
	}
	for k, _ := range mli {
		*li = append(*li, k)
	}
	return li, nil
}
