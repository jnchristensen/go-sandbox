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

func (p person) pSpeak() {
	fmt.Println(p.fname, "says hello")
}

func (sa secretAgent) saSpeak() {
	fmt.Println("Keep it secret. Keep it safe.")
}

func main() {
	p := person{"josh", "c"}
	sa := secretAgent{person{"SA", "231"}, "dirty dawg"}

	fmt.Println(p.fname)
	p.pSpeak()

	fmt.Println(sa.codeName)
	sa.saSpeak()
	sa.pSpeak()

}
