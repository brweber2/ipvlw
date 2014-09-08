package rpvlw

import (
	"fmt"
	"log"
	"github.com/brweber2/interwebs/ipvlw"
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

	r := Router{System{uint8(s)}, &controlPlane, RouterDataPlane{}}
	controlPlane.Router = &r
	r.Start()
	return &r
}

func (r Router) Start() {
	r.ControlPlane.Start()
	r.DataPlane.Start()
}

type RouterControlPlane struct {
	Router *Router
	// local to this network
	LocalBlocks []*ipvlw.Block
	Computers []Nic
	Addresses map[*ipvlw.Address]Nic
	// external to this network
	Nics []*Router // todo rename me!
	Routes map[*ipvlw.Block]*System
	Interfaces map[*System]*Router
}

func (r *RouterControlPlane) String() string {
	return fmt.Sprintf("router\n\tsystem: %v\n\trouters: %#v\n\tnics: %#v\n\troutes: %#v\n\t", r.Router.System, r.Routes, r.Nics, r.Interfaces)
}

func (r *RouterControlPlane) AddressInUse(a *ipvlw.Address) bool {
	if _,ok := r.Addresses[a]; ok {
		return true
	}
	return false
}

func (r *RouterControlPlane) UnusedAddress() (*ipvlw.Address, error) {
	for _, block := range(r.LocalBlocks) {
		for _, addr := range(block.Addresses()) {
			if ! r.AddressInUse(addr) {
				return addr, nil
			}
		}
	}
	return &ipvlw.Address{}, fmt.Errorf("Unable to find an available ip address in %#v\n", r)
}

func (r *RouterControlPlane) AddComputer(n Nic) error {
	// find unused ip address
	addr, err := r.UnusedAddress()
	if err != nil {
		return err
	}
	n.addr(addr)
	r.Addresses[addr] = n
	r.Computers = append(r.Computers, n)
	return nil
}

type RouterDataPlane struct {

}

type Runnable interface {
	Start()
	Stop()
}

func (r *RouterControlPlane) Start() {
	fmt.Printf("starting control plane\n")
}

func (r RouterDataPlane) Start() {
	fmt.Printf("starting data plane\n")
}

func (r *RouterControlPlane) Stop() {
	fmt.Printf("stopping control plane\n")
}

func (r RouterDataPlane) Stop() {
	fmt.Printf("stopping data plane\n")
}

func (r Router) ConnectTo(routers ... *Router) error {
	for _, router := range(routers) {
		err := r.ControlPlane.AddNic(router)
		if err != nil {
			return err
		}
		err = router.ControlPlane.AddNic(&r)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RouterControlPlane) AddNic(rtr *Router) error {
	log.Printf("adding router %v to %v\n", rtr, r.Nics)
	r.Nics = append(r.Nics, rtr)
	log.Printf("added router %v to %v\n", rtr, r.Nics)
	r.Interfaces[&r.Router.System] = rtr
	return nil
}

func (r *RouterControlPlane) Puters() []Nic {
	return r.Computers
}

func (r *RouterControlPlane) Routers() []*Router {
	return r.Nics
}

func (r *RouterControlPlane) AddRoute(s *System, b *ipvlw.Block) error {
	if s.Identifier == r.Router.System.Identifier {
		r.LocalBlocks = append(r.LocalBlocks, b)
	}
	r.Routes[b] = s
	return nil
}

func (r Router) Announce(b *ipvlw.Block) error {
	log.Printf("router %v originating %v\n", r.System, b)
	r.ControlPlane.AddRoute(&r.System, b)
	for _, router := range(r.ControlPlane.Routers()) {
		router.ControlPlane.AddRoute(&r.System, b)
	}
	return nil
}

type Computer struct {
	Addr ipvlw.Address
}

func (c *Computer) addr(a *ipvlw.Address) {
	c.Addr = *a
}

func (c *Computer) Address() ipvlw.Address {
	return c.Addr
}

func MakeNic() Nic {
	log.Printf("making a nic")
	return &Computer{ipvlw.Address{0}}
}

type RouterDhcp struct {
	Routers []*Router
}

func MakeDhcp(routers ... *Router) Dhcp {
	dhcp := RouterDhcp{make([]*Router, 0, 16)}
	return dhcp
}

func (d RouterDhcp) ConnectTo(r *Router, nics ... Nic) error {
	for _, nic := range(nics) {
		err := r.ControlPlane.AddComputer(nic)
		if err != nil {
			return err
		}
	}
	return nil
}

