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

	//Making shure that we are working with brand new data after
	//	append or other modification
	//TODO: are we even care that we are modifying in []byte
	tmp := slices.Clip(in)

	//Creating brand new slice just
	//	because scanner returns []byte without new line
	//TODO: there should be a way to get []byte with new line
	tmp = append(tmp, '\n')
	req.Data = append(req.Data, tmp)

	return nil
}
func (req *Request) InsertInto(out *os.File) error {

	//Preparing request to be loaded in memory
	req.setNameFromOptions()
	req.ReadDataFile()

	//WTF is this?
	//TODO: Rewrite this in more readable way
	lastRow := req.Data[len(req.Data)-1]
	req.Data[len(req.Data)-1] = lastRow[:len(lastRow)-1]

	//TODO: Hardcode, rewrite
	out.Write([]byte("\tfmt.Print(`"))

	//Writing to file
	for _, row := range req.Data {
		inserted, err := out.Write(row)
		if inserted == 0 || err != nil {
			return generr.Err("Failed to write to file", err)
		}
	}

	//TODO: Hardcode, rewrite
	out.Write([]byte("`)\n"))

	req.Reset()
	return nil
}
func (req *Request) Reset() {
	tmp := NewRequest()
	*req = tmp
}
func (req *Request) Stop() {
	//Preparing request to be written to file
	req.deleteLastRow()
	req.setNameFromOptions()

	//Creating data file and data folder if not exists yet
	createDataFolder()
	datafile := createDataFile(req.Name)

	//Flush data to data file
	//There should be a way to do this,
	//	not calling Write on every row
	for _, row := range req.Data {
		datafile.Write(row)
	}
	datafile.Close()
	req.Reset()
}
func (req *Request) ReadDataFile() *os.File {

	// Search for data file with name as in Request
	// /tmp/snapshot.data
	datafile, err := os.Open(viper.GetString("tmpdir") + "/" + req.Name + ".data")
	if err != nil {
		log.Fatal(generr.Err("Failed to open data file: "+viper.GetString("tmpdir")+"/"+req.Name+".data", err))
	}
	defer datafile.Close()

	//TODO: bytes.Buffer
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
