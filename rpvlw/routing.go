package rpvlw

import (
	"github.com/brweber2/interwebs/ipvlw"
)

type System struct {
	Identifier uint8
}

type Router struct {
	System System
	ControlPlane ControlPlane
	DataPlane DataPlane
}

type Routable interface {
	Announce(b ipvlw.Block, p RoutingPath) error
}

type Runnable interface {
	Start()
	Stop()
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
	AddRoute(p RoutingPath, b *ipvlw.Block) error
	PrintRoutes()
	AddComputer(c Nic) error
	Puters() []Nic
	GetRoutes() map[ipvlw.Block]RoutingPath
	RouteFor(a ipvlw.Address) (RoutingPath, error)
}

type DataPlane interface {
	Runnable

	Send(ipvlw.Message) error
}

type RoutingPath interface {
	Hops() []System
	Length() uint8
	First() System
	Last() System
	Clone() RoutingPath
	PrintHops()
	Pop() RoutingPath
}

type Nic interface {
	addr(a *ipvlw.Address)
	rtr(*Router)
	handler() (func(Nic,ipvlw.Message) error, error)

	Router() *Router
	Address() ipvlw.Address
	Send(ipvlw.Message) error
	RegisterCallback(func(n Nic, m ipvlw.Message) error) error
	MacAddress() string
}

type Dhcp interface {
	ConnectTo(r *Router, n ... Nic) error
}

