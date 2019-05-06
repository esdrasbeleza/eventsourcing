package main

import "github.com/google/uuid"

type Person struct {
	Id      uuid.UUID
	Name    string
	Email   string
	Address map[string]interface{}
}

type PersonEvent interface {
	Apply(person *Person)
}

func NewPerson() *Person {
	return &Person{
		Id: uuid.New(),
	}
}

func GetPerson(events []PersonEvent) *Person {
	person := new(Person)

	for _, event := range events {
		event.Apply(person)
	}

	return person
}
