package strategyImplementations

import (
	"fmt"
)

type NoFlyBehaviour struct{}

func (nf *NoFlyBehaviour) Fly() {
	fmt.Println("No Flying")
}
