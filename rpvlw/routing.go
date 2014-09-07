package rpvlw

import (
	"fmt"
	"log"
	"github.com/brweber2/interwebs/ipvlw"
)

func MakeAndStartRouter(s int) *Router {
	controlPlane := RouterControlPlane{nil, make([]*Router, 0, 16), make(map[*ipvlw.Block]*System), make(map[*System]*Router)}
	r := Router{System{uint8(s)}, &controlPlane, RouterDataPlane{}}
	controlPlane.Router = &r
	r.Start()
	return &r
}

func (r Router) Start() {
	r.ControlPlane.Start()
	r.DataPlane.Start()
}

type RouterControlPlane struct {
	Router *Router
	Nics []*Router
	Routes map[*ipvlw.Block]*System
	Interfaces map[*System]*Router
}

func (r *RouterControlPlane) String() string {
	return fmt.Sprintf("router\n\tsystem: %v\n\trouters: %#v\n\tnics: %#v\n\troutes: %#v\n\t", r.Router.System, r.Routes, r.Nics, r.Interfaces)
}

type RouterDataPlane struct {

}

type Runnable interface {
	Start()
	Stop()
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

func (r *RouterControlPlane) Routers() []*Router {
	return r.Nics
}

func (r *RouterControlPlane) AddRoute(s *System, b *ipvlw.Block) error {
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
