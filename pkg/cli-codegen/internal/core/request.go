package core

import (
	generr "backend-visualiser/cli-codegen/internal/errors/codegenError"
	"bufio"
	"log"
	"os"
	"slices"

	"github.com/spf13/viper"
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
		Query:  make([]option, 0, viper.GetInt("query_size")),
		Data:   make([][]byte, 0, viper.GetInt("data_rows_size")),
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
	req.setNameFromOptions()
	req.ReadDataFile()
	lastRow := req.Data[len(req.Data)-1]
	req.Data[len(req.Data)-1] = lastRow[:len(lastRow)-1]
	out.Write([]byte("\tfmt.Print(`"))
	for _, row := range req.Data {
		inserted, err := out.Write(row)
		if inserted == 0 || err != nil {
			return generr.Err("Failed to write to file", err)
		}
	}
	out.Write([]byte("`)\n"))
	req.Reset()
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
	/* this makes a lot of allcotions, i might just default some of values? */
}
func (req *Request) Stop() {
	req.deleteLastRow()

	/* query handling , abadon later when */
	req.setNameFromOptions()
	createDataFolder()
	/*__TODO: there i should create symlink to tmp/__NAME__.data , for use in name=latest()*/
	datafile := createDataFile(req.Name)
	for _, row := range req.Data {
		datafile.Write(row)
	}
	defer datafile.Close()

	//__TODO handle :start?OptionS&OptionS
	// . . . handling . . .
	//__TODO write .data to file
	req.Reset()
}
func (req *Request) ReadDataFile() *os.File {
	datafile, err := os.Open(viper.GetString("tmpdir") + "/" + req.Name + ".data")
	if err != nil {
		log.Fatal(generr.Err("Failed to open data file: "+viper.GetString("tmpdir")+"/"+req.Name+".data", err))
	}
	defer datafile.Close()
	scanner := bufio.NewScanner(datafile)
	for rowIdxidx := 0; scanner.Scan(); rowIdxidx++ {
		req.Data = append(req.Data, slices.Concat(scanner.Bytes(), []byte("\n")))
	}
	return datafile

}

/* ===  utilitary functions === */
func (req *Request) deleteLastRow() {
	req.Data = slices.Delete(req.Data, len(req.Data)-1, len(req.Data))
}
func createDataFolder() {
	err := os.Mkdir(viper.GetString("tmpdir"), os.ModePerm)
	if err != nil {
		log.Println("TODO: handle if tmp dir exists", err)
	}
}
func createDataFile(name string) *os.File {
	tmpf, err := os.Create(viper.GetString("tmpdir") + "/" + name + ".data")
	if err != nil {
		log.Println("TODO: handle if tmp/__NAME__.data file exists", err)
	}
	return tmpf
}
func (req *Request) setNameFromOptions() {
	for _, o := range req.Query {
		if o.Key() == "name" {
			req.Name = o.Value()
		}
	}
}
