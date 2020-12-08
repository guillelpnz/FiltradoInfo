package main

import (
	// "fmt"

	"github.com/guillelpnz/TextAnalyzer/src/micro"

	"github.com/gin-gonic/gin"
)

var textos *micro.Textos

func introducirTexto(c *gin.Context) {
	contenido := c.PostForm("contenido")
	autor := c.PostForm("autor")

	textos.IntroducirTexto(contenido, autor)

	c.JSON(200, gin.H{
		"mensaje":   "Texto añadido con éxito",
		"contenido": contenido,
		"autor":     autor,
	})
}

// go get -u github.com/gin-gonic/gin
func main() {
	textos = micro.NewTextos()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/introducir-texto", introducirTexto)
	router.Run() // listen and serve on 0.0.0.0:8080
}
