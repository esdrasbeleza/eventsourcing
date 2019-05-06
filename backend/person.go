package main

type Person struct {
	Name string
}

func GetPerson(changePersonName []ChangePersonName) *Person {
	person := new(Person)

	for _, event := range changePersonName {
		person.Name = event.Name
	}

	return person
}
