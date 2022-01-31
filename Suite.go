// Suite
package main

type Suite struct {
}

func NewSuite() *Suite {
	return new(Suite)
}

func (st *Suite) Scalar() *Scalar {
	return new(Scalar)
}

func (st *Suite) Point() *Point {
	return new(Point)
}
