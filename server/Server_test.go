package server

import "testing"

/*

birds est un slice et il faudrait un map pour ce test.

func TestRegisterNewBird(t *testing.T) {

	var TestBird0 Bird

	TestBird0.ID = "zero"
	TestBird0.BirdName = "TestBirdZ"

	RegisterNewBird(TestBird0)

	if birds[len(birds)-1] != TestBird0 {
		t.Errorf("RegisterNewBird n'a pas correctement enregistré TestBird0.\n")
	}

	for i, v := range birds {
		if v == TestBird0 {
			delete(birds, i)
		}
	}

}*/

func TestGetNumberOfOneBird(t *testing.T) {

	a := GetNumberOfOneBird("tesTBIrdA") // mélange de majuscules et minuscules pour tester la resistance à la casse
	b := GetNumberOfOneBird("teStbiRdB")

	if a != 0 {
		t.Errorf("GetNumberOfOneBird renvoie testbirdA = %d au lieu de 0.\n", a)
	}

	if b != 1 {
		t.Errorf("GetNumberOfOneBird renvoie testbirdB = %d au lieu de 1.\n", b)
	}

}

func TestGetNumberOfAllBirds(t *testing.T) {

	c := GetNumberOfAllBirds()

	if c["testbirda"] != 0 || c["testbirdb"] != 1 {
		t.Errorf("GetNumberOfAllBirds ne renvoie pas le map attendue.\n")
	}

}
