package main 
import ( 
	"fmt"
	"log"
) 
type point struct { 
	x float64 
	y float64
	} 
func pointnew(x float64, y float64) *point { 
		p := new(point)
		p.x = x 
		p.y = y
		return p
		} 
func pointgetx(p *point) float64 { 
	return p.x
	} 
func pointgety(p *point) float64 { 
	return p.y
	} 
func pointsetx(p *point, x float64) { 
	p.x = x
	} 
func pointsety(p *point, y float64) { 
	p.y = y
	} 
func main() { 
	p := pointnew(1, 1)
	fmt.Println(pointnew(1,1))
	fmt.Println(pointgetx(p))
	fmt.Println(pointgety(p))

	pointsetx(p, 3.0)
    pointsety(p, 4.0)
	if !(pointgetx(p) == 3.0) {
        log.Fatal("Wrong value")
    }
 
    if !(pointgety(p) == 4.0) {
        log.Fatal("Wrong value")
    }
}

