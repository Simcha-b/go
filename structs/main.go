package main

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {
    p := person{}
    p.age = 42
    return &p
}