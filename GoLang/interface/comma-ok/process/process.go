package process

type Station struct{
	Name string
}

// func NewStation(name string) *Station {
// 	return &Station{name:name}
// }

func (x Station) GetName() string {
	return x.Name
}

func (x Station) StateNum() string {
	return "工站編號"
}

func (x Station) Machine() string {
	return "機台名稱"
}