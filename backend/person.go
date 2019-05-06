package main

type Person struct {
	Name    string
	Email   string
	Address map[string]interface{}
}

type PersonEvent interface {
	Apply(person *Person)
}

func GetPerson(events []PersonEvent) *Person {
	person := new(Person)

	for _, event := range events {
		event.Apply(person)
	}

	return person
}
