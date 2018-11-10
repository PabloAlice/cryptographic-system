package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(w http.ResponseWriter, r *http.Request) {
	println(r)
}

func main() {
	r := gin.Default()
	r.POST("/api/encryption", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		method := c.PostForm("method")
		key := c.PostForm("key")
		iv := c.PostForm("iv")
		println(file)
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, fmt.Sprintf("./uploads/%s", file.Filename))

		// ! TODO encrypt file

		c.JSON(http.StatusOK, gin.H{
			"fileName": file.Filename,
			"method":   method,
			"key":      key,
			"iv":       iv,
		})
	})
	r.POST("/api/decryption", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		method := c.PostForm("method")
		key := c.PostForm("key")
		iv := c.PostForm("iv")

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, fmt.Sprintf("./uploads/%s", file.Filename))

		// ! TODO decrypt file

		c.JSON(http.StatusOK, gin.H{
			"fileName": file.Filename,
			"method":   method,
			"key":      key,
			"iv":       iv,
		})
	})

	r.Static("/static", "./client/build/static")

	r.GET("/", func(c *gin.Context) {
		c.File("./client/build/index.html")
	})
	log.Println("Listening...")
	r.Run() // listen and serve on 0.0.0.0:8080

}
