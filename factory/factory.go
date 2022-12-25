package main

import "fmt"

//抽象层
type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit() Fruit
}

//实现层
type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("apple created")
}

func (fac *AppleFactory) CreateFruit() Fruit {
	return new(Apple)
}


//基础工厂模块

type AppleFactory struct {
	AbstractFactory
}


