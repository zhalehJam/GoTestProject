package pets

import (
	"fmt"
	"time"
)

type Animal struct {
	lastAte time.Time
}

func (a Animal) Feed(food string) string {
	a.lastAte = time.Now()
	return fmt.Sprintf("the animal is eating %s", food)
}
func (a Animal) IsHungry() bool {
	return time.Now().Sub(a.lastAte) > 2*time.Hour
}
