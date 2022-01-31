// const
package main

import (
	"math/big"
)

var primemod, _ = new(big.Int).SetString("57896044618658097711785492504343953926634992332820282019728792003956564819949", 10)

func rightcopy(l []byte) [32]byte {
	r := [32]byte{0}
	var j int = 31
	for i := len(l) - 1; i >= 0; i-- {
		r[j] = l[i]
		j--
	}
	return r
}

func reversOrder(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

/*
type Scalar interface {
	NewKeyFromRandom() Scalar
	Add(a, b Scalar) Scalar
	Bytes() []byte
	String()
}
*/
