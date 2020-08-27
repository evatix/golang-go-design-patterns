package strategyimplementations

type NoFlyBehaviour struct{}

func (_ NoFlyBehaviour) Fly() string {
	return "No Flying"
}
