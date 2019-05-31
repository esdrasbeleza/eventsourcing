package person

type AddAddress struct {
	Name    string
	Address string
}

func (event AddAddress) Apply(person *Person) {
	if person.Address == nil {
		person.Address = make(map[string]interface{})
	}

	person.Address[event.Name] = event.Address
}

func (event AddAddress) JSON() []byte {
	return outputJSON(event)
}
