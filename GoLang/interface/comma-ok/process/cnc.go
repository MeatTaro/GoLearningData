package process

type CNC struct{
	Station
}

func (x CNC) StateNum() string {
	return "No.1"
}

func (x CNC) Machine() string {
	return "Milling machine"
}