package person

type RemoveAddress struct {
	Name string
}

func (event RemoveAddress) Type() string {
	return "RemoveAddress"
}

func (e RemoveAddress) Apply(person *Person) {
	if person.Address == nil {
		person.Address = make(map[string]string)
		return
	}

	delete(person.Address, e.Name)
}

func (event RemoveAddress) JSON() []byte {
	return outputJSON(event)
}
