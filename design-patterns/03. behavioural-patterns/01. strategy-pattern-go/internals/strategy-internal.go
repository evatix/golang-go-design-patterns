package strategyInternals

import . "github.com/evatix/golang-go-design-patterns/strategy/internals/interfaces"

type Bird struct {
	flyer   Flyer
	speaker Speaker
	swimmer Swimmer
	Flyer
}

func (b Bird) Fly() {
	b.flyer.Fly()
}

func (b Bird) Swim() {
	b.swimmer.Swim()
}

func (b Bird) Speak() {
	b.speaker.Speak()
}

func NewBird(flyer Flyer, speaker Speaker, swimmer Swimmer) *Bird {
	return &Bird{flyer: flyer, speaker: speaker, swimmer: swimmer}
}
