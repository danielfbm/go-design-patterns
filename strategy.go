package main

import "fmt"

// First lets define our behavior

type WalkBehavior interface {
	Walk()
}

type JumpBehavior interface {
	Jump()
}

type Human struct {
	WalkBehavior
	JumpBehavior
}

// --- Normal

type NormalWalk struct{}

func (NormalWalk) Walk() {
	fmt.Println("This is a normal walk")
}

type NormalJump struct{}

func (NormalJump) Jump() {
	fmt.Println("This is a normal jump!")
}

// MoonWalk

type MoonWalk struct{}

func (MoonWalk) Walk() {
	fmt.Println("Moon walk like Michael Jackson!!!!")
}

type RabbitJump struct{}

func (RabbitJump) Jump() {
	fmt.Println("I am just jumping like a rabbit!")
}

// --- starting

var walker Human

func init() {
	walker = Human{
		WalkBehavior: NormalWalk{},
		JumpBehavior: NormalJump{},
	}
}

func main() {

	fmt.Println("This is a normal human, Joe Doe..")
	fmt.Println("How does Joe walks?")
	walker.Walk()

	fmt.Println("How does Joe jump?")
	walker.Jump()

	fmt.Println("-----")
	fmt.Println("Joe enters in Michael Jackson Mode")
	walker.JumpBehavior = RabbitJump{}
	walker.WalkBehavior = MoonWalk{}

	fmt.Println("How does Joe walks?")
	walker.Walk()

	fmt.Println("How does Joe jump?")
	walker.Jump()
}
