package core

type Session struct {
	options []map[option]OptionHandler
}

func NewSession() *Session {
	return &Session{options: make([]map[option]OptionHandler, METHODS_COUNT)}
}
func (s *Session) InitOptions() error {
	s.options[START] = initSTART()
	s.options[STOP] = initSTOP()
	s.options[INSERT] = initINSERT()
	return nil
}
func initSTART() map[option]OptionHandler {
	tmpmap := make(map[option]OptionHandler, OPTIONS_POSSIBLE_MAP_DEFAULT_SIZE)
	return tmpmap
}
func initSTOP() map[option]OptionHandler {
	tmpmap := make(map[option]OptionHandler, OPTIONS_POSSIBLE_MAP_DEFAULT_SIZE)
	return tmpmap
}
func initINSERT() map[option]OptionHandler {
	tmpmap := make(map[option]OptionHandler, OPTIONS_POSSIBLE_MAP_DEFAULT_SIZE)
	return tmpmap
}
