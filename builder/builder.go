package main

import "fmt"

// interface do builder
type Ibuilder interface {
	setWindowsType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(buildertype string) Ibuilder {
	if buildertype == "normal" {
		return newNormalBuilder()
	}

	if buildertype == "igloo" {
		return newIgloobuilder()
	}

	return nil
}

// builder concreto 1
type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowsType() {
	b.windowType = "wooden windows"
}

func (b *NormalBuilder) setDoorType() {
	b.doorType = "wooden door"
}

func (b *NormalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *NormalBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}

// builder concreto 2
type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// setDoortype implements Ibuilder.

func newIgloobuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (c *IglooBuilder) setWindowsType() {
	c.windowType = "snow windows"
}

// setDoorType implements Ibuilder.
func (c *IglooBuilder) setDoorType() {
	c.doorType = "snow door"
}

func (c *IglooBuilder) setNumFloor() {
	c.floor = 1
}

func (c *IglooBuilder) getHouse() House {
	return House{
		windowType: c.windowType,
		doorType:   c.doorType,
		floor:      c.floor,
	}
}

// Produto
type House struct {
	windowType string
	doorType   string
	floor      int
}

// o director

type Director struct {
	builder Ibuilder
}

func newDirector(b Ibuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b Ibuilder) {
	d.builder = b
}

func (d *Director) buildGHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowsType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

// main
func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildGHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	fmt.Println("Novais 1")
	iglooHouse := director.buildGHouse()
	fmt.Println("Novais 2")

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

}

// mini-prática de builder
