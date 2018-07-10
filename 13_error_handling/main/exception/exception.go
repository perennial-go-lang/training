package exception

type Exception interface {
	getError() error
}

type DivideByZero struct {
	Error error
}

func (d DivideByZero) getError() error {
	return d.Error
}
