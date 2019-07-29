package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

const pagepath = "../assets/pages/"
const pathend = "_skald"

func TestPutFile(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano())

	salt := strconv.Itoa(rand.Int())

	w := new(WikiPage)
	w.Title = "test title " + salt
	w.Creator = "creator " + salt
	w.Content = "# some content \n\nand more\n"

	n := "test_doc_" + salt
	vn := pagepath + n + pathend

	fmt.Println(n)

	//add the doc to the folder
	PutDoc(n, w)

	//Check the files exist
	assert.True(t, fileExists(vn+".json"))
	assert.True(t, fileExists(vn+".md"))

	//Lets delete the files we created
	deleteFile(vn + ".json")
	deleteFile(vn + ".md")

}
func fileExists(d string) bool {
	info, err := os.Stat(d)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func deleteFile(n string) {
	// delete file
	var err = os.Remove(n)
	if err != nil {
		return
	}

	fmt.Println("==> done deleting file")
}

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
