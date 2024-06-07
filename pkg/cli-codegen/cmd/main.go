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

	originalFile, replaceFile := wrapp.GetFiles(core.ORIGINAL_FILE_DEFAULT_PATH + core.ORIGINAL_FILE_DEFAULT_NAME)
	defer originalFile.Close()
	defer replaceFile.Close()
	sess := core.NewSession()
	sess.InitOptions()

	scanner := bufio.NewScanner(originalFile)
	vReq := core.NewRequest()
	for rowIdx := 0; scanner.Scan(); rowIdx++ {
		prefixStartIndex := bytes.Index(scanner.Bytes(), core.PREFIX_BYTES)
		errorRowMsg := " " + originalFile.Name() + ":" + strconv.Itoa(rowIdx)

		switch vReq.Method {
		case "start":
			vReq.ReadLine(scanner.Bytes())
		case "stop":
			vReq.Stop()
		case "insert":
			vReq.InsertInto(replaceFile)
			vReq.Reset()
		case "":
		default:
			log.Fatal(generr.Err("Unknown method: "+vReq.Method+errorRowMsg, nil))

		}
		// Write original row
		if prefixStartIndex == -1 {
			replaceFile.Write(scanner.Bytes())
			replaceFile.Write([]byte("\n"))
			continue
		}

		queryStartindex := prefixStartIndex + len(core.PREFIX)
		if len(scanner.Bytes()) <= queryStartindex {
			log.Fatal(generr.Err("Visualiser comment exists, but does not contain method "+errorRowMsg, nil))
		}
		query := scanner.Bytes()[queryStartindex:]

		OptionsStartIndex := bytes.IndexByte(query, '?')
		if OptionsStartIndex == -1 {
			vReq.Method = string(query)
			continue
		}
		vReq.Method = string(query[:OptionsStartIndex])
		tempParams := bytes.Split(query[OptionsStartIndex+1:], []byte("&"))
		for _, v := range tempParams {

			bytes := bytes.Split(v, []byte("="))
			if len(bytes) != 2 {
				log.Fatal(generr.Err("Incorrect query Option structure, must be key=value"+errorRowMsg, nil))
			}
			vReq.Query = append(vReq.Query, core.MakeOption(string(bytes[0]), string(bytes[1])))
		}

		//__TODO: parse Options if exists
		fmt.Println("DEBUG_found method ", vReq.Method)
		fmt.Println("DEBUG_found query ", vReq.Query)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(generr.Err("Failed to read file", err))
	}
	replace(originalFile, replaceFile)
}

func replace(originalFile *os.File, replaceFile *os.File) {
	os.Rename(originalFile.Name(), strings.TrimSuffix(originalFile.Name(), ".go")+"_old.go")

	os.Rename(replaceFile.Name(), strings.TrimSuffix(replaceFile.Name(), "_replace"))

}
