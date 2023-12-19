//  Copyright (C) 晓白齐齐,版权所有.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bqqsrc/bqqg/file"
	"github.com/bqqsrc/gson"
)

func main() {
	gsonFile := ""
	args := ""
	flag.StringVar(&gsonFile, "gson", gsonFile, "gson配置文件路径，包括文件名")
	flag.StringVar(&args, "arg", args, "要获取的参数，如果是多级参数，用英文冒号连接，如arg1:arg2:arg3，获取arg1下的arg2下的arg3")
	flag.Parse()
	err := false
	if args == "" {
		fmt.Println("没有传入参数args，使用--help查看使用说明")
		err = true
	}
	var data []byte
	var fileExist bool
	if gsonFile == "" {
		fmt.Println("没有传入文件gson，使用--help查看使用说明")
		err = true
	} else {
		if !file.Exist(gsonFile) {
			if wd, e := os.Getwd(); e != nil {
				fmt.Println(e)
				err = true
			} else {
				gsonFile = file.JoinPath(wd, gsonFile)
				if !file.Exist(gsonFile) {
					fmt.Printf("传入文件'%s'不存在\n", gsonFile)
					err = true
				} else {
					fileExist = true
				}
			}
		} else {
			fileExist = true
		}
	}

	if fileExist {
		var e error
		if data, e = ioutil.ReadFile(gsonFile); e != nil {
			fmt.Println(e)
			err = true
		}
	}

	if err {
		os.Exit(1)
	}
	originArr := strings.Split(args, ":")
	argArr := originArr[:]
	val := ""
	keyEventCallback := func(d *gson.Decoder, l *gson.Lexer, isFound bool) bool {
		key := l.String()
		if isFound {
			if argArr != nil && len(argArr) > 0 {
				if key == argArr[0] {
					argArr = argArr[1:]
				}
			}
			if len(argArr) == 0 {
				d.SetAnyTarget(&val, true)
			}
		} else {
			if len(originArr) != len(argArr) {
				argArr = originArr[len(originArr)-len(argArr)-1:]
			}
		}
		return true
	}
	if e := gson.DecodeData(data, keyEventCallback, nil, nil); e != nil {
		fmt.Printf("解析gson文件出错：%s\n", e)
		os.Exit(1)
	}

	fmt.Printf(val)
}
