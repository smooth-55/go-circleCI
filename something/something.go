package something

type Something struct {
	NumberA int
	NumberB int
}

func (s Something) Add() int {
	return s.NumberA + s.NumberB
}
