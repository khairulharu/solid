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

type WantRobusta interface {
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

func IWantRobustaAndArabica(robusta WantRobusta, arabica WantArabica) {
	robusta.Robusta()
	arabica.Arabica()
}

func main() {
	arabica := Arabica{}
	robusta := Robusta{}

	IWantArabica()
}
