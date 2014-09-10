package main

import (
	"github.com/brweber2/interwebs/rpvlw"
	"log"
	"github.com/brweber2/interwebs/ipvlw"
)

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

//	fmt.Printf("routers: %#v %#v %#v %#v %#v\n", router_1, router_2, router_3, router_4, router_5)

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

	for _, rtr := range(router_1.ControlPlane.Routers()) {
		log.Printf("router %d is connected to %d\n", router_1.System.Identifier, rtr.System.Identifier)
	}
	for _, rtr := range(router_2.ControlPlane.Routers()) {
		log.Printf("router %d is connected to %d\n", router_2.System.Identifier, rtr.System.Identifier)
	}
	for _, rtr := range(router_3.ControlPlane.Routers()) {
		log.Printf("router %d is connected to %d\n", router_3.System.Identifier, rtr.System.Identifier)
	}
	for _, rtr := range(router_4.ControlPlane.Routers()) {
		log.Printf("router %d is connected to %d\n", router_4.System.Identifier, rtr.System.Identifier)
	}
	for _, rtr := range(router_5.ControlPlane.Routers()) {
		log.Printf("router %d is connected to %d\n", router_5.System.Identifier, rtr.System.Identifier)
	}

//	fmt.Printf("after routers: %#v %#v %#v %#v %#v\n", router_1, router_2, router_3, router_4, router_5)
//	fmt.Printf("router 1 connected to %d routers\n", len(router_1.ControlPlane.Routers()))
//	fmt.Printf("router 2 connected to %d routers\n", len(router_2.ControlPlane.Routers()))
//	fmt.Printf("router 3 connected to %d routers\n", len(router_3.ControlPlane.Routers()))
//	fmt.Printf("router 4 connected to %d routers\n", len(router_4.ControlPlane.Routers()))
//	fmt.Printf("router 5 connected to %d routers\n", len(router_5.ControlPlane.Routers()))

	// announce some routes
	err = router_1.Originate(&ipvlw.Block{ipvlw.Address{4}, 6})
	if err != nil {
		log.Fatalf("Unable to originate routes with error %v\n", err)
	}
	err = router_1.Originate(&ipvlw.Block{ipvlw.Address{20}, 6})
	if err != nil {
		log.Fatalf("Unable to originate routes with error %v\n", err)
	}
	err = router_2.Originate(&ipvlw.Block{ipvlw.Address{40}, 6})
	if err != nil {
		log.Fatalf("Unable to originate routes with error %v\n", err)
	}
	err = router_3.Originate(&ipvlw.Block{ipvlw.Address{60}, 6})
	if err != nil {
		log.Fatalf("Unable to originate routes with error %v\n", err)
	}
	err = router_4.Originate(&ipvlw.Block{ipvlw.Address{100}, 4})
	if err != nil {
		log.Fatalf("Unable to originate routes with error %v\n", err)
	}
	err = router_5.Originate(&ipvlw.Block{ipvlw.Address{200}, 6})
	if err != nil {
		log.Fatalf("Unable to originate routes with error %v\n", err)
	}

	log.Printf("**** ROUTING TABLES ****\n")
	router_1.ControlPlane.PrintRoutes()
	router_2.ControlPlane.PrintRoutes()
	router_3.ControlPlane.PrintRoutes()
	router_4.ControlPlane.PrintRoutes()
	router_5.ControlPlane.PrintRoutes()

//	fmt.Printf("after routers: %#v %#v %#v %#v %#v\n", router_1, router_2, router_3, router_4, router_5)

	// define some computers
	nic_1 := rpvlw.MakeNic()
	nic_2 := rpvlw.MakeNic()
	nic_3 := rpvlw.MakeNic()
	nic_4 := rpvlw.MakeNic()
	nic_5 := rpvlw.MakeNic()
	nic_6 := rpvlw.MakeNic()
	nic_7 := rpvlw.MakeNic()
	nic_8 := rpvlw.MakeNic()
	nic_9 := rpvlw.MakeNic()
	nic_10 := rpvlw.MakeNic()
	nic_11 := rpvlw.MakeNic()
	nic_12 := rpvlw.MakeNic()
	nic_13 := rpvlw.MakeNic()
	nic_14 := rpvlw.MakeNic()
	nic_15 := rpvlw.MakeNic()
	nic_16 := rpvlw.MakeNic()

//	log.Printf("making nics %#v,%#v,%#v,%#v,%#v,%#v,%#v,%#v,%#v,%#v,%#v,%#v,%#v,%#v,%#v,%#v \n",
//	nic_1, nic_2, nic_3, nic_4, nic_5, nic_6, nic_7, nic_8, nic_9, nic_10, nic_11, nic_12, nic_13,
//	nic_14, nic_15, nic_16)

	// "dhcp" will assing computers to routers (as if they were physically plugged in there)
	// and it will assign IPvLW addresses to the network interfaces
	dhcp := rpvlw.MakeDhcp(router_1, router_2, router_3, router_4, router_5)

//	log.Printf("dhcp: %#v\n", dhcp)


	// plug computers into the network
	dhcp.ConnectTo(router_1, nic_1, nic_2, nic_3, nic_4)
	dhcp.ConnectTo(router_2, nic_5, nic_6, nic_7)
	dhcp.ConnectTo(router_3, nic_8, nic_9, nic_10)
	dhcp.ConnectTo(router_4, nic_11, nic_12, nic_13, nic_14)
	dhcp.ConnectTo(router_5, nic_15, nic_16)

	// send a message from computer 2 to computer 4 (same router - stays internal)
	to := nic_4.Address()
	from := nic_2.Address()

	nic_4.RegisterCallback(func(n rpvlw.Nic, m ipvlw.Message) error {
		log.Printf("nic 4 recieved message %v\n", m)
		resp, err := ipvlw.MakeResponse(m, "right back at ya!")
		if err != nil {
			return err
		}
		log.Printf("response message %v\n", resp)
		return n.Send(resp)
	})

	nic_2.RegisterCallback(func(n rpvlw.Nic, m ipvlw.Message) error {
		log.Printf("nic 2 recieved message %v\n", m)
		return nil
	})

	log.Printf("let's send a message from %v to %v\n", from, to)
	err = nic_2.Send(ipvlw.TextMessage{from, to, ipvlw.GenerateId(), "hello there"})
	if err != nil {
		log.Fatalf("unable to send message: %v\n", err)
	}

	log.Printf("***** DONE WITH INTERAL ROUTING EXAMPLE *****\n")

	from = nic_1.Address()
	to = nic_15.Address()

	log.Printf("trying to send message from %v to %v\n", from, to)

	nic_15.RegisterCallback(func(n rpvlw.Nic, m ipvlw.Message) error {
		log.Printf("nic 15 recieved message %v\n", m)
		resp, err := ipvlw.MakeResponse(m, "oh did you now?")
		if err != nil {
			return err
		}
		log.Printf("response message %v\n", resp)
		return n.Send(resp)
	})

	nic_1.RegisterCallback(func(n rpvlw.Nic, m ipvlw.Message) error {
		log.Printf("nic 1 recieved message %v\n", m)
		return nil
	})

	err = nic_1.Send(ipvlw.TextMessage{from, to, ipvlw.GenerateId(), "hello world"})
	if err != nil {
		log.Fatalf("unable to send message from 1 to 15\n")
	}

	log.Printf("***** DONE WITH INTERAL ROUTING EXAMPLE *****\n")
	// send a message from computer 2 to computer 15 (different routers - external traffic)

}





