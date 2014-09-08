package rpvlw

import (
	"github.com/brweber2/interwebs/ipvlw"
	"fmt"
)

type Computer struct {
	Rtr *Router
	Addr ipvlw.Address
	Callback func(Nic, ipvlw.Message) error
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
