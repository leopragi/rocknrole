package main

import (
	"log"

	"roles/platform/action"
)

// import "github.com/gin-gonic/gin"

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run()
// }

func main() {
	data := map[string]string{"name": "Hi"}
	action := action.MakeAction(data)
	log.Print(action)
	action = action.Rename("Hello")
	log.Print(action)
}
