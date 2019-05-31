package person

type PersonEvent interface {
	Apply(person *Person)
	Type() string
	JSON() []byte
}
