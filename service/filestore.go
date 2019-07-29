package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

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

func docExists(d string) bool {
	info, err := os.Stat(d)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
