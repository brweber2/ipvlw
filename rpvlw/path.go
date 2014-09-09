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
