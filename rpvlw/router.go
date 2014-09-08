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
}

type Nic interface {
	Address() ipvlw.Address
	addr(a *ipvlw.Address)
}

type Dhcp interface {
	ConnectTo(r *Router, n ... Nic) error
}



