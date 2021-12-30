package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Usuarios struct {
	Id            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        bool   `json:"activo"`
	FechaCreacion string `json:"fechaCreacion"`
}

func main() {

	//-------- ejercicio 2
	// Crea router con gin
	router := gin.Default()

	// captura la solicitud get "/hola"
	router.GET("/saludar", func(c *gin.Context) {

		c.JSON(200, gin.H{

			"message": "hola Alexander",
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
