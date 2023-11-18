package main

import "fmt"

func main() {
	Greet()
	PersonalGreeting("Andrey")
	MultipleGreeting("Andrey", "Maxim")
	FioGreeting("Pupkin", "Andrey", "Maximovich")
}

// Тут я буду жеска объявлять функции

func Greet() {
	fmt.Println("Hello, World!")
}
func PersonalGreeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
func MultipleGreeting(name1, name2 string) {
	fmt.Printf("Hello, %s and %s. Have a nice day, guys.\n", name1, name2)
}
func FioGreeting(fname, name, patronim string) {
	fmt.Printf("My full name is %s %s %s.\n", fname, name, patronim)
}
