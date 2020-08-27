package strategyinternals

import (
	"fmt"
	. "github.com/evatix/golang-go-design-patterns/strategy/internals/implementations"
)

type Chicken struct {
	Bird
}

func (c Chicken) Display() string {
	return fmt.Sprintf("I am Chicken!Bahves are:\n%s", c.BaseDisplay())
}

func NewChicken() *Chicken {
	newBird := NewBird(NoFlyBehaviour{}, SpeakBehaviour{}, NoSwimBehaviour{})
	return &Chicken{*newBird}
}
