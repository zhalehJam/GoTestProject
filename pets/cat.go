package pets

import (
	"fmt"
	"time"
)

type Cat struct {
	Name      string
	Color     string
	Breed     string
	lastSlept time.Time
	Animal
}

func (c Cat) GiveAttention(activity string) string {
	return fmt.Sprintf("The cat is ignoring your attemtp to %s", activity)
}

func NewCat(name, color, breed string, lastSlept time.Time) Pet {
	return Cat{
		Name:      name,
		Color:     color,
		Breed:     breed,
		lastSlept: lastSlept,
		Animal:    Animal{lastAte: time.Now()},
	}
}
