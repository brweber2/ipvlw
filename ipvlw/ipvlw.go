package ipvlw

import (
	"math/rand"
	"fmt"
	"log"
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
	log.Printf("comparing %v and %v\n", s, e)
	return s == e
}

func (b Block) Addresses() []*Address {
	answer := make([]*Address, 1 << (8 - b.Bits))
	var i uint8
	for i = 0; i < (1 << (8 - b.Bits)); i++ {
		answer[i] = &Address{b.Start.Address + i}
	}
	return answer
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

func MakeResponse( m Message, resp string) (Message, error) {
	f := m.From()
	if &f == nil {
		return nil, fmt.Errorf("No from address on original message %v\n", m)
	}
	t := m.To()
	if &t == nil {
		return nil, fmt.Errorf("No to address on original message %v\n", m)
	}
	i := m.Id()
	if &i == nil {
		return nil, fmt.Errorf("No id on original message %v\n", m)
	}
	return TextMessage{f, t, i, resp}, nil
}
