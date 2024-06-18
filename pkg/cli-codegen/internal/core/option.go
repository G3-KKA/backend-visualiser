package core

import (
	generr "backend-visualiser/cli-codegen/internal/errors/codegenError"
	"backend-visualiser/cli-codegen/internal/errors/current"
	"bytes"
	"log"
)

type BigOption struct {
	Options       option
	OptionHandler func(*Request, option) error
}
type option struct {
	key   string
	value string
}

//	These two functions are used for incapsulation of fields
//	There is not a single scenario where you need
//	to change key/value after they was set up

func (o *option) Key() string {
	return o.key
}
func (o *option) Value() string {
	return o.value
}

func MakeOption(key string, value string) option {
	return option{key: key, value: value}
}

func MakeOptionFromRaw(raw []byte) option {
	const (
		key   = 0
		value = 1
	)
	pair := bytes.Split(raw, []byte("="))
	if len(pair) != 2 {
		log.Fatal(generr.Err("Wrong query options, must be key=value"+current.Phase, nil))
	}

	return MakeOption(string(pair[key]), string(pair[value]))
}
