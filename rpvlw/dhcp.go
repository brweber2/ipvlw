package rpvlw

type RouterDhcp struct {
	Routers []*Router
}

func (d RouterDhcp) ConnectTo(r *Router, nics ... Nic) error {
	for _, nic := range(nics) {
		err := r.ControlPlane.AddComputer(nic)
		if err != nil {
			return err
		}
	}
	return nil
}
