package core

const PREFIX string = "//visualize:"

/*
cannot make it const even though its fully dependant on
compile-time constants, for better incapsulation i will
make it global function that returns prefix in bytes
by that it will never be changed during runtime
*/
func PrefixBytes() []byte { return []byte(PREFIX) }

/*
const OPTIONS_POSSIBLE_MAP_DEFAULT_SIZE = 8
const OPTION_QUERY_DEFAULT_SIZE = 8
const DATA_ROWS_DEFAULT_SIZE = 16
*/
/* const ORIGINAL_FILE_DEFAULT_PATH = "../test/"
const ORIGINAL_FILE_DEFAULT_NAME = "toinsert.go" */

/* var METHODS_COUNT int = len(METHODS) */
/* var METHODS = []string{"start", "stop", "insert"} */
var REMAIN_COMMENTS = option{key: "remainComments", value: "true"}

const (
	START = iota
	STOP
	INSERT
)
