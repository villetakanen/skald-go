package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type WikiPageMeta struct {
	Title   string `json:"Title"`
	Creator string `json:"Creator"`
}

func GetDoc(d string) (string, string) {
	d = d + "_skald"
	meta := ""
	md := ""

	if docExists("../assets/pages/" + d + ".json") {
		meta = getFile("../assets/pages/" + d + ".json")
		md = getFile("../assets/pages/" + d + ".md")
	} else {
		meta = getFile("../assets/default/" + d + ".json")
		md = getFile("../assets/default/" + d + ".md")
	}
	return meta, md
}

func PutDoc(n string, w *WikiPage) {
	//Write the document metadata
	m := new(WikiPageMeta)
	m.Title = w.Title
	m.Creator = w.Creator
	putMetaFile("../assets/pages/"+n+"_skald.json", *m)

	//Write the document content
	c := w.Content
	putMdFile("../assets/pages/"+n+"_skald.md", c)
}

func getFile(n string) string {
	f, err := os.Open(n)
	if err != nil {
		fmt.Print(err)
		return ""
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	return string(b)
}
func putMetaFile(n string, d WikiPageMeta) {
	f, _ := json.MarshalIndent(d, "", " ")
	err := ioutil.WriteFile(n, f, 0644)
	if err != nil {
		fmt.Print(err)
	}
}
func putMdFile(n string, d string) {
	f := []byte(d)
	err := ioutil.WriteFile(n, f, 0644)
	if err != nil {
		fmt.Print(err)
	}
}
func docExists(d string) bool {
	info, err := os.Stat(d)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
