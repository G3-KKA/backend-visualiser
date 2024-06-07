package wrapp

import (
	generr "backend-visualiser/cli-codegen/internal/errors/codegenError"
	"log"
	"os"
)

func GetFiles(originalPath string) (original *os.File, replace *os.File) {
	log.Println("Need new logger:", "func Files(originalPath string) (original *os.File, replace *os.File)")
	originalFile, err := os.Open(originalPath)
	if err != nil {
		log.Fatal(generr.Err("Failed to open original file: "+originalPath, err))
	}
	replacePath := originalFile.Name() + "_replace"
	replaceFile, err := os.Create(replacePath)
	if err != nil {
		log.Fatal(generr.Err("Failed to create replace file: "+replacePath, err))
	}
	return originalFile, replaceFile
}
