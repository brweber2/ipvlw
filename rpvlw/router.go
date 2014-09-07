package rpvlw

import "github.com/brweber2/interwebs/ipvlw"

type Router struct {
	System System
	ControlPlane ControlPlane
	DataPlane DataPlane
}

type RouteDatabase struct {
	Routes map[ipvlw.Block]System
	Interfaces map[System]Nic
	Nics []Nic
}

type PolicyDatabase struct {
	Policies []Policy
}

// **** behavior ****

type Policy interface {
	Check(a Announcement) bool
}

type Announcable interface {
	Announce(a Announcement) error
	Listen() ([]Announcement, error) // these may be accepted or rejected
	Announcements() ([]Announcement, error)
}

type Fiterable interface {
	AddPolicy(p Policy) error
	RemovePolicy(p Policy) error
	Policies() ([]Policy, error)
}

type Routable interface {
	Connect(n Nic) error
	Disconnect(n Nic) error
	ConnectedRouters() ([]Router, error)
}

type ControlPlane interface {
	Routable



}

type DataPlane interface {
	Start() error
	Stop() error
}

type Nic interface {
	Address() ipvlw.Address
	Send(m ipvlw.Message) (ipvlw.Message, error)
	Listen() ([]ipvlw.Message, error)
}

type Dhcp interface {
	Register(n Nic) (Router, ipvlw.Address, error)
}
