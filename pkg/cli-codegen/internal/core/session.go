package core

type Session struct {
	options []map[option]OptionHandler
}

func NewSession() *Session {
	return &Session{options: make([]map[option]OptionHandler, METHODS_COUNT)}
}
func (sess *Session) InitOptions() error {
	sess.options[START] = initSTART()
	sess.options[STOP] = initSTOP()
	sess.options[INSERT] = initINSERT()
	return nil
}
func initSTART() map[option]OptionHandler {
	tmp := make(map[option]OptionHandler, OPTIONS_POSSIBLE_MAP_DEFAULT_SIZE)
	return tmp
}
func initSTOP() map[option]OptionHandler {
	tmp := make(map[option]OptionHandler, OPTIONS_POSSIBLE_MAP_DEFAULT_SIZE)
	return tmp
}
func initINSERT() map[option]OptionHandler {
	tmp := make(map[option]OptionHandler, OPTIONS_POSSIBLE_MAP_DEFAULT_SIZE)
	return tmp
}
