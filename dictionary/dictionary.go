package dictionary

func Search(d map[string]string, x string) string {
	return d[x]
}

type Dictionary map[string]string

func (d Dictionary) Search(a string) string {
	return d[a]
}
