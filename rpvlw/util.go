package rpvlw

import (
	"github.com/brweber2/interwebs/ipvlw"
	"log"
)

func MakeAndStartRouter(s int) *Router {
	controlPlane := RouterControlPlane{
		Router: nil,
		LocalBlocks: make([]*ipvlw.Block, 0, 256),
		Computers: make([]Nic, 0, 16),
		Addresses: make(map[*ipvlw.Address]Nic),
		Nics: make([]*Router, 0, 16),
		Routes: make(map[*ipvlw.Block]*System),
		Interfaces: make(map[*System]*Router),
	}
	dataPlane := RouterDataPlane{
		nil,
	}

	r := Router{System{uint8(s)}, &controlPlane, &dataPlane}
	controlPlane.Router = &r
	dataPlane.Router = &r

	r.Start()
	return &r
}

func MakeNic() Nic {
	log.Printf("making a nic")
	return &Computer{nil, ipvlw.Address{0}, nil}
}

func MakeDhcp(routers ... *Router) Dhcp {
	dhcp := RouterDhcp{make([]*Router, 0, 16)}
	return dhcp
}
