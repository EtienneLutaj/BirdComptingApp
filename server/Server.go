package server

import "fmt"

type Bird struct {
	ID       string `json:"id"`
	BirdName string `json:"birdname"`
	//NumberOfBirds int    `json:"numberofbirds"`
}

var birds []Bird

func RegisterNewBird(newBird Bird) {

	birds = append(birds, newBird)

	fmt.Printf("L'oiseau %v a bien été enregistré !\n", newBird.BirdName)

}

func GetNumberOfOneBird(birdName string) int { //ajouter option lowercase

	a := 0

	for _, b := range birds {

		if b.BirdName == birdName {
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

		v, ok := BirdMap[bird.BirdName]

		if ok {
			BirdMap[bird.BirdName] = v + 1
		} else {
			BirdMap[bird.BirdName] = 1
		}

	}

	return BirdMap

}
