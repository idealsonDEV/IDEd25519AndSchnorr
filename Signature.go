// Signature
package main

import (
	"crypto/sha256"
)

type Signature struct {
	R Scalar
	S Point
}

func (sig *Signature) Bytes() []byte {
	r := sig.R.Bytes()
	s := sig.S.Bytes()
	binary := append(r, s...)
	return binary
}

func (sig *Signature) SetBytes(data []byte) *Signature {
	sig.R.SetBytes(data[:32])
	sig.S.SetBytes(data[32:])
	return sig
}

func HashMsg(msg []byte) []byte {
	sum := sha256.Sum256(msg)
	return sum[:]
}
