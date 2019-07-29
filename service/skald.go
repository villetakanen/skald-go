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
	r.PUT("/page/:id", putPage)
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

	c.JSON(http.StatusOK, r)
}

/**
 * Creates a wikipage from the PUT request
 *
 * TODO, just saves the payload for now, no sanity or even checks!
 * TODO, returns the page from disk after save, good for dev, pretty slow for prod
 */
func putPage(c *gin.Context) {

	p := new(WikiPage)

	err := c.BindJSON(p)
	if err != nil {
		fmt.Printf("something happened on c.BindJSON \n")
		fmt.Println(err)
		c.AbortWithError(400, err)
		return
	}

	putFilestoreDoc(c.Param("id"), p)

	getPage(c)
}
func putFilestoreDoc(name string, p *WikiPage) {

}

func getFilestoreDoc(name string) *WikiPage {
	meta, content := GetDoc(name)
	p := new(WikiPage)

	json.Unmarshal([]byte(meta), &p)

	p.Content = content

	return p
}
