package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//Server_test
	//Server
)

type Bird struct {
	ID       string `json:"id"`
	BirdName string `json:"birdname"`
	//NumberOfBirds int    `json:"numberofbirds"`
}

var birds []Bird

//make(chan int) ?

func PostBird(context *gin.Context) {

	var newBird Bird

	if err := context.BindJSON(&newBird); err != nil {
		fmt.Print(err)
		return
	}

	birds = append(birds, newBird)

	context.IndentedJSON(http.StatusCreated, newBird)

	fmt.Printf("L'oiseau %v", newBird.BirdName)

}

func GetBird(context *gin.Context) {
	birdName := context.Param("birdname")
	number := GetNumberOfOneBird(birdName)

	context.IndentedJSON(http.StatusOK, number)
}

func GetBirds(context *gin.Context) {
	BirdMap := GetNumberOfAllBirds()
	context.IndentedJSON(http.StatusOK, BirdMap)
}

func GetNumberOfOneBird(birdName string) int {

	a := 0

	for _, b := range birds {

		if b.BirdName == birdName {
			a++
		}

	}

	if a == 0 {
		fmt.Printf("L'oiseau %v n'existe pas ou n'a jamais été aperçu pour l'instant", birdName)
	} else {
		fmt.Printf("L'oiseau %v a été aperçu %v fois !", birdName, a)
	}

	return a

}

func GetNumberOfAllBirds() map[string]int {

	BirdMap := make(map[string]int)

	for _, bird := range birds {

		v, ok := BirdMap[bird.BirdName]

		if ok {
			BirdMap[bird.BirdName] = v + 1
		} else {
			BirdMap[bird.BirdName] = 1
		}

	}

	return BirdMap

}

func main() {
	router := gin.Default()
	fmt.Print("miaou miaou")
	router.POST("/bird", PostBird)
	router.GET("/bird", GetBird)
	router.GET("/bird", GetBirds)
	router.Run("localhost:8080")
}
