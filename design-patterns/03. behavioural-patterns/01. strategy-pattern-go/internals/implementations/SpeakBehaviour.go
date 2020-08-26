package strategyImplementations

import (
	"fmt"
)

type SpeakBehaviour struct{}

func (nf *SpeakBehaviour) Speak() {
	fmt.Println("Speaking Default.")
}
