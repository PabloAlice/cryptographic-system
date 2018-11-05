package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(w http.ResponseWriter, r *http.Request) {
	println(r)
}

func main() {
	r := gin.Default()
	r.GET("/api/encrypt", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "encrypted",
		})
	})
	r.GET("/api/decrypt", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "decrypted",
		})
	})

	r.Static("/static", "./client/build/static")

	r.GET("/", func(c *gin.Context) {
		c.File("./client/build/index.html")
	})
	log.Println("Listening...")
	r.Run() // listen and serve on 0.0.0.0:8080

}
