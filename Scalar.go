// Scalar
package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"io"
	"math/big"
)

type Scalar struct {
	v [32]byte
}

func (s *Scalar) NewKey() *Scalar {
	seed := make([]byte, 32)
	io.ReadFull(rand.Reader, seed)
	aa := new(big.Int).SetBytes(seed[:])
	cc := new(big.Int).Mod(aa, primemod)
	rc := rightcopy(cc.Bytes())
	copy(s.v[:], reversOrder(rc[:]))
	return s
}

func (s *Scalar) Bytes() []byte {
	return s.v[:]
}

func (s *Scalar) String() string {
	return hex.EncodeToString(s.Bytes())
}

func (s *Scalar) Add(a, b *Scalar) *Scalar {
	aa := new(big.Int).SetBytes(reversOrder(a.v[:]))
	bb := new(big.Int).SetBytes(reversOrder(b.v[:]))
	cc := new(big.Int).Add(aa, bb)
	cc = new(big.Int).Mod(cc, primemod)
	rc := rightcopy(cc.Bytes())
	copy(s.v[:], reversOrder(rc[:]))
	return s
}

func (s *Scalar) Sub(a, b *Scalar) *Scalar {
	aa := new(big.Int).SetBytes(reversOrder(a.v[:]))
	bb := new(big.Int).SetBytes(reversOrder(b.v[:]))
	cc := new(big.Int).Sub(aa, bb)
	cc = new(big.Int).Mod(cc, primemod)
	rc := rightcopy(cc.Bytes())
	copy(s.v[:], reversOrder(rc[:]))
	return s
}

func (s *Scalar) Mul(a, b *Scalar) *Scalar {
	aa := new(big.Int).SetBytes(reversOrder(a.v[:]))
	bb := new(big.Int).SetBytes(reversOrder(b.v[:]))
	cc := new(big.Int).Mul(aa, bb)
	cc = new(big.Int).Mod(cc, primemod)
	rc := rightcopy(cc.Bytes())
	copy(s.v[:], reversOrder(rc[:]))
	return s
}

func (s *Scalar) Div(a, b *Scalar) *Scalar {
	aa := new(big.Int).SetBytes(reversOrder(a.v[:]))
	bb := new(big.Int).SetBytes(reversOrder(b.v[:]))
	cc := new(big.Int).Div(aa, bb)
	cc = new(big.Int).Mod(cc, primemod)
	rc := rightcopy(cc.Bytes())
	copy(s.v[:], reversOrder(rc[:]))
	return s
}

func (s *Scalar) Neg(a *Scalar) *Scalar {
	aa := new(big.Int).SetBytes(reversOrder(a.v[:]))
	cc := new(big.Int).Neg(aa)
	cc = new(big.Int).Mod(cc, primemod)
	rc := rightcopy(cc.Bytes())
	copy(s.v[:], reversOrder(rc[:]))
	return s
}

func (s *Scalar) Inv(a *Scalar) *Scalar {
	aa := new(big.Int).SetBytes(reversOrder(a.v[:]))
	cc := new(big.Int).ModInverse(aa, primemod)
	rc := rightcopy(cc.Bytes())
	copy(s.v[:], reversOrder(rc[:]))
	return s
}

func (s *Scalar) Set(a *Scalar) *Scalar {
	copy(s.v[:], a.v[:])
	return s
}

func (s *Scalar) SetInt64(a int64) *Scalar {
	aa := new(big.Int).SetInt64(a)
	cc := new(big.Int).Mod(aa, primemod)
	rc := rightcopy(cc.Bytes())
	copy(s.v[:], reversOrder(rc[:]))
	return s
}

func (s *Scalar) SetUint64(a uint64) *Scalar {
	aa := new(big.Int).SetUint64(a)
	cc := new(big.Int).Mod(aa, primemod)
	rc := rightcopy(cc.Bytes())
	copy(s.v[:], reversOrder(rc[:]))
	return s
}

func (s *Scalar) SetBytes(b []byte) *Scalar {
	aa := new(big.Int).SetBytes(reversOrder(b))
	cc := new(big.Int).Mod(aa, primemod)
	rc := rightcopy(cc.Bytes())
	copy(s.v[:], reversOrder(rc[:]))
	return s
}

func (s *Scalar) Zero() *Scalar {
	s.v = [32]byte{0}
	return s
}

func (s *Scalar) One() *Scalar {
	s.SetInt64(1)
	return s
}

func (s *Scalar) Clone() *Scalar {
	s2 := *s
	return &s2
}

func (s *Scalar) Equal(a *Scalar) bool {
	return bytes.Equal(a.v[:], s.v[:])
}
