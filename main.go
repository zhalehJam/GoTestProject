package main

import (
	"fmt"
	"main/pets"
	"time"
)

func main() {
	// pet := pets.Dog{

	// 	Name:  "Oreo",
	// 	Color: "Black",
	// 	Breed: "Mixed",
	// }
	sleepTime := time.Now() // time.now().Add(time.Duration(-5)*time.Hour)
	var animals []pets.Pet
	animals = append(animals, pets.NewDog("Odi", "White", "Mixed", sleepTime))
	animals = append(animals, pets.NewCat("pishan", "balckandWhite", "Mixed", sleepTime))

	pet := pets.NewDog("Oreo", "Black", "Mixed", sleepTime)

	if pet.IsHungry() {
		fmt.Println(pet.Feed("kibble"))
	} else {
		fmt.Println("Your animal isn't hungry")
		time.Sleep(5 * time.Second)
		fmt.Println(pet.Feed("kibble"))
	}
	fmt.Println(pet.Feed("steak"))
	fmt.Println(pet.GiveAttention("play fetch"))

}
