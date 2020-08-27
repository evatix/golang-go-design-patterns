package strategyimplementations

type NoSpeakBehaviour struct{}

func (_ NoSpeakBehaviour) Speak() string {
	return "No Speaking Behaviour."
}
