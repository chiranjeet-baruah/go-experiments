package goInterfaceStructMethods

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) update(n string, a int) {
	p.name = n
	p.age = a
	fmt.Println(p.name, p.age)
}

func display(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

func RunGoMethodsExample() {
	var p Person
	p.update("Chiranjeet", 28)

	mess := display("Hello!")
	fmt.Println(mess)
}
