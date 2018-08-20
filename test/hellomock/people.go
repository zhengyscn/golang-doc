package hellomock

import "fmt"

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{
		name: name,
	}
}

func (p *Person) SayHello(name string) string {
	return fmt.Sprintf("Hello %s, my name is %s.", name, p.name)
}
