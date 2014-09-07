package rpvlw

import (
	"testing"
	"github.com/brweber2/interwebs/ipvlw"
	"log"
)

func TestAnnouncementMessage(t *testing.T) {
	expectedFrom := ipvlw.Address{1}
	expectedTo := ipvlw.Address{2}
	expectedId := uint64(3)
	expectedBody := "ack"

	m := AnnouncementMessage{expectedFrom, expectedTo, expectedId,
		Announcement{System{4}, []ipvlw.Block{ipvlw.Block{ipvlw.Address{16}, 7}}}}

	log.Printf("annoucement: %v\n", m)

	resp, err := ipvlw.MakeResponse(m, expectedBody)
	if err != nil {
		t.Errorf("failed to make an annoucement response with error %v\n", err)
	}
	if resp.From() != expectedFrom {
		t.Errorf("failed to make an annoucement response with bad from address\n")
	}
	if resp.To() != expectedTo {
		t.Errorf("failed to make an annoucement response with bad to address\n")
	}
	if resp.Id() != expectedId {
		t.Errorf("failed to make an annoucement response with bad mesage id\n")
	}
	if resp.Payload() != expectedBody {
		t.Errorf("failed to make an annoucement response with bad response message\n")
	}
}
