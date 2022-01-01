package gosupport

type Metadata map[string]string

func (m Metadata) Get(key string) (value string, ok bool) {
	value, ok = m[key]
	return
}

func (m Metadata) Set(key string, value string) {
	m[key] = value
}

func (m Metadata) Del(key string) {
	delete(m, key)
}

func (m Metadata) Clone() StrMaper {
	c := make(Metadata)
	for k, v := range m {
		c[k] = v
	}
	return c
}
