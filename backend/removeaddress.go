package main

type RemoveAddress struct {
	Name string
}

func (e RemoveAddress) Apply(person *Person) {
	if person.Address == nil {
		person.Address = make(map[string]interface{})
		return
	}

	delete(person.Address, e.Name)
}
