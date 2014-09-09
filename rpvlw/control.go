package rpvlw

import (
	"github.com/brweber2/interwebs/ipvlw"
	"fmt"
	"log"
)

type RouterControlPlane struct {
	Router *Router
	// local to this network
	LocalBlocks []*ipvlw.Block
	Computers []Nic
	Addresses map[ipvlw.Address]Nic
	// external to this network
	Nics []*Router // todo rename me!
	Routes map[ipvlw.Block]RoutingPath
	Interfaces map[System]Router
}

func (r *RouterControlPlane) Start() {
	fmt.Printf("starting control plane\n")
}

func (r *RouterControlPlane) Stop() {
	fmt.Printf("stopping control plane\n")
}

func (r *RouterControlPlane) isLocal(a ipvlw.Address) bool {
	for _, block := range(r.LocalBlocks) {
		if block.Contains(a) {
			log.Printf("%v is local to %v\n", a, r)
			return true
		}
	}
	return false
}

func (r *RouterControlPlane) nicFor(a ipvlw.Address) (Nic,error) {
	if nic, ok := r.Addresses[a]; ok {
		return nic, nil
	}
	return nil, fmt.Errorf("Unable to find NIC for %v\n", a)
}

func (r *RouterControlPlane) routeFor(a ipvlw.Address) bool {
	log.Printf("looking for route to %v in %v\n", a, r.Routes)
	for block, _ := range(r.Routes) {
		log.Printf("checking if block %v contains %v\n", block, a)
		if block.Contains(a) {
			return true
		}
	}
	return false
}

func (r *RouterControlPlane) systemFor(a ipvlw.Address) (System, error) {
	for block, routingPath := range(r.Routes) {
		if block.Contains(a) {
			return routingPath.First(), nil
		}
	}
	return System{}, fmt.Errorf("Unable to find system for %v\n", a)
}

func (r *RouterControlPlane) routerFor(s System) (Router, error) {
	if router, ok := r.Interfaces[s]; ok {
		return router, nil
	}
	return Router{}, fmt.Errorf("Unable to find router for %v\n", s)
}

func (r *RouterControlPlane) String() string {
	return fmt.Sprintf("router\n\tsystem: %v\n\trouters: %#v\n\tnics: %#v\n\troutes: %#v\n\t", r.Router.System, r.Routes, r.Nics, r.Interfaces)
}

func (r *RouterControlPlane) AddressInUse(a *ipvlw.Address) bool {
	log.Printf("checking if %v is in %v\n", a, r.Addresses)
	if _,ok := r.Addresses[*a]; ok {
		return true
	}
	return false
}

func (r *RouterControlPlane) UnusedAddress() (*ipvlw.Address, error) {
	for _, block := range(r.LocalBlocks) {
		for _, addr := range(block.Addresses()) {
			if ! r.AddressInUse(addr) {
				log.Printf("returning unused address %v\n", addr)
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
	r.Addresses[*addr] = n
	r.Computers = append(r.Computers, n)
	return nil
}

func (r *RouterControlPlane) AddNic(rtr *Router) error {
	log.Printf("adding router %v to %v\n", rtr, r.Nics)
	r.Nics = append(r.Nics, rtr)
	log.Printf("added router %v to %v\n", rtr, r.Nics)
	r.Interfaces[r.Router.System] = *rtr
	return nil
}

func (r *RouterControlPlane) Puters() []Nic {
	return r.Computers
}

func (r *RouterControlPlane) Routers() []*Router {
	return r.Nics
}

func (r *RouterControlPlane) AddRoute(p RoutingPath, b *ipvlw.Block) error {
	if p.First().Identifier == r.Router.System.Identifier {
		r.LocalBlocks = append(r.LocalBlocks, b)
	}
	r.Routes[*b] = p
	return nil
}
