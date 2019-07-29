package main

import (
	"fmt"

	// "cloud.google.com/go/firestore"

	"github.com/gin-gonic/gin"
)

type BookMetaData struct {
	Title   string `json:"title"`
	Creator string `json:"creator"`
}

func main() {
	route := gin.Default()
	route.PUT("/book", createBook)
	route.Run()
}

func createBook(c *gin.Context) {

	/*/create a firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "adaptivestoryengine")
	if err != nil {
		c.AbortWithError(400, err)
		fmt.Printf("something happened on firestore.NewClient")
		fmt.Println(err)
		return
	}
	test := client.Doc("books/test")
	fmt.Println(test) */

	data := new(BookMetaData)

	errJ := c.BindJSON(data)
	if errJ != nil {
		fmt.Printf("something happened on c.BindJSON \n")
		fmt.Println(errJ)
		c.AbortWithError(400, errJ)
		return
	}

	c.String(200, fmt.Sprintf("%#v", data))

}
