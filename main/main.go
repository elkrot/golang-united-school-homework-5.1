package main
import (
	"fmt"
	"square"
)
func main()  {
	p:=square.Point{X:1,Y:2}
	s:= square.Square{p,5}
	fmt.Println(s.Area())
}
