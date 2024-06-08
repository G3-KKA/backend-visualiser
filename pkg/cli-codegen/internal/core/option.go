package core

import (
	generr "backend-visualiser/cli-codegen/internal/errors/codegenError"
	"backend-visualiser/cli-codegen/internal/errors/current"
	"bytes"
	"log"
)

type option struct {
	key   string
	value string
}

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
	bytes := bytes.Split(raw, []byte("="))
	if len(bytes) != 2 {
		log.Fatal(generr.Err("Wrong query options, must be key=value"+current.Phase, nil))
	}

	return option{key: string(bytes[0]), value: string(bytes[1])}
}
