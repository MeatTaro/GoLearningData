package process

type Iprocess interface {
	GetName() string
	StateNum() string
	Machine() string
}