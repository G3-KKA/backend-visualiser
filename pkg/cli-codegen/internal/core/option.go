package core

type option struct {
	Key   string
	Value string
}

func MakeOption(key string, value string) option {
	return option{Key: key, Value: value}
}
