//  Copyright (C) 晓白齐齐,版权所有.

package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	typeArr := []string{
		"Int",
		"Int8",
		"Int16",
		"Int32",
		"Int64",
		"Uint",
		"Uint8",
		"Uint16",
		"Uint32",
		"Uint64",
		"Float32",
		"Float64",
		"String",
		"Bool",
	}
	tarPath := ""
	srcPath := "../opt.go"
	codeStr, _ := ioutil.ReadFile(srcPath)
	codeRet := ""
	for _, typeStr := range typeArr {
		typeStrT := strings.ToLower(typeStr)
		codeRet = strings.Replace(strings.Replace(string(codeStr), "$Opt", typeStr, -1), "$opt", typeStrT, -1)
		tarPath = "../../g" + typeStrT + ".go"
		ioutil.WriteFile(tarPath, []byte(codeRet), 0644)
	}
}
