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
	log.Printf("starting control plane on router %d\n", r.Router.System.Identifier)
}

func (r *RouterControlPlane) Stop() {
	log.Printf("stopping control plane on router %d\n", r.Router.System.Identifier)
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

func (r *RouterControlPlane) RouteFor(a ipvlw.Address) (RoutingPath, error) {
	for block, routingPath := range(r.GetRoutes()) {
	    if block.Contains(a) {
			return routingPath, nil
		}
	}
	return nil, fmt.Errorf("no routing path found for address %v in router %d\n", a, r.Router.System)
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
	if _,ok := r.Addresses[*a]; ok {
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
	log.Printf("assigning computer %v ip address %d\n", n.MacAddress(), addr)
	n.rtr(r.Router)
	n.addr(addr)
	r.Addresses[*addr] = n
	r.Computers = append(r.Computers, n)
	return nil
}

func (r *RouterControlPlane) AddNic(rtr *Router) error {
	r.Nics = append(r.Nics, rtr)
	r.Interfaces[rtr.System] = *rtr
	return nil
}

func (r *RouterControlPlane) Puters() []Nic {
	return r.Computers
}

func (r *RouterControlPlane) Routers() []*Router {
	return r.Nics
}

func (r *RouterControlPlane) GetRoutes() map[ipvlw.Block]RoutingPath {
	return r.Routes
}

func (r *RouterControlPlane) PrintRoutes() {
	log.Printf("router: %d\n", r.Router.System)
	for block, path := range(r.GetRoutes()) {
		log.Printf("\tblock: %v/%d\n", block.Start, block.Bits)
		for _, hop := range(path.Hops()) {
			log.Printf("\t\thop: %d\n", hop.Identifier)
		}
	}
}

func (r *RouterControlPlane) AddRoute(rp RoutingPath, b *ipvlw.Block) error {
	if systemInPath(r, rp) {
		return nil
	}
	p := addThisSystemToPath(r, rp)
	if shortestPath(r, p, b) {
		log.Printf("shortest path %v at router %d\n", p, r.Router.System)
		if localToThisRouter(r, p) {
			r.LocalBlocks = append(r.LocalBlocks, b)
		}
		if pth, ok := r.GetRoutes()[*b]; ok {
		    log.Printf("replacing path %v with %v on router %d\n", pth, p, r.Router.System.Identifier )
		}
		r.Routes[*b] = p
		for _, neighbor := range(r.Routers()) {
			err := neighbor.Announce(*b, p.Clone())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func systemInPath(r *RouterControlPlane, p RoutingPath) bool {
	// it is ok to be first in the path, but not anywhere else (that would cause a loop)
	first_hop := true
	for _, system := range(p.Hops()) {
		if system == r.Router.System && !first_hop {
			return true
		}
		first_hop = false
	}
	return false
}

func addThisSystemToPath(r *RouterControlPlane, p RoutingPath) RoutingPath {
	return SystemPath{Systems: append(p.Hops(), r.Router.System)}
}

func shortestPath(r *RouterControlPlane, p RoutingPath, b *ipvlw.Block) bool {
	if currentPath, ok := r.Routes[*b]; ok {
		return p.Length() < currentPath.Length()
	} else {
		return true
	}
}

func localToThisRouter(r *RouterControlPlane, p RoutingPath) bool {
	return p.First() == r.Router.System
}
