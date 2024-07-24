package main

import "fmt"

// a interface do componente
type IPiza interface {
	getPrice() int
}

// componente concreto
/*o componente base que vai receber multiplos
comportamentos de forma dinamica*/
type VeggieMania struct {
}

func (v *VeggieMania) getPrice() int {
	return 15
}

// decorador concreto
/*1- os comportamentos de forma dinamica*/
type TomatoTopping struct {
	pizza IPiza
}

func (t *TomatoTopping) getPrice() int {
	price := t.pizza.getPrice()

	return price + 7
}

// decorador concreto
/* 2-os comportamentos de forma dinamica*/
type CheeseTopping struct {
	pizza IPiza
}

func (c *CheeseTopping) getPrice() int {
	price := c.pizza.getPrice()

	return price + 10
}

// main
func main() {
	pizza := &VeggieMania{}

	pizzaWithCheeze := &CheeseTopping{
		pizza: pizza,
	}

	pizzaWithTomate := &TomatoTopping{
		pizza: pizza,
	}

	pizzaWithChesseAndTomate := &TomatoTopping{
		pizza: pizzaWithCheeze,
	}

	fmt.Println("PIZZA BASE: ", pizza.getPrice())
	fmt.Println("PIZZA COM CHEESE: ", pizzaWithCheeze.getPrice())
	fmt.Println("PIZZA TOMATE: ", pizzaWithTomate.getPrice())
	fmt.Println("PIZZA COM CHEESE E TOMATE: ", pizzaWithChesseAndTomate.getPrice())
}
