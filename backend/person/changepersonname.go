package person

type ChangePersonName struct {
	Name string
}

func (event ChangePersonName) Type() string {
	return "ChangePersonName"
}

func (event ChangePersonName) Apply(person *Person) {
	person.Name = event.Name
}

func (event ChangePersonName) JSON() []byte {
	return outputJSON(event)
}
