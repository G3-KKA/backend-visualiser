package core

import (
	generr "backend-visualiser/cli-codegen/internal/errors/codegenError"
	"backend-visualiser/cli-codegen/internal/errors/current"
	"bytes"
	"log"
)

/*
	 TODO
		Идеальный дизайн это когда
		 для добавления фитчи
		 нужно модифицировать
		 лишь одно место в коде.

		Хочу сделать cлайс BigOption'ов к которому будут всегда обращаться
		 когда нужно модифицировать состояние реквеста.
		 Отложу т.к. нужно многое переписывать
		PS разобраться как хендлить имя, возможно оно вообще
		 не должно быть частью опций, а должно быть явно указано
		 при объявлении
		 --visualise:__METHOD__:name=__NAME__?__QUERY__
*/
type BigOption struct {
	Options       option
	OptionHandler func(*Request, option) error
}
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
