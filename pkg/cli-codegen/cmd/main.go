package main

import (
	"backend-visualiser/cli-codegen/internal/config"
	"backend-visualiser/cli-codegen/internal/core"
	generr "backend-visualiser/cli-codegen/internal/errors/codegenError"
	"backend-visualiser/cli-codegen/internal/errors/current"
	"backend-visualiser/cli-codegen/internal/lib/wrapp"
	"backend-visualiser/cli-codegen/internal/logger"
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	logger := logger.InitLogger()
	logger.Info("cli-codegen")
	original, replace := wrapp.GetFiles(viper.GetString("file"))
	defer original.Close()
	defer replace.Close()
	sess := core.NewSession()
	sess.InitOptions()
	/* builder := strings.Builder{} */
	scanner := bufio.NewScanner(original)
	req := core.NewRequest()
	logger.Debug("Preparations succeeded")
	logger.Sync()
	defer logger.Sync()
	panic("everything before this succeded")
	for rowIdx := 0; scanner.Scan(); rowIdx++ {
		startIndex := bytes.Index(scanner.Bytes(), core.PrefixBytes())
		current.Phase = " " + original.Name() + ":" + strconv.Itoa(rowIdx)

		switch req.Method {
		case "start":
			req.ReadLine(scanner.Bytes())
		case "stop":
			req.Stop()
		case "insert":
			req.InsertInto(replace)
		case "":
		default:
			log.Fatal(generr.Err("Unknown method: "+req.Method+current.Phase, nil))
		}
		// Write original rows, ignore rows with visualiser comment
		if startIndex == -1 {
			replace.Write(scanner.Bytes())
			replace.Write([]byte("\n"))
			continue
		}

		queryindex := startIndex + len(core.PREFIX)
		if len(scanner.Bytes()) <= queryindex {
			log.Fatal(generr.Err("Missing method at: "+current.Phase, nil))
		}
		//__TODO: wrap to query logic
		query := scanner.Bytes()[queryindex:]
		optionsIndex := bytes.IndexByte(query, '?')
		if optionsIndex == -1 {
			req.Method = string(query)
			continue
		}
		req.Method = string(query[:optionsIndex])
		options := bytes.Split(query[optionsIndex+1:], []byte("&"))
		for _, optRaw := range options {
			opt := core.MakeOptionFromRaw(optRaw)
			req.Query = append(req.Query, opt)
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
