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
	// todo test outside valid range (ie 0 > n > 8)

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

func TestBlockContainsAddress(t *testing.T) {
	// every byte is valid in a slash 0
	slash_0 := Block{Address{0}, 0}
	EnsureBlockContains(slash_0, 0, t)
	EnsureBlockContains(slash_0, 1, t)
	EnsureBlockContains(slash_0, 2, t)
	EnsureBlockContains(slash_0, 3, t)
	// ...
	EnsureBlockContains(slash_0, 127, t)
	EnsureBlockContains(slash_0, 128, t)
	EnsureBlockContains(slash_0, 129, t)

	// ...
	EnsureBlockContains(slash_0, 253, t)
	EnsureBlockContains(slash_0, 254, t)
	EnsureBlockContains(slash_0, 255, t)

	slash_8 := Block{Address{17}, 8}

	EnsureBlockDoesNOTContain(slash_8, 0, t)
	EnsureBlockDoesNOTContain(slash_8, 1, t)
	EnsureBlockDoesNOTContain(slash_8, 2, t)
	// ...
	EnsureBlockDoesNOTContain(slash_8, 15, t)
	EnsureBlockDoesNOTContain(slash_8, 16, t)
	EnsureBlockContains(slash_8, 17, t)
	EnsureBlockDoesNOTContain(slash_8, 18, t)
	EnsureBlockDoesNOTContain(slash_8, 19, t)
	// ...
	EnsureBlockDoesNOTContain(slash_8, 254, t)
	EnsureBlockDoesNOTContain(slash_8, 255, t)

	slash_7 := Block{Address{8}, 7}

	EnsureBlockDoesNOTContain(slash_7, 0, t)
	EnsureBlockDoesNOTContain(slash_7, 1, t)
	EnsureBlockDoesNOTContain(slash_7, 2, t)
	// ...
	EnsureBlockDoesNOTContain(slash_7, 6, t)
	EnsureBlockDoesNOTContain(slash_7, 7, t)
	EnsureBlockContains(slash_7, 8, t)
	EnsureBlockContains(slash_7, 9, t)
	EnsureBlockDoesNOTContain(slash_7, 10, t)
	EnsureBlockDoesNOTContain(slash_7, 11, t)
	// ...
	EnsureBlockDoesNOTContain(slash_7, 254, t)
	EnsureBlockDoesNOTContain(slash_7, 255, t)

	slash_5 := Block{Address{168}, 5}

	EnsureBlockDoesNOTContain(slash_5, 0, t)
	EnsureBlockDoesNOTContain(slash_5, 1, t)
	EnsureBlockDoesNOTContain(slash_5, 2, t)
	// ...
	EnsureBlockDoesNOTContain(slash_5, 164, t)
	EnsureBlockDoesNOTContain(slash_5, 165, t)
	EnsureBlockDoesNOTContain(slash_5, 166, t)
	EnsureBlockDoesNOTContain(slash_5, 167, t)
	EnsureBlockContains(slash_5, 168, t)
	EnsureBlockContains(slash_5, 169, t)
	EnsureBlockContains(slash_5, 170, t)
	EnsureBlockContains(slash_5, 171, t)
	EnsureBlockContains(slash_5, 172, t)
	EnsureBlockContains(slash_5, 173, t)
	EnsureBlockContains(slash_5, 174, t)
	EnsureBlockContains(slash_5, 175, t)
	EnsureBlockDoesNOTContain(slash_5, 176, t)
	EnsureBlockDoesNOTContain(slash_5, 177, t)
	EnsureBlockDoesNOTContain(slash_5, 178, t)
	// ...
	EnsureBlockDoesNOTContain(slash_5, 254, t)
	EnsureBlockDoesNOTContain(slash_5, 255, t)
}

func EnsureBlockContains(b Block, a int, t *testing.T) {
	if ! b.Contains(Address{uint8(a)}) {
		t.Errorf("Block %v should contain %v\n", b, a)
	}
}

func EnsureBlockDoesNOTContain(b Block, a int, t *testing.T) {
	if b.Contains(Address{uint8(a)}) {
		t.Errorf("Block %v should NOT contain %v\n", b, a)
	}
}

func TestMakeResponse(t *testing.T) {
	// MESSAGES
	valid := TextMessage{Address{1}, Address{2}, 3, "valid"}

	no_from := new(TextMessage)
	no_from.ToAddr = Address{2}
	no_from.Identifier = 3
	no_from.Body = "no from"

	no_to := new(TextMessage)
	no_to.FromAddr = Address{1}
	no_to.Identifier = 3
	no_to.Body = "no to"

	no_id := new(TextMessage)
	no_id.FromAddr = Address{1}
	no_id.ToAddr = Address{2}
	no_id.Body = "no to"

	no_message := new(TextMessage)
	no_message.FromAddr = Address{1}
	no_message.ToAddr = Address{2}
	no_message.Identifier = 3

	empty_message := TextMessage{Address{1}, Address{2}, 3, ""}

	// RESPONSES
	valid_response := TextMessage{Address{2}, Address{1}, 3, "valid response"}
	no_message_response := TextMessage{Address{2}, Address{1}, 3, "no message response"}
	empty_message_response := TextMessage{Address{2}, Address{1}, 3, "empty message response"}

	// all filled out
	r_valid, err := MakeResponse(valid, "valid response")
	if err != nil {
		t.Errorf("Did not make proper response for a valid message. Error %v\n", err)
	}
	if r_valid != valid_response {
		t.Errorf("Did not make proper response for a valid message. Expected %v Actual %v\n", valid_response, r_valid)
	}
	log.Printf("r_valid: %v\n", r_valid)

	// empty message
	r_empty_message, err := MakeResponse(empty_message, "empty message response")
	if err != nil {
		t.Errorf("Did not make proper response for an 'empty message' message. Error %v\n", err)
	}
	if r_empty_message != empty_message_response {
		t.Errorf("Did not make proper response for an 'empty message' message. Expected %v Actual %v\n", empty_message_response, r_empty_message)
	}

	// no message
	r_no_message, err := MakeResponse(no_message, "no message response")
	if err != nil {
		t.Errorf("Did not make proper response for a 'no message' message. Error %v\n", err)
	}
	if r_no_message != no_message_response {
		t.Errorf("Did not make proper response for a 'no message' message. Expected %v Actual %v\n", no_message_response, r_no_message)
	}

//	// no from
//	_, err = MakeResponse(no_from, "ignored")
//	if err == nil {
//		t.Errorf("Did not get an error when making a response to a message with no from. orig: %v\n", no_from)
//	}
//
//	// no to
//	_, err = MakeResponse(no_to, "ignored")
//	if err == nil {
//		t.Errorf("Did not get an error when making a response to a message with no to.\n")
//	}
//
//	// no id
//	_, err = MakeResponse(no_id, "ignored")
//	if err == nil {
//		t.Errorf("Did not get an error when making a response to a message with no id.\n")
//	}
}

func TestBlockAddresses(t *testing.T) {
	block := Block{Address{148}, 6}
	addrs := block.Addresses()
	for _, addr := range(addrs) {
		log.Printf("block contains %v\n", addr)
	}
}

func TestBlockPointerAddresses(t *testing.T) {
	block := &Block{Address{148}, 6}
	addrs := block.Addresses()
	for _, addr := range(addrs) {
		log.Printf("block contains %v\n", addr)
	}
}
