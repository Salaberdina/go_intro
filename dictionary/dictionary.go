package dictionary

import "errors"

func Search(d map[string]string, x string) string {
	return d[x]
}

type Dictionary map[string]string

func (d Dictionary) Search(a string) string {
	return d[a]
}

func (d Dictionary) Add(t, s string) error {
	_, ok := d[t]
	if ok {
		return errors.New("error")
	}
	d[t] = s
	return nil
}

func (d Dictionary) Update(g, j string) {
	d[g] = j
}

func (d Dictionary) Delete(l string) {
	delete(d, l)
}
