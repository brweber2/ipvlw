package ipvlw

import (
	"testing"
	"log"
)

func TestIPvLW(t *testing.T) {
	t.Logf("Hello from the ipvlw test!\n")
}

func TestIPvLWTextMessage(t *testing.T) {
	var expectedFrom uint8 = 1
	var expectedTo uint8 = 2
	var expectedId uint64 = 3
	var expectedMessage string = "hello world message"

	tm := TextMessage{Address{expectedFrom}, Address{expectedTo}, expectedId, expectedMessage}

	// not exactly an interesting test...
	if tm.FromAddr.Address != expectedFrom {
		t.Errorf("From address not set correctly. Expected: %d Actual %d\n", expectedFrom, tm.FromAddr.Address)
	}
	if tm.ToAddr.Address != expectedTo {
		t.Errorf("To address not set correctly. Expected: %d Actual %d\n", expectedTo, tm.ToAddr.Address)
	}
	if tm.Identifier != expectedId {
		t.Errorf("Message id not set properly. Expected: %d Actual %d\n", expectedId, tm.Identifier)
	}
	if tm.Body != expectedMessage {
		t.Errorf("Message body net set properly. Expected '%s' Actual '%s'\n", expectedMessage, tm.Body)
	}
}

func TestGenerateId(t *testing.T) {
	id_1 := GenerateId()
	id_2 := GenerateId()
	id_3 := GenerateId()
	id_4 := GenerateId()
	id_5 := GenerateId()
	id_6 := GenerateId()
	id_7 := GenerateId()
	id_8 := GenerateId()
	id_9 := GenerateId()
	id_10 := GenerateId()

	set := make(map[uint64]bool)
	set[id_1] = true
	set[id_2] = true
	set[id_3] = true
	set[id_4] = true
	set[id_5] = true
	set[id_6] = true
	set[id_7] = true
	set[id_8] = true
	set[id_9] = true
	set[id_10] = true

	cnt := 0
	for k, v := range(set) {
		log.Printf("ignoring %d and %v\n", k, v)
	    cnt++
	}

	if cnt != 10 {
		t.Errorf("Did not generate 10 unique ids")
	}
}

func TestMask(t *testing.T) {
	                                      // dec   binary    cidr
	expected_0 := uint8(256 - (1<<(8-0))) // 0     00000000  /0
	expected_1 := uint8(256 - (1<<(8-1))) // 128   10000000  /1
	expected_2 := uint8(256 - (1<<(8-2))) // 192   11000000  /2
	expected_3 := uint8(256 - (1<<(8-3))) // 224   11100000  /3
	expected_4 := uint8(256 - (1<<(8-4))) // 240   11110000  /4
	expected_5 := uint8(256 - (1<<(8-5))) // 248   11111000  /5
	expected_6 := uint8(256 - (1<<(8-6))) // 252   11111100  /6
	expected_7 := uint8(256 - (1<<(8-7))) // 254   11111110  /7
	expected_8 := uint8(256 - (1<<(8-8))) // 255   11111111  /8

	mask_0 := Mask(uint8(0))
	mask_1 := Mask(uint8(1))
	mask_2 := Mask(uint8(2))
	mask_3 := Mask(uint8(3))
	mask_4 := Mask(uint8(4))
	mask_5 := Mask(uint8(5))
	mask_6 := Mask(uint8(6))
	mask_7 := Mask(uint8(7))
	mask_8 := Mask(uint8(8))

	maskExpect(0, mask_0, expected_0, t)
	maskExpect(1, mask_1, expected_1, t)
	maskExpect(2, mask_2, expected_2, t)
	maskExpect(3, mask_3, expected_3, t)
	maskExpect(4, mask_4, expected_4, t)
	maskExpect(5, mask_5, expected_5, t)
	maskExpect(6, mask_6, expected_6, t)
	maskExpect(7, mask_7, expected_7, t)
	maskExpect(8, mask_8, expected_8, t)

}

func maskExpect(i int, m, e uint8, t *testing.T) {
	log.Printf("looking at %b\n", m)
	if m != e {
		t.Errorf("mask %d is not created properly. Expected: %d Actual %d\n", i, e, m)
	}
}
