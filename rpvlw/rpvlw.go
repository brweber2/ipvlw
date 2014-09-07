package rpvlw

import (
	"github.com/brweber2/interwebs/ipvlw"
	"fmt"
)

type System struct {
	Identifier uint8
}

type Announcement struct {
	System System
	Blocks []ipvlw.Block
}

func (a Announcement) String() string {
	return fmt.Sprintf("%d:%v", a.System.Identifier, a.Blocks)
}

type AnnoucementMessage struct {
	FromAddr ipvlw.Address
	ToAddr ipvlw.Address
	Announcement Announcement
}

func (t AnnoucementMessage) From() ipvlw.Address {
return t.FromAddr
}

func (t AnnoucementMessage) To() ipvlw.Address {
	return t.ToAddr
}

func (t AnnoucementMessage) Payload() string {
	return t.Announcement.String()
}
