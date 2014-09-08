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

	AddNic(r *Router) error // todo rename!
	Routers() []*Router
	AddRoute(s *System, b *ipvlw.Block) error
	AddComputer(c Nic) error
	Puters() []Nic
}

type DataPlane interface {
	Runnable

	Send(ipvlw.Message) error
	Callback(func(Nic, ipvlw.Message) error) error
}

type Nic interface {
	Router() *Router
	rtr(*Router)
	Address() ipvlw.Address
	addr(a *ipvlw.Address)
	Send(ipvlw.Message) error
	Callback(func(n Nic, m ipvlw.Message) error) error
}

type Dhcp interface {
	ConnectTo(r *Router, n ... Nic) error
}




