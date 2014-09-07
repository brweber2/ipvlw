package main

import (
	"github.com/brweber2/interwebs/rpvlw"
//	"github.com/brweber2/interwebs/ipvlw"
	"fmt"
//	"log"
	"log"
	"github.com/brweber2/interwebs/ipvlw"
)

//type RouterNic struct {
//	Addr ipvlw.Address
//	IO chan ipvlw.Message
//}
//
//func (r RouterNic) Send(m ipvlw.Message) (ipvlw.Message, error) {
//	r.IO <- m
//	resp := <- r.IO
//	return resp, nil
//}
//
//func (r RouterNic) Listen() ([]ipvlw.Message, error) {
//	resp := <-r.IO
//	var a []ipvlw.Message
//	a = append(a, resp)
//	return a, nil
//}
//
//type RouterData struct {
//	Running bool
//	IO chan ipvlw.Message
//}
//
//func (r RouterData) Start() error {
//	if r.Running {
//		panic("router is already running")
//	}
//	r.Running = true
//	var routeDb rpvlw.RouteDatabase
//	for {
//		if r.Running == false {
//			break
//		}
//
//		message := <-r.IO
//		to := message.To()
//		block, err := containsAddress(routeDb.Routes, to)
//		if err != nil {
//			log.Printf("error while looking for the route that contains %v: %v\n", to, err)
//		}
//		system := routeDb.Routes[*block]
//		nic := routeDb.Interfaces[system]
//		resp, err := nic.Send(message)
//		if err != nil {
//			log.Printf("error while trying to send message to %v: %v\n", to, err)
//		}
//		if resp != nil {
//			r.IO <- resp
//		}
//
//	}
//	return nil
//}
//
//func containsAddress(routes map[ipvlw.Block]rpvlw.System, a ipvlw.Address) (*ipvlw.Block, error) {
//	for block := range routes {
//		if block.Contains(a) {
//			return &block, nil
//		}
//	}
//	return nil, nil
//}
//
//func (r RouterData) Stop() error {
//	r.Running = false
//	return nil
//}
//
//func main() {
//	io := make(chan ipvlw.Message)
//	routerNic := RouterNic{Addr: ipvlw.Address{1}, IO: io}
//	fmt.Println(routerNic)
//	hello := ipvlw.TextMessage{FromAddr: ipvlw.Address{3}, ToAddr: ipvlw.Address{4}, Identifier: 72, Body: "yeah"}
////	io <-hello
//	routerNic.Send(hello)
////	router := rpvlw.Router{System: System{15}, ControlPlane: , DataPlane: }
//}

func main() {
	simulateRouting()
}

func simulateRouting() {

	/*
				1
			/	   \
		  /         \
		2	-------- 3	------	4
		   \
		   	\
		   	 \
		   	  \
				5


	 */

	// define some routers
	router_1 := rpvlw.MakeAndStartRouter(1)
	router_2 := rpvlw.MakeAndStartRouter(2)
	router_3 := rpvlw.MakeAndStartRouter(3)
	router_4 := rpvlw.MakeAndStartRouter(4)
	router_5 := rpvlw.MakeAndStartRouter(5)

	fmt.Printf("routers: %#v %#v %#v %#v %#v\n", router_1, router_2, router_3, router_4, router_5)

	// build the network topology
	err := router_1.ConnectTo(router_2, router_3)
	if err != nil {
		log.Fatalf("Unable to connect routers with error %v\n", err)
	}
	err = router_2.ConnectTo(router_3, router_5)
	if err != nil {
		log.Fatalf("Unable to connect routers with error %v\n", err)
	}
	err = router_3.ConnectTo(router_4)
	if err != nil {
		log.Fatalf("Unable to connect routers with error %v\n", err)
	}

	fmt.Printf("after routers: %#v %#v %#v %#v %#v\n", router_1, router_2, router_3, router_4, router_5)
	fmt.Printf("router 1 connected to %d\n", len(router_1.ControlPlane.Routers()))
	fmt.Printf("router 2 connected to %d\n", len(router_2.ControlPlane.Routers()))
	fmt.Printf("router 3 connected to %d\n", len(router_3.ControlPlane.Routers()))
	fmt.Printf("router 4 connected to %d\n", len(router_4.ControlPlane.Routers()))
	fmt.Printf("router 5 connected to %d\n", len(router_5.ControlPlane.Routers()))

	// announce some routes
	err = router_1.Announce(ipvlw.Block{ipvlw.Address{4}, 6})
	if err != nil {
		log.Fatalf("Unable to announce routes with error %v\n", err)
	}
	err = router_1.Announce(ipvlw.Block{ipvlw.Address{20}, 6})
	if err != nil {
		log.Fatalf("Unable to announce routes with error %v\n", err)
	}
	err = router_2.Announce(ipvlw.Block{ipvlw.Address{40}, 6})
	if err != nil {
		log.Fatalf("Unable to announce routes with error %v\n", err)
	}
	err = router_3.Announce(ipvlw.Block{ipvlw.Address{60}, 6})
	if err != nil {
		log.Fatalf("Unable to announce routes with error %v\n", err)
	}
	err = router_4.Announce(ipvlw.Block{ipvlw.Address{100}, 4})
	if err != nil {
		log.Fatalf("Unable to announce routes with error %v\n", err)
	}
	err = router_5.Announce(ipvlw.Block{ipvlw.Address{200}, 6})
	if err != nil {
		log.Fatalf("Unable to announce routes with error %v\n", err)
	}


	//
//	// define some computers
//	nic_1 := MakeNic()
//	nic_2 := MakeNic()
//	nic_3 := MakeNic()
//	nic_4 := MakeNic()
//	nic_5 := MakeNic()
//	nic_6 := MakeNic()
//	nic_7 := MakeNic()
//	nic_8 := MakeNic()
//	nic_9 := MakeNic()
//	nic_10 := MakeNic()
//	nic_11 := MakeNic()
//	nic_12 := MakeNic()
//	nic_13 := MakeNic()
//	nic_14 := MakeNic()
//	nic_15 := MakeNic()
//	nic_16 := MakeNic()
//
//	// "dhcp" will assing computers to routers (as if they were physically plugged in there)
//	// and it will assign IPvLW addresses to the network interfaces
//	dhcp := MakeDhcp(router_1, router_2, router_3, router_4, router_5)
//
//	// plug computers into the network
//	dhcp.ConnectTo(router_1, nic_1, nic_2, nic_3, nic_4)
//	dhcp.ConnectTo(router_2, nic_5, nic_6, nic_7)
//	dhcp.ConnectTo(router_3, nic_8, nic_9, nic_10)
//	dhcp.ConnectTo(router_4, nic_11, nic_12, nic_13, nic_14)
//	dhcp.ConnectTo(router_5, nic_15, nic_16)
//
//	// send a message from computer 2 to computer 4 (same router - stays internal)
//	to := nic_15.Address()
//	from := nic_2.Address()
//
//	nic_2.Send(ipvlw.TextMessage{from, to, ipvlw.GenerateId(), "hello world"})
//	responses, err := nic_2.Listen()
//
//	// send a message from computer 2 to computer 15 (different routers - external traffic)

}

//func MakeNic() rpvlw.Nic {
//	return nil
//}
//
//func MakeDhcp(routers ... Router) rpvlw.Dhcp {
//	return nil
//}



