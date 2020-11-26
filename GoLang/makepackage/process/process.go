package process

type Station struct {
	name string
}

func NewStation(name string) *Station {
	return &Station{name:name}
}

func (x Station) GetName() string  {
	return x.name
}

func (x Station) StateNum() string {
	return "工站編號"
}

func (x Station) Machine() string  {
	return "機台名稱"
}