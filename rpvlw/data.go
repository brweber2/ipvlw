package rpvlw

import (
	"fmt"
	"github.com/brweber2/interwebs/ipvlw"
	"log"
)

type RouterDataPlane struct {
	Router *Router
}

func (r *RouterDataPlane) Start() {
	log.Printf("starting data plane on router %d\n", r.Router.System.Identifier)
}

func (r *RouterDataPlane) Stop() {
	log.Printf("stopping data plane on router %d\n", r.Router.System.Identifier)
}

// this ignores loops, number of hops, qos, etc.
func (r *RouterDataPlane) Send(m ipvlw.Message) error {
	// todo we should add a hop to the message here!
	to := m.To()
	if r.Router.ControlPlane.isLocal(to) {
		log.Printf("sending message via local traffic to %v on router %d\n", to, r.Router.System.Identifier)
		// if to is local, send to nic
		targetNic, err := r.Router.ControlPlane.nicFor(to)
		if err != nil {
			return err
		}
		f, err := targetNic.handler()
		if err != nil {
			return err
		}
//		log.Printf("going to call handler... sending %v\n", m)
		return f(targetNic, m)
	} else {
		log.Printf("sending message via external traffic to %v on router %d\n", to, r.Router.System.Identifier)
		systemPath, err := r.Router.ControlPlane.RouteFor(to)
		if err != nil {
			return err
		}
//		log.Printf("going to route this message via following path:\n")
		systemPath = systemPath.Pop()
//		systemPath.PrintHops()
		nextHop := systemPath.Last()
		log.Printf("next hop is %v on router %d\n", nextHop, r.Router.System.Identifier)
		if nextHop.Identifier == r.Router.System.Identifier {
			log.Printf("route should be local, about to bomb! router: %d\n", r.Router.System)
			return fmt.Errorf("This route should be local... %v on router %d\n", nextHop, r.Router.System)
		}
		nextRouter, err := r.Router.ControlPlane.routerFor(nextHop)
		if err != nil {
//			log.Printf("--- bbw ---\n")
			log.Print(err)
			return err
		}
//		log.Printf("going to send traffic via router %v from router %d\n", nextRouter, r.Router.System.Identifier)
		return nextRouter.DataPlane.Send(m)
	}
}
