package process

type CNC struct {
	*Station
}

func (y CNC) StateNum() string  {
	return "No.1"
}

func (y CNC) Machine() string {
	return "Milling machine"
}