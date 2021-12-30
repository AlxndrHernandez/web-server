package main

import "github.com/gin-gonic/gin"

// Crea router con gin
func main() {

	router := gin.Default()

	// captura la solicitud get "/hola"

	router.GET("/hello-world", func(c *gin.Context) {

		c.JSON(200, gin.H{

			"message": "hola Alexander",
		})

	})

	router.Run() // corremos nuestro servidor sobre el puerto 8080

}
