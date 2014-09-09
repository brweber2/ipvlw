package rpvlw

import (
	"fmt"
	"github.com/brweber2/interwebs/ipvlw"
)

type RouterDataPlane struct {
	Router *Router
}

func (r *RouterDataPlane) Start() {
	fmt.Printf("starting data plane\n")
}

func (r *RouterDataPlane) Stop() {
	fmt.Printf("stopping data plane\n")
}

// this ignores loops, number of hops, qos, etc.
func (r *RouterDataPlane) Send(m ipvlw.Message) error {
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
