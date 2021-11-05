package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	codeName string
}

func (p person) speak() {
	fmt.Println(p.fname, "says hello")
}

func (sa secretAgent) speak() {
	fmt.Println("Keep it secret. Keep it safe.")
}

type human interface {
	speak()
}

func doThing(thingType human) {
	thingType.speak()
}

func main() {
	doThing(person{"joe", "Shmoe"})
	doThing(secretAgent{person{"joe", "smith"}, "Don't Ask"})
}
