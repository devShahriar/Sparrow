package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	go h.run()

	router := gin.New()
	router.LoadHTMLFiles("index.html")
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})

	router.Run(":8080")
}
