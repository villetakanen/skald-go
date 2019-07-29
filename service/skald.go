package main

import (
	"fmt"

	// "cloud.google.com/go/firestore"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WikiPage struct {
	Title   string `json:"Title"`
	Creator string `json:"Creator"`
	Content string `json:"Content"`
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/page/:id", getPage)
	//route.PUT("/book", createBook)
	return r

}

func main() {
	r := SetupRouter()
	r.Run()
}

/**
 * Returns a wikipage with metadata
 *
 * TODO: add MVP implementation
 * TODO: add pageProvider (local files, firestore, etc)
 */
func getPage(c *gin.Context) {
	fmt.Println("getPage")

	// TODO: check mode
	// if local
	r := getFilestoreDoc(c.Param("id"))

	// if firestore
	// TODO: add firestore

	//demopage := new(WikiPage)
	//demopage.Title = "test title"
	//demopage.Creator = "test creator"
	//demopage.Content = "test content"

	c.JSON(http.StatusOK, r)
}
func getFilestoreDoc(name string) *WikiPage {
	meta, content := GetDoc(name)
	p := new(WikiPage)

	json.Unmarshal([]byte(meta), &p)

	p.Content = content

	return p
}
