package main

import (
	"github.com/brweber2/interwebs/rpvlw"
	"github.com/brweber2/interwebs/ipvlw"
	"fmt"
	"log"
)

type RouterNic struct {
	Addr ipvlw.Address
	IO chan ipvlw.Message
}

func (r RouterNic) Send(m ipvlw.Message) (ipvlw.Message, error) {
	r.IO <- m
	resp := <- r.IO
	return resp, nil
}

func (r RouterNic) Listen() ([]ipvlw.Message, error) {
	resp := <-r.IO
	var a []ipvlw.Message
	a = append(a, resp)
	return a, nil
}

type RouterData struct {
	Running bool
	IO chan ipvlw.Message
}

func (r RouterData) Start() error {
	if r.Running {
		panic("router is already running")
	}
	r.Running = true
	var routeDb rpvlw.RouteDatabase
	for {
		if r.Running == false {
			break
		}

		message := <-r.IO
		to := message.To()
		block, err := containsAddress(routeDb.Routes, to)
		if err != nil {
			log.Printf("error while looking for the route that contains %v: %v\n", to, err)
		}
		system := routeDb.Routes[*block]
		nic := routeDb.Interfaces[system]
		resp, err := nic.Send(message)
		if err != nil {
			log.Printf("error while trying to send message to %v: %v\n", to, err)
		}
		if resp != nil {
			r.IO <- resp
		}

	}
	return nil
}

func containsAddress(routes map[ipvlw.Block]rpvlw.System, a ipvlw.Address) (*ipvlw.Block, error) {
	for block := range routes {
		if block.Contains(a) {
			return &block, nil
		}
	}
	return nil, nil
}

func (r RouterData) Stop() error {
	r.Running = false
	return nil
}

func main() {
	io := make(chan ipvlw.Message)
	routerNic := RouterNic{Addr: ipvlw.Address{1}, IO: io}
	fmt.Println(routerNic)
	hello := ipvlw.TextMessage{FromAddr: ipvlw.Address{3}, ToAddr: ipvlw.Address{4}, Identifier: 72, Body: "yeah"}
//	io <-hello
	routerNic.Send(hello)
//	router := rpvlw.Router{System: System{15}, ControlPlane: , DataPlane: }
}

