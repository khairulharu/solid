package main

import "fmt"

type Coffe interface {
	Arabica()
	Robusta()
	Batista()
}

type WantArabica interface {
	Arabica()
}

type WanRobusta interface {
	Robusta()
}

type WantBatista interface {
	Batista()
}

type Arabica struct{}

func (a *Arabica) Arabica() {
	fmt.Println("Arabica")
}

type Robusta struct{}

func (r Robusta) Robusta() {
	fmt.Println("Robusta")
}

func IWantArabica(a Arabica) {
	a.Arabica()
}

func IWantRobustaAndArabica(c Coffe, a WantArabica) {
	c.Robusta()
	a.Arabica()
}

func main() {
	arabica := Arabica{}
	robusta := Robusta{}

	IWantArabica(arabica)
}
