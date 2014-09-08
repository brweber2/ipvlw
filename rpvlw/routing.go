package rpvlw

import (
	"fmt"
	"log"
	"github.com/brweber2/interwebs/ipvlw"
)

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

type RouterDataPlane struct {
	Router *Router
}

type Computer struct {
	Rtr *Router
	Addr ipvlw.Address
	Callback func(Nic, ipvlw.Message) error
}

type RouterDhcp struct {
	Routers []*Router
}

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

	r := Router{System{uint8(s)}, &controlPlane, dataPlane}
	controlPlane.Router = &r
	dataPlane.Router = &r

	r.Start()
	return &r
}

func (r Router) Start() {
	r.ControlPlane.Start()
	r.DataPlane.Start()
}

func (r *RouterControlPlane) isLocal(a ipvlw.Address) bool {
	for _, block := range(r.LocalBlocks) {
		if block.Contains(a) {
			return true
		}
	}
	return false
}

func (r *RouterControlPlane) nicFor(a ipvlw.Address) (Nic,error) {
	if nic, ok := r.Addresses[&a]; ok {
		return nic, nil
	}
	return nil, fmt.Errorf("Unable to find NIC for %v\n", a)
}

func (r *RouterControlPlane) routeFor(a ipvlw.Address) bool {
	for block, _ := range(r.Routes) {
		if block.Contains(a) {
			return true
		}
	}
	return false
}

func (r *RouterControlPlane) systemFor(a ipvlw.Address) (System, error) {
	for block, system := range(r.Routes) {
		if block.Contains(a) {
			return *system, nil
		}
	}
	return System{}, fmt.Errorf("Unable to find system for %v\n", a)
}

func (r *RouterControlPlane) routerFor(s System) (Router, error) {
	if router, ok := r.Interfaces[&s]; ok {
		return *router, nil
	}
	return Router{}, fmt.Errorf("Unable to find router for %v\n", s)
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
	n.rtr(r.Router)
	n.addr(addr)
	r.Addresses[addr] = n
	r.Computers = append(r.Computers, n)
	return nil
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

func (c *Computer) addr(a *ipvlw.Address) {
	c.Addr = *a
}

func (c *Computer) Address() ipvlw.Address {
	return c.Addr
}

func (c *Computer) Router() *Router {
	return c.Rtr
}

func (c *Computer) rtr(r *Router) {
	c.Rtr = r
}

func (c *Computer) Send(m ipvlw.Message) error {
	return c.Router().DataPlane.Send(m)
}

func (c *Computer) RegisterCallback(f func(n Nic, m ipvlw.Message) error) error {
	c.Callback = f
	return nil
}

func (c *Computer) handler() (func(Nic,ipvlw.Message) error, error) {
	if c.Callback != nil {
		return c.Callback, nil
	}
	return nil, fmt.Errorf("No callback registered!\n")
}

func MakeNic() Nic {
	log.Printf("making a nic")
	return &Computer{nil, ipvlw.Address{0}, nil}
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

// this ignores loops, number of hops, qos, etc.
func (r RouterDataPlane) Send(m ipvlw.Message) error {
	// todo we should add a hop to the message here!
	to := m.To()
	if r.Router.ControlPlane.isLocal(to) {
		// if to is local, send to nic
		targetNic, err := r.Router.ControlPlane.nicFor(to)
		if err != nil {
			return err
		}
		f, err := targetNic.handler()
		if err != nil {
			return err
		}
		return f(targetNic, m)
	} else if r.Router.ControlPlane.routeFor(to) {
		// if to is external, send to router
		targetSystem, err := r.Router.ControlPlane.systemFor(to)
		if err != nil {
			return err
		}
		targetRouter, err := r.Router.ControlPlane.routerFor(targetSystem)
		if err != nil {
			return err
		}
		return targetRouter.DataPlane.Send(m)
	} else {
		return fmt.Errorf("Unable to find router for %v\n", to)
	}
}


