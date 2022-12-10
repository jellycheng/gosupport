package xlanguage

//单个语言的key-vaule集合
type LangSet map[string]string

func (l LangSet) Add(key, value string) {
	l[key] = value
}

func (l LangSet) Get(key string) string {
	if ret, ok := l[key]; ok {
		return ret
	}
	return ""
}

func (l LangSet) Combine(set LangSet) LangSet {
	for k, s := range set {
		l[k] = s
	}
	return l
}

type LangMap map[string]LangSet
