package rpvlw

import (
	"github.com/brweber2/interwebs/ipvlw"
	"log"
)


func (r Router) Start() {
	r.ControlPlane.Start()
	r.DataPlane.Start()
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

func (r Router) Originate(b *ipvlw.Block) error {
	log.Printf("router %v originating %v\n", r.System, b)
	routingPath := SystemPath{Systems: []System{r.System}}

	r.ControlPlane.AddRoute(routingPath, b)
	log.Printf("announce to routers: %v\n", r.ControlPlane.Routers())
	for _, router := range(r.ControlPlane.Routers()) {
		log.Printf("announce %v out of %v\n", b, r.System)
		router.ControlPlane.AddRoute(routingPath, b)
	}
	return nil
}





