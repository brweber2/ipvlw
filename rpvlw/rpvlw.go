package rpvlw

//import (
//	"github.com/brweber2/interwebs/ipvlw"
//	"fmt"
//)



//type Announcement struct {
//	System System
//	Blocks []ipvlw.Block
//}
//
//func (a Announcement) String() string {
//	return fmt.Sprintf("%d:%v", a.System.Identifier, a.Blocks)
//}
//
//type AnnouncementMessage struct {
//	FromAddr ipvlw.Address
//	ToAddr ipvlw.Address
//	Identifier uint64
//	Announcement Announcement
//}
//
//func (t AnnouncementMessage) From() ipvlw.Address {
//return t.FromAddr
//}
//
//func (t AnnouncementMessage) To() ipvlw.Address {
//	return t.ToAddr
//}
//
//func (t AnnouncementMessage) Id() uint64 {
//	return t.Identifier
//}
//
//func (t AnnouncementMessage) Payload() string {
//	return t.Announcement.String()
//}
