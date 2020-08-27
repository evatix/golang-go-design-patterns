package strategyimplementations

type FlyBehaviour struct{}

func (_ FlyBehaviour) Fly() string {
	return "Flying"
}
