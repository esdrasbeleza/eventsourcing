package person

type AddEmail struct {
	Email string
}

func (event AddEmail) Type() string {
	return "AddEmail"
}

func (event AddEmail) Apply(person *Person) {
	person.Email = event.Email
}

func (event AddEmail) JSON() []byte {
	return outputJSON(event)
}
