package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/Joker/hpp"
	"github.com/Joker/jade"
)

func main() {
	srcPath := "./index.pug"
	destPath := "."

	fInfo, err := os.Stat(srcPath)
	if err != nil {
		log.Fatalf("parse error: %v", err)
	}
	pathSep := string(os.PathSeparator)
	fileName := fileNameWithoutExt(fInfo.Name())
	if err != nil {
		log.Fatalf("parse error: %v", err)
	}
	destPath += pathSep + fileName + ".html"

	htmlTpl, err := jade.ParseFile(srcPath)
	if err != nil {
		log.Fatalf("parse error: %v", err)
	}

	hppTpl := hpp.PrPrint(htmlTpl)

	err = ioutil.WriteFile(destPath, []byte(hppTpl), fInfo.Mode())
	if err != nil {
		log.Fatalf("parse error: %v", err)
	}
}

func fileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
