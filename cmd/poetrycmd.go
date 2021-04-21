package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"path/filepath"
)

type Portey struct {
	Subject string
	Dynasty string
	Author  string
	Content string
}

func (PP *Portey) String() string {
	return PP.Subject + PP.Content + PP.Author + PP.Dynasty
}

type Processer struct {
	Filespath     string
	ResultPorteys []*Portey
	Keyname       string
}

func NewProcesser(filespath, keyname string) *Processer {
	return &Processer{Filespath: filespath, Keyname: keyname}
}
func (p *Processer) Getfileslist() []string {

	if !strings.HasSuffix(p.Filespath, "/") {
		p.Filespath = p.Filespath + "/"
	}

	paths, err := filepath.Glob(p.Filespath + "*.csv")
	if err != nil {
		log.Println(err)
	}

	return paths
}

func (p *Processer) Analysisfile(filename string) error {

	r, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return err
	}
	subcontents := strings.Split(string(r), "\n")
	for _, v := range subcontents {

		if strings.Contains(v, p.Keyname) {
			//log.Println(" matched",string(v))
			subject := strings.Split(v, ",")[0]
			dynasty := strings.Split(v, ",")[1]
			author := strings.Split(v, ",")[2]
			content := strings.Split(v, ",")[3]
			Matchedportey := &Portey{Subject: subject, Dynasty: dynasty, Author: author, Content: content}

			p.ResultPorteys = append(p.ResultPorteys, Matchedportey)
		}
	}
	return nil
}

func (p *Processer) Printpoetry() {
	paths := p.Getfileslist()
	if len(paths) == 0 {
		log.Println("诗歌数据文件未找到，程序退出")

	}
	for _, targetfilename := range paths {
		err := p.Analysisfile(targetfilename)
		if err != nil {
			log.Println(targetfilename, err)
		}

	}

	for _, result := range p.ResultPorteys {
		fmt.Println("名称： " + result.Subject + " 朝代：" + result.Dynasty +
			"  作者：" + result.Author + " 内容: " + result.Content)
	}

}
