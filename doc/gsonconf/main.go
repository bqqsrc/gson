//  Copyright (C) 晓白齐齐,版权所有.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/bqqsrc/gson"
)

func main() {
	confFile := ""
	keys := ""
	valType := ""
	flag.StringVar(&confFile, "conf", confFile, "gson文件")
	flag.StringVar(&keys, "keys", keys, "要取值的键，如果有多层键，则用英文冒号（:）分隔开，如果传入空值，则默认将整个gson文件转为指定的配置值")
	flag.StringVar(&valType, "type", valType, "数据类型，取值string，或number，或bool，如果没有传入值或者传入非法值，默认为string")
	flag.Parse()

	if confFile == "" {
		log.Fatal("gson文件不能为空，请使用-conf传入gson文件参数")
	}
	file, err := os.Stat(confFile)
	if err != nil {
		log.Fatalf("%s不存在或报错，请检查文件", confFile)
	}
	if file.IsDir() {
		log.Fatalf("%s是一个路径，必需传入一个gson文件", confFile)
	}
	data, _ := ioutil.ReadFile(confFile)

	if valType == "number" {
		if err, val := gson.GetFloat64WithSep(data, keys, ":"); err != nil {
			log.Fatalf("gson error, GetBoolWithSep err: %s", err)
		} else {
			fmt.Print(val)
		}
	} else if valType == "bool" {
		if err, val := gson.GetBoolWithSep(data, keys, ":"); err != nil {
			log.Fatalf("gson error, GetBoolWithSep err: %s", err)
		} else {
			fmt.Print(val)
		}
	} else {
		if err, val := gson.GetStringWithSep(data, keys, ":"); err != nil {
			log.Fatalf("gson error, GetBoolWithSep err: %s", err)
		} else {
			fmt.Print(val)
		}
	}
}
