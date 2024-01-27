package pets

import (
	"fmt"
	"strings"
	"time"
)

type Dog struct {
	Name      string
	Color     string
	Breed     string
	lastSlept time.Time
	Animal
}

func (d Dog) needsSleep() bool {
	return time.Now().Sub(d.lastSlept) > 4*time.Hour
}
func (d Dog) sleep() {
	d.lastSlept = time.Now()
}

func (d Dog) Feed(food string) string {
	return fmt.Sprintf("%s eating %s", d.Name, food)
}

func (d Dog) GiveAttention(activity string) string {
	response := ""
	if d.needsSleep() {
		d.sleep()
		return "Your dog is tired and need to rest"
	}

	switch strings.ToUpper(activity) {
	case "PET":
		response = "wags tail"
	case "Playing Fetch":
		response = "Return the ball and jump waiting for you to throw it again"
	default:
		response = "Barks"
	}
	return fmt.Sprintf("%s loves attention , %s will cause hin to %s", d.Name, activity, response)
}

func NewDog(name, color, breed string, lastSleep time.Time) Pet { //Dog {
	return Dog{
		Name:      name,
		Color:     color,
		Breed:     breed,
		lastSlept: lastSleep,
		Animal:    Animal{lastAte: time.Now()},
	}
}
