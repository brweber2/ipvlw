package ipvlw

import (
	"math/rand"
)

type Address struct {
	Address uint8 // this is a huge interweb!
}

type Block struct {
	Start Address
	Bits uint8
}

type Header interface {
	From() Address
	To() Address
	Id() uint64
}

type Message interface {
	Header
	Payload() string
}

type TextMessage struct {
	FromAddr Address
	ToAddr Address
	Identifier uint64
	Body string
}

func (t TextMessage) From() Address {
	return t.FromAddr
}

func (t TextMessage) To() Address {
	return t.ToAddr
}

func (t TextMessage) Id() uint64 {
	return t.Identifier
}

func (t TextMessage) Payload() string {
	return t.Body
}

func (b Block) Contains(a Address) bool {
	bits := b.Bits
	s := b.Start.Address & Mask(bits)
	e := a.Address & Mask(bits)
	return s == e
}

func Mask(bits uint8) uint8 {
	return 0xFF << (8-bits)
}

func GenerateId() uint64 {
	l := rand.Uint32()
	r := rand.Uint32()

	a := uint64(r)
	b := uint64(l) << 32

	id := a | b
	return id
}
