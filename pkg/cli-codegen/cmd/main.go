package main

import (
	"backend-visualiser/cli-codegen/internal/core"
	generr "backend-visualiser/cli-codegen/internal/errors/codegenError"
	"backend-visualiser/cli-codegen/internal/lib/wrapp"
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
__TODO:





*/

func main() {

	original, replace := wrapp.GetFiles(core.ORIGINAL_FILE_DEFAULT_PATH + core.ORIGINAL_FILE_DEFAULT_NAME)
	defer original.Close()
	defer replace.Close()
	sess := core.NewSession()
	sess.InitOptions()

	scanner := bufio.NewScanner(original)
	req := core.NewRequest()
	for rowIdx := 0; scanner.Scan(); rowIdx++ {
		prefixStartIndex := bytes.Index(scanner.Bytes(), core.PREFIX_BYTES)
		errorRowMsg := " " + original.Name() + ":" + strconv.Itoa(rowIdx)

		switch req.Method {
		case "start":
			req.ReadLine(scanner.Bytes())
		case "stop":
			req.Stop()
		case "insert":
			req.InsertInto(replace)
			req.Reset()
		case "":
		default:
			log.Fatal(generr.Err("Unknown method: "+req.Method+errorRowMsg, nil))

		}
		// Write original rows, ignore rows with visualiser comment
		if prefixStartIndex == -1 {
			replace.Write(scanner.Bytes())
			replace.Write([]byte("\n"))
			continue
		}

		queryStartindex := prefixStartIndex + len(core.PREFIX)
		if len(scanner.Bytes()) <= queryStartindex {
			log.Fatal(generr.Err("Missing method at: "+errorRowMsg, nil))
		}
		query := scanner.Bytes()[queryStartindex:]
		OptionsStartIndex := bytes.IndexByte(query, '?')
		if OptionsStartIndex == -1 {
			req.Method = string(query)
			continue
		}
		req.Method = string(query[:OptionsStartIndex])
		tempParams := bytes.Split(query[OptionsStartIndex+1:], []byte("&"))
		for _, v := range tempParams {

			bytes := bytes.Split(v, []byte("="))
			if len(bytes) != 2 {
				log.Fatal(generr.Err("Wrong query options, must be key=value"+errorRowMsg, nil))
			}
			req.Query = append(req.Query, core.MakeOption(string(bytes[0]), string(bytes[1])))
		}

		//__TODO: parse Options if exists
		fmt.Println("DEBUG_found method ", req.Method)
		fmt.Println("DEBUG_found query ", req.Query)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(generr.Err("Failed to read file", err))
	}
	swapFiles(original, replace)
}

func swapFiles(original *os.File, replace *os.File) {
	os.Rename(original.Name(), strings.TrimSuffix(original.Name(), ".go")+"_old.go")

	os.Rename(replace.Name(), strings.TrimSuffix(replace.Name(), "_replace"))

}
