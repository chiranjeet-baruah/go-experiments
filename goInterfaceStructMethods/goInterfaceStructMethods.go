package goInterfaceStructMethods

import "fmt"

type Animaler interface {
	Eat()
	Move()
	Speak()
	Error()
}

type SuperAnimals struct {
	locomotion string
}

type Animals struct {
	SuperAnimals // inheritance
	food         string
	sound        string
}

func (x Animals) Eat() {
	fmt.Println(x.food)
}

func (x Animals) Move() {
	fmt.Println(x.locomotion)
}

func (x Animals) Speak() {
	fmt.Println(x.sound)
}

func (x Animals) Error() {
	fmt.Println("Invalid query entered!")
}

func RunGoInterfaceStructMethodsExample() {
	m := map[string]Animals{
		"cow":   {SuperAnimals{"walk"}, "grass", "moo"},
		"Cow":   {SuperAnimals{"walk"}, "grass", "moo"},
		"Bird":  {SuperAnimals{"fly"}, "worms", "peep"},
		"bird":  {SuperAnimals{"fly"}, "worms", "peep"},
		"Snake": {SuperAnimals{"slither"}, "mice", "hsss"},
		"snake": {SuperAnimals{"slither"}, "mice", "hsss"},
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Enter animal name & query (eat / move / speak): ")
		fmt.Print(">")
		var animal, op string
		fmt.Scan(&animal)
		fmt.Print(">")
		fmt.Scan(&op)
		if op == "eat" {
			m[animal].Eat()
		} else if op == "move" {
			m[animal].Move()
		} else if op == "speak" {
			m[animal].Speak()
		} else {
			m[animal].Error()
		}
	}
}
