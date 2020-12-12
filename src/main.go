package main

import (
	// "fmt"

	"strconv"

	"github.com/guillelpnz/TextAnalyzer/src/micro"

	"github.com/gin-gonic/gin"
)

var textos *micro.Textos

// curl -d "contenido=mi nombre es Guillermo y esto es una prueba" -d "autor=Guillermo Lupiáñez Tapia" http://localhost:8080/introducir-texto
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

// curl http://localhost:8080/obtener-redundantes?posicion=0
func obtenerRedundantes(c *gin.Context) {
	i := c.Query("posicion")

	pos, _ := strconv.Atoi(i)

	c.JSON(200, gin.H{
		"mensaje":              "Texto analizado con éxito",
		"palabras redundantes": textos.ObtenerRedundantes(pos),
	})
}

// curl http://localhost:8080/obtener-sin-redundantes?posicion=0
func obtenerSinRedundantes(c *gin.Context) {
	i := c.Query("posicion")

	pos, _ := strconv.Atoi(i)

	c.JSON(200, gin.H{
		"mensaje":               "Texto analizado con éxito",
		"Texto sin redundantes": textos.ObtenerSinRedundantes(pos),
	})
}

// curl http://localhost:8080/obtener-personas?posicion=0
func obtenerPersonas(c *gin.Context) {
	i := c.Query("posicion")

	pos, _ := strconv.Atoi(i)

	c.JSON(200, gin.H{
		"mensaje":  "Texto analizado con éxito",
		"Personas": textos.ObtenerPersonas(pos),
	})
}

// curl http://localhost:8080/obtener-estadisticas?posicion=0
func obtenerEstadisticas(c *gin.Context) {
	i := c.Query("posicion")

	pos, _ := strconv.Atoi(i)

	c.JSON(200, gin.H{
		"mensaje":      "Texto analizado con éxito",
		"Estadísticas": textos.ObtenerEstadisticas(pos),
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

	router.GET("/obtener-redundantes", obtenerRedundantes)

	router.GET("/obtener-sin-redundantes", obtenerSinRedundantes)
	router.GET("/obtener-personas", obtenerPersonas)
	router.GET("/obtener-estadisticas", obtenerEstadisticas)
	router.Run() // listen and serve on 0.0.0.0:8080
}
