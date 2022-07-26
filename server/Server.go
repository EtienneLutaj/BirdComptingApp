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

var testbirdB Bird = Bird{"7", "testbirdB", 1} //possibilité de changer l'id si besoin

var birds []Bird = []Bird{testbirdB}

func RegisterNewBird(newBird Bird) {

	for i := 0; i < newBird.NumberOfBirds; i++ {
		birds = append(birds, newBird)
	}

	fmt.Printf("L'oiseau %v vient d'être enregistré %v fois !\n", newBird.BirdName, newBird.NumberOfBirds)

}

func GetNumberOfOneBird(birdName string) int { //ajouter option lowercase

	a := 0

	for _, b := range birds {

		if strings.ToLower(b.BirdName) == strings.ToLower(birdName) {
			a++
		}

	}

	if a == 0 {
		fmt.Printf("L'oiseau %v n'existe pas ou n'a jamais été aperçu pour l'instant.\n", birdName)
	} else {
		fmt.Printf("L'oiseau %v a été aperçu %v fois !\n", birdName, a)
	}

	return a

}

func GetNumberOfAllBirds() map[string]int { //ajouter option lowercase

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
