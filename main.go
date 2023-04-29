package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const addForm = `
<html><body>
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
</html></body>
`

var store = NewURLStore()

func main() {
	router := gin.Default()

	// handleFunc for "/"  GET
	router.GET("/:key", Redirect)
	router.GET("/", func(c *gin.Context) {
		c.Writer.WriteString(addForm)
	})

	// handle Func for "/add" POST
	router.POST("/add", Add)

	// Start the server on 8080
	router.Run(":8080")
}
func Add(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	url := c.PostForm("url")
	if url == "" {
		c.Writer.WriteString(addForm)
		return
	}
	key := store.Put(url)

	c.Writer.WriteString(fmt.Sprintf("%s", key))
}
func Redirect(c *gin.Context) {
	key := c.Param("key")
	url := store.Get(key)
	if url == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Redirect(http.StatusFound, url)
}
