package strategyimplementations

type SwimBehaviour struct{}

func (_ SwimBehaviour) Swim() string {
	return "Swimming"
}

type NoSwimBehaviour struct{}

func (_ NoSwimBehaviour) Swim() string {
	return "No Swimming"
}
