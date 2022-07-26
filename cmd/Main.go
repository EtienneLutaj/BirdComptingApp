package main

import (
	"fmt"
	"net/http"

	"github.com/EtienneLutaj/BirdComptingApp/server"

	"flag"

	"github.com/gin-gonic/gin"
)

//make(chan int) ?

func PostBird(context *gin.Context) {

	var newBird server.Bird

	if err := context.BindJSON(&newBird); err != nil {
		fmt.Print(err)
		return
	}

	context.IndentedJSON(http.StatusCreated, newBird)

	server.RegisterNewBird(newBird)

}

func GetBird(context *gin.Context) {
	birdName := context.Param("birdname")
	number := server.GetNumberOfOneBird(birdName)

	context.IndentedJSON(http.StatusOK, number)
}

func GetBirds(context *gin.Context) {
	BirdMap := server.GetNumberOfAllBirds()
	context.IndentedJSON(http.StatusOK, BirdMap)
}

func main() {

	port := flag.String("port", "8080", "port paramétrable")

	flag.Parse()

	router := gin.Default()
	router.POST("/bird", PostBird) // ne prend les JSON que 1 par 1 et ne prend pas de liste avec plusieurs JSON en entré.
	router.GET("/bird/:birdname", GetBird)
	router.GET("/bird", GetBirds)
	router.Run("localhost:" + *port)

}
