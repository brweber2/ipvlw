package rpvlw

type SystemPath struct {
	Systems []System
}

func (s SystemPath) Hops() []System {
	return s.Systems
}

func (s SystemPath) Length() uint8 {
	return uint8(len(s.Systems))
}

func (s SystemPath) First() System {
	return s.Systems[0]
}

func (s SystemPath) Clone() RoutingPath {
	p := make([]System, len(s.Hops()))
	for i, hop := range(s.Hops()) {
		p[i] = System{Identifier: hop.Identifier}
	}
	return SystemPath{Systems: p}
}
