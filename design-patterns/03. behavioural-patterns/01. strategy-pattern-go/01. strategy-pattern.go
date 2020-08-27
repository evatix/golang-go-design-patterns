package main

import (
	"fmt"
	. "github.com/evatix/golang-go-design-patterns/strategy/internals"
)

func main() {
	chicken := NewChicken()
	fmt.Println(chicken.Display())
}
