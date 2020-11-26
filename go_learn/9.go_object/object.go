package main 
import ( 
	"fmt" 
) 
//Point is a type 
type Point struct{	//將type宣告所使用的struct為class 
	X float64		// `X` and `Y` are public fields. 
	Y float64 
} 
func newpoint(x, y float64) *Point  { //建構的func作為obiect 
	p := new(Point) 
	p.X = x 
	p.Y = y 
	return p 
} 
func main()  { 
	p := newpoint(3, 4) 
	fmt.Println(p.X, p.Y) 
}