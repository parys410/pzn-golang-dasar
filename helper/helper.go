package helper

import "fmt"

type Person struct {
	Name string
	BdayYear int
	Age int
}
func (p *Person) Greetings() {
	fmt.Printf("Hello Guys, I'm %s and today I'm %d years old.", p.Name, p.Age)
}

func SayHelloInHelper(name string) {
	fmt.Println("Hello", name)
}