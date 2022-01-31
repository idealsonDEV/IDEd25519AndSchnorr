// ed25519
package main

import (
	"fmt"
)

var curve = Suite{}

func main() {
	p1 := curve.Point().SetInt64(2)
	p2 := curve.Point().Base()
	fmt.Println(curve.Point().Add(p1, p2))
	fmt.Println(new(Point).Neg(new(Point).SetUnt64(5)))
}
