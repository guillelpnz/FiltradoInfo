package main

import (
	// "fmt"

	"log"
	"strconv"
	"time"

	"github.com/guillelpnz/TextAnalyzer/src/micro"

	"github.com/gin-gonic/gin"
)

var textos *micro.Textos

// curl -d "contenido=mi nombre es Guillermo y esto es una prueba" -d "autor=Guillermo Lupiáñez Tapia" http://localhost:8080/texto
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
	i := c.Param("id")

	pos, _ := strconv.Atoi(i)

	c.JSON(200, gin.H{
		"mensaje":              "Texto analizado con éxito",
		"palabras redundantes": textos.ObtenerRedundantes(pos),
	})
}

// curl http://localhost:8080/obtener-sin-redundantes?posicion=0
func obtenerSinRedundantes(c *gin.Context) {
	i := c.Param("id")

	pos, _ := strconv.Atoi(i)

	c.JSON(200, gin.H{
		"mensaje":               "Texto analizado con éxito",
		"Texto sin redundantes": textos.ObtenerSinRedundantes(pos),
	})
}

// curl http://localhost:8080/texto/:id/:redundantes
func obtenerPersonas(c *gin.Context) {
	i := c.Param("id")

	pos, _ := strconv.Atoi(i)

	c.JSON(200, gin.H{
		"mensaje":  "Texto analizado con éxito",
		"Personas": textos.ObtenerPersonas(pos),
	})
}

// curl http://localhost:8080/obtener-estadisticas?posicion=0
func obtenerEstadisticas(c *gin.Context) {
	i := c.Param("id")

	pos, _ := strconv.Atoi(i)

	c.JSON(200, gin.H{
		"mensaje":      "Texto analizado con éxito",
		"Estadísticas": textos.ObtenerEstadisticas(pos),
	})
}

// go get -u github.com/gin-gonic/gin

func setupServer() *gin.Engine {
	textos = micro.NewTextos()
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	//router := gin.Default()

	router.Use(Logger())

	router.POST("/texto", introducirTexto)

	router.GET("/texto/:id/redundantes", obtenerRedundantes)

	router.GET("/texto/:id/sin-redundantes", obtenerSinRedundantes)
	router.GET("/texto/:id/personas", obtenerPersonas)
	router.GET("/texto/:id/estadisticas", obtenerEstadisticas)

	return router
}

// Logger is a custom Middleware to log TextAnalyzer api
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		path := c.Request.URL.Path

		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()
		method := c.Request.Method
		errors := c.Errors.String()

		log.Printf("Latencia: %s, Código: %v, Path: %s", latency, status, path)
		log.Printf("Tipo de petición: %s, Errores: %s", method, errors)
	}
}

func main() {
	setupServer().Run() // listen and serve on 0.0.0.0:8080
}
