package main

import (
	"io/ioutil"
	"os"

	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDefaultDoc(t *testing.T) {
	// Grab the files
	vjsonf, err := os.Open("../assets/default/index_skald.json")
	if err != nil {
		log.Fatal(err)
	}
	defer vjsonf.Close()

	vmdf, err := os.Open("../assets/default/index_skald.md")
	if err != nil {
		log.Fatal(err)
	}
	defer vmdf.Close()

	vjson, _ := ioutil.ReadAll(vjsonf)
	vmd, _ := ioutil.ReadAll(vmdf)
	tjson, tmd := GetDoc("index")

	assert.Equal(t, string(vjson), tjson)
	assert.Equal(t, string(vmd), tmd)
}

func TestGetUpdatedDoc(t *testing.T) {
	// Grab the files
	vjsonf, err := os.Open("../assets/pages/1_skald.json")
	if err != nil {
		log.Fatal(err)
	}
	defer vjsonf.Close()

	vmdf, err := os.Open("../assets/pages/1_skald.md")
	if err != nil {
		log.Fatal(err)
	}
	defer vmdf.Close()

	vjson, _ := ioutil.ReadAll(vjsonf)
	vmd, _ := ioutil.ReadAll(vmdf)
	tjson, tmd := GetDoc("1")

	assert.Equal(t, string(vjson), tjson)
	assert.Equal(t, string(vmd), tmd)
}
