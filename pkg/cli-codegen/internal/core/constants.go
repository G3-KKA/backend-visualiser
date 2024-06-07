package core

const PREFIX string = "//visualize:"
const OPTIONS_POSSIBLE_MAP_DEFAULT_SIZE = 8
const OPTION_QUERY_DEFAULT_SIZE = 8
const DATA_ROWS_DEFAULT_SIZE = 16
const ORIGINAL_FILE_DEFAULT_PATH = "./"
const ORIGINAL_FILE_DEFAULT_NAME = "toinsert.go"
const QUERY_EXIST = true

var METHODS_COUNT int = len(METHODS)
var METHODS = []string{"start", "stop", "insert"}
var REMAIN_COMMENTS = option{Key: "remainComments", Value: "true"}
var PREFIX_BYTES []byte = []byte(PREFIX)

const (
	START = iota
	STOP
	INSERT
)
