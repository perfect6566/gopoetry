package main

import (
	"flag"
	"github.com/perfect6566/gopoetry/cmd"
	"log"
)


func main()  {

	poetrypath:=flag.String("p","","请输入字符串文件路径,默认为data/poetry")
	keyword:=flag.String("k","","请输入关键字匹配,默认为敏")
	flag.Parse()

	if *poetrypath==""{
		log.Println("请输入诗歌库路径 -p path")
	}
	if *keyword==""{
		log.Println("请输入搜索的关键字 -k keyword")
	}


p:=cmd.NewProcesser(*poetrypath,*keyword)

p.Printpoetry()

}
