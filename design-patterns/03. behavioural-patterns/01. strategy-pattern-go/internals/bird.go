package strategyinternals

import (
	"fmt"
	. "github.com/evatix/golang-go-design-patterns/strategy/internals/interfaces"
)

// Embedding https://golang.org/doc/effective_go.html#embedding
type Bird struct {
	flyer   Flyer
	speaker Speaker
	swimmer Swimmer
	// naming can have improper words which is fine according to the convention
	BaseDisplayer
	Flyer
	Speaker
	Swimmer
	Displayer
}

func (b Bird) BaseDisplay() string {
	flyBehave := b.Fly()
	speakBehave := b.Speak()
	swimBehave := b.Swim()

	return fmt.Sprintf(flyBehave, speakBehave, swimBehave)
}

func (b Bird) Fly() string {
	return b.flyer.Fly()
}

func (b Bird) Swim() string {
	return b.swimmer.Swim()
}

func (b Bird) Speak() string {
	return b.speaker.Speak()
}

func NewBird(flyer Flyer, speaker Speaker, swimmer Swimmer) *Bird {
	return &Bird{flyer: flyer, speaker: speaker, swimmer: swimmer}
}
