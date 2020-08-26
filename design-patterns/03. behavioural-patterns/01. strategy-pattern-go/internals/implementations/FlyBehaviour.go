package strategyImplementations

import (
	"fmt"
)

type FlyBehaviour struct{}

func (f *FlyBehaviour) Fly() {
	fmt.Println("Flying")
}
