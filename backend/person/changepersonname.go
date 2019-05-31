package person

type ChangePersonName struct {
	Name string
}

func (event ChangePersonName) Apply(person *Person) {
	person.Name = event.Name
}
