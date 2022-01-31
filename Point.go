// Point
package main

import (
	"IDEd25519AndSchnorr/edwards25519"
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"io"
	"math/big"
)

type Point struct {
	w [32]byte
}

func (p *Point) MultBase(a *Scalar) *Point {
	var A edwards25519.ExtendedGroupElement
	edwards25519.GeScalarMultBase(&A, &a.v)
	A.ToBytes(&p.w)
	return p
}

func (p *Point) Bytes() []byte {
	return p.w[:]
}

func (p *Point) Base() *Point {
	one := new(Scalar).One()
	p.MultBase(one)
	return p
}

func (p *Point) Null() *Point {
	zero := new(Scalar).Zero()
	p.MultBase(zero)
	return p
}

func (p *Point) String() string {
	return hex.EncodeToString(p.Bytes())
}

func (p *Point) Clone() *Point {
	p2 := *p
	return &p2
}

func (p *Point) Equal(a *Point) bool {
	return bytes.Equal(a.w[:], p.w[:])
}

func (p *Point) Set(a *Point) *Point {
	copy(p.w[:], a.w[:])
	return p
}

func (p *Point) SetInt64(a int64) *Point {
	aa := new(big.Int).SetInt64(a)
	rc := rightcopy(aa.Bytes())
	copy(p.w[:], reversOrder(rc[:]))
	return p
}

func (p *Point) SetUnt64(a uint64) *Point {
	aa := new(big.Int).SetUint64(a)
	rc := rightcopy(aa.Bytes())
	copy(p.w[:], reversOrder(rc[:]))
	return p
}

func (p *Point) SetBytes(b []byte) *Point {
	copy(p.w[:], b[:])
	return p
}

func (p *Point) NewPoint() *Point {
	seed := make([]byte, 32)
	io.ReadFull(rand.Reader, seed)
	for {
		if bytes.Equal(seed[:], new(Point).Null().Bytes()) {
			io.ReadFull(rand.Reader, seed)
		} else if bytes.Equal(seed[:], new(Point).Base().Bytes()) {
			io.ReadFull(rand.Reader, seed)
		} else {
			break
		}
	}
	copy(p.w[:], seed[:])
	return p
}

func (p *Point) Mul(a *Scalar, M *Point) *Point {
	var A edwards25519.ExtendedGroupElement
	var R edwards25519.ProjectiveGroupElement
	A.FromBytes(&M.w)
	edwards25519.GeScalarMult(&R, &a.v, &A)
	R.ToBytes(&p.w)
	return p
}

func (p *Point) Add(a, b *Point) *Point {
	var R edwards25519.CompletedGroupElement
	var P edwards25519.ExtendedGroupElement
	P.FromBytes(&a.w)
	var q edwards25519.ExtendedGroupElement
	q.FromBytes(&b.w)
	var Q edwards25519.CachedGroupElement
	q.ToCached(&Q)
	edwards25519.GeAdd(&R, &P, &Q)
	var r edwards25519.ExtendedGroupElement
	R.ToExtended(&r)
	r.ToBytes(&p.w)
	return p
}

func (p *Point) Sub(a, b *Point) *Point {
	var R edwards25519.CompletedGroupElement
	var P edwards25519.ExtendedGroupElement
	P.FromBytes(&a.w)
	var q edwards25519.ExtendedGroupElement
	q.FromBytes(&b.w)
	var Q edwards25519.CachedGroupElement
	q.ToCached(&Q)
	edwards25519.GeSub(&R, &P, &Q)
	var r edwards25519.ExtendedGroupElement
	R.ToExtended(&r)
	r.ToBytes(&p.w)
	return p
}

func (p *Point) Neg(a *Point) *Point {
	p.Sub(new(Point).Null(), a)
	return p
}
