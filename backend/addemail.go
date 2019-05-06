package main

type AddEmail struct {
	Email string
}

func (event AddEmail) Apply(person *Person) {
	person.Email = event.Email
}
