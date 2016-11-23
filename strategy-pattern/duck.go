package main

import "fmt"

/*
The code below is an example of the strategy pattern is Go

*/

// Duck interface is the main interface for all duck structs
// it's constructed in order for other structs to respond
// to duck-type methods in the same way
type Duck interface {
	Display()
	Quack()
	Fly()
	Swim()
}

// FlyWithWings struct is the fly strategy for ducks that can fly
type FlyWithWings struct{}

func (f FlyWithWings) Fly() {
	fmt.Println("I'm flying!!")
}

// FlyNoWay struct is the fly strategy for ducks that cannot fly
type FlyNoWay struct{}

func (f FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}

// Squeak struct is the quack strategy for ducks that are rubber
type Squeak struct{}

func (s Squeak) Quack() {
	fmt.Println("Squeak")
}

// RealQuack struct is the quack strategy for ducks that are real
type RealQuack struct{}

func (r RealQuack) Quack() {
	fmt.Println("Quack")
}

// MuteQuack struct is the quack strategy for ducks that are fake
type MuteQuack struct{}

func (r MuteQuack) Quack() {
	fmt.Println("<< Silence >>")
}

// AbstractDuck is the struct for all ducks
type AbstractDuck struct {
}

// Swim is the only method that all ducks share
// through inheritance
func (a AbstractDuck) Swim() {
	fmt.Println("All ducks swim")
}

// MallardDuck struct is a type of duck via inheritence (which allows it
// to use the Swim) method and also implements the `Duck` interface
type MallardDuck struct {
	AbstractDuck
	RealQuack
	FlyWithWings
}

// Display displays the mallard duck
func (m MallardDuck) Display() {
	fmt.Println("I am a mallard")
}

// RubberDuck  struct is a type of duck via inheritence (which allows it
// to use the Swim) method and also implements the `Duck` interface
type RubberDuck struct {
	AbstractDuck
	Squeak
	FlyNoWay
}

// Display displays the rubber duck
func (m RubberDuck) Display() {
	fmt.Println("I am a rubber")
}

// ModelDuck struct is a type of duck via inheritence (which allows it
// to use the Swim) method and also implements the `Duck` interface
type ModelDuck struct {
	AbstractDuck
	MuteQuack
	FlyNoWay
}

// Display displays the model duck
func (m ModelDuck) Display() {
	fmt.Println("I am a model duck")
}

func main() {
	// This is a slice of ducks
	duckHouse := []Duck{
		MallardDuck{},
		RubberDuck{},
		ModelDuck{},
	}
	for idx, d := range duckHouse {
		fmt.Printf("This is duck #%d\n", idx)
		d.Display()
		d.Fly()
		d.Quack()
		d.Swim()
		fmt.Println("_____")
	}
}
