package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
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
func TestPutGetFile(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano())

	salt := strconv.Itoa(rand.Int())

	w := new(WikiPage)
	w.Title = "test title " + salt
	w.Creator = "creator " + salt
	w.Content = "# some content \n\nand more\n"

	n := "eq_test_doc_" + salt
	vn := pagepath + n + pathend

	fmt.Println(n)

	//add the doc to the folder
	PutDoc(n, w)
	vw := GetDoc(n)

	//Check the files exist
	assert.True(t, reflect.DeepEqual(w, vw))
	//assert.Equal(t, vmd, w.Content)

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
		log.Fatal(err)
		return
	}
}
