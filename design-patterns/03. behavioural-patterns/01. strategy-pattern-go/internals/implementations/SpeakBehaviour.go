package strategyimplementations

type SpeakBehaviour struct{}

func (_ SpeakBehaviour) Speak() string {
	return "Speaking Default."
}
