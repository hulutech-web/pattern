package main

import "fmt"

//抽象层
type Fruit interface {
	Show()
}

//实现层
type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("apple")
}

//================================================================================================
type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("banana")

}

//================================================================================================
type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("pear")
}

//工厂模式
type Factory struct{}

func (fac *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit
}

