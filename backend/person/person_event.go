package person

type PersonEvent interface {
	Apply(person *Person)
	JSON() []byte
}
