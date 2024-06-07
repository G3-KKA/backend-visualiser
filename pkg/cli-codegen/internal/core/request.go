package core

import (
	generr "backend-visualiser/cli-codegen/internal/errors/codegenError"
	"os"
	"slices"
)

type Request struct {
	Method string
	Name   string
	Query  []option
	Data   [][]byte
}

func NewRequest() Request {
	return Request{
		Name:   "",
		Method: "",
		Query:  make([]option, 0, OPTION_QUERY_DEFAULT_SIZE),
		Data:   make([][]byte, 0, DATA_ROWS_DEFAULT_SIZE),
	}
}

func (req *Request) ReadLine(in []byte) error {
	tmp := slices.Clip(in)
	tmp = append(tmp, '\n')
	req.Data = append(req.Data, tmp)
	return nil
}
func (req *Request) InsertInto(out *os.File) error {
	//__TODO
	// this should seek in the data file for name ,
	// then insert this exact data to the replace file
	out.Write([]byte("\tfmt.Print(`"))
	for _, row := range req.Data {
		inserted, err := out.Write(row)
		if inserted == 0 || err != nil {
			return generr.Err("Failed to write to file", err)
		}
	}
	out.Write([]byte("`)\n"))
	return nil
}
func (req *Request) Reset() {
	//__TODO
	// reset should be call'd at the end of the req.stop()
	// at the time of writing this comment its call'd at the end of the insert(),
	// which is drastically shirnkens variants of use to
	// start > stop > insert,
	// i want it to have ability of multiple insertions of the same data
	// and
	// multiple start > stop call's in a row before any if any insert happens
	tmp := NewRequest()
	*req = tmp
}
func (req *Request) Stop() {
	req.Data = slices.Delete(req.Data, len(req.Data)-1, len(req.Data))
	//__TODO handle :start?OptionS&OptionS
	// . . . handling . . .
	//__TODO write .data to file
}
