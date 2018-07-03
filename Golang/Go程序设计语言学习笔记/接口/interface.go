package main

import (
	"fmt"
)

func main() {
	var p People
	p = &Student{Name: "Zhou Yang", Age: 30}
	fmt.Println(p.SayHello())
}

type Student struct {
	Name string
	Age  int
}

type People interface {
	SayHello() string
}

func (s *Student) SayHello() string {
	s.Age++
	return fmt.Sprintf("My name is %s with %d years old", s.Name, s.Age)
}
