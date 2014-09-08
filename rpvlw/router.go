package rpvlw

import (
	"github.com/brweber2/interwebs/ipvlw"
)

type Router struct {
	System System
	ControlPlane ControlPlane
	DataPlane DataPlane
}

type ControlPlane interface {
	Runnable

	isLocal(a ipvlw.Address) bool
	nicFor(a ipvlw.Address) (Nic, error)
	routeFor(a ipvlw.Address) bool
	systemFor(a ipvlw.Address) (System, error)
	routerFor(s System) (Router, error)

	AddNic(r *Router) error // todo rename!
	Routers() []*Router
	AddRoute(s *System, b *ipvlw.Block) error
	AddComputer(c Nic) error
	Puters() []Nic
}

type DataPlane interface {
	Runnable

	Send(ipvlw.Message) error
}

type Nic interface {
	addr(a *ipvlw.Address)
	rtr(*Router)
	handler() (func(Nic,ipvlw.Message) error, error)

	Router() *Router
	Address() ipvlw.Address
	Send(ipvlw.Message) error
	RegisterCallback(func(n Nic, m ipvlw.Message) error) error
}

type Dhcp interface {
	ConnectTo(r *Router, n ... Nic) error
}




