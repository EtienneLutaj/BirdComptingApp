package server

import (
	"fmt"
	"strings"
)

type Bird struct {
	ID            string `json:"id"`
	BirdName      string `json:"birdname"`
	NumberOfBirds int    `json:"numberofbirds"`
}

var testbirdB Bird = Bird{"7", "testbirdB", 1} //possibilité de changer l'id si besoin //le B est en majuscule pour tester la resistance à la casse de GetNumberOfAllBirds
var birds []Bird = []Bird{testbirdB}           //nécessaire pour les tests (voir Server_test.go)

func RegisterNewBird(newBird Bird) {

	for i := 0; i < newBird.NumberOfBirds; i++ {
		birds = append(birds, newBird)
	}

	fmt.Printf("L'oiseau %v vient d'être enregistré %v fois !\n", newBird.BirdName, newBird.NumberOfBirds)

}

func GetNumberOfOneBird(birdName string) int {

	totalCountOfThisBird := 0

	for _, bird := range birds {

		if strings.ToLower(bird.BirdName) == strings.ToLower(birdName) {
			totalCountOfThisBird++
		}

	}

	if totalCountOfThisBird == 0 {
		fmt.Printf("L'oiseau %v n'existe pas ou n'a jamais été aperçu pour l'instant.\n", birdName)
	} else {
		fmt.Printf("L'oiseau %v a été aperçu %v fois !\n", birdName, totalCountOfThisBird)
	}

	return totalCountOfThisBird

}

func GetNumberOfAllBirds() map[string]int {

	BirdMap := make(map[string]int)

	for _, bird := range birds {

		v, ok := BirdMap[strings.ToLower(bird.BirdName)]

		if ok {
			BirdMap[strings.ToLower(bird.BirdName)] = v + 1
		} else {
			BirdMap[strings.ToLower(bird.BirdName)] = 1
		}

	}

	return BirdMap

}
