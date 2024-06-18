package core

import "github.com/spf13/viper"

//
//
//
//
//
//
//
//	This code is unused for now
//
//
//
//
//
//
//

type Session struct {
	options []map[option]OptionHandler
}

func NewSession() *Session {
	return &Session{options: make([]map[option]OptionHandler, len(viper.GetStringSlice("methods")))}
}
func (sess *Session) InitOptions() error {
	//Setting up possible visualise:__OPTION__
	sess.options[START] = initSTART()
	sess.options[STOP] = initSTOP()
	sess.options[INSERT] = initINSERT()

	return nil
}
func initSTART() map[option]OptionHandler {
	tmp := make(map[option]OptionHandler, viper.GetInt("options_size"))
	return tmp
}
func initSTOP() map[option]OptionHandler {
	tmp := make(map[option]OptionHandler, viper.GetInt("options_size"))
	return tmp
}
func initINSERT() map[option]OptionHandler {
	tmp := make(map[option]OptionHandler, viper.GetInt("options_size"))
	return tmp
}
