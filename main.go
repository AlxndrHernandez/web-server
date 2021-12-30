package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Usuarios struct {
	id            int    `json:"id"`
	nombre        string `json:"nombre"`
	apellido      string `json:"apellido"`
	email         string `json:"email"`
	edad          int    `json:"edad"`
	altura        int    `json:"altura"`
	activo        bool   `json:"activo"`
	fechaCreacion string `json:"fechaCreacion"`
}

func main() {

	//-------- ejercicio 2
	// Crea router con gin
	router := gin.Default()

	// captura la solicitud get "/hola"
	router.GET("/saludar : nombre", func(c *gin.Context) {

		nombre := c.Param("nombre")
		c.JSON(200, gin.H{

			"message": "hola " + nombre,
		})

	})

	//--------- ejercicio 3

	datosJson, err := ioutil.ReadFile("users.json")
	if err != nil {

		log.Fatal(err)
	}

	usuarios := Usuarios{}
	err = json.Unmarshal(datosJson, &usuarios)
	if err != nil {

		log.Fatal(err)
	}

	router.GET("/usuarios", func(c *gin.Context) {

		c.JSON(200, usuarios)

	})

	router.Run() // corremos nuestro servidor sobre el puerto 8080

}
