package core

const PREFIX string = "//visualize:"

/*
cannot make it const even though its fully dependant on
compile-time constants, for better incapsulation i will
make it global function that returns prefix in bytes
by that it will never be changed during runtime
*/
func PrefixBytes() []byte { return []byte(PREFIX) }

var REMAIN_COMMENTS = option{key: "remainComments", value: "true"}

const (
	START = iota
	STOP
	INSERT
)
