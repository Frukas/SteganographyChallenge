package steganography

import "testing"

func TestByteToString(t *testing.T) {

	test := []byte("Hello This is a test.")
	testResult := "Hello This is a test."

	rs := byteToString(test)

	if rs != testResult {
		t.Error("Expected:", testResult, "Got:", rs)
	}
}

func TestWriteListOnByte(t *testing.T) {
	testSet := [][]uint8{{1, 0, 1, 0, 1, 0, 1, 0}, {1, 1, 1, 1, 1, 1, 1, 1}, {0, 0, 0, 0, 0, 0, 0, 0}, {1, 0, 0, 0, 0, 0, 0, 1}}
	testResult := []byte{170, 255, 0, 129}

	for index, tr := range testResult {
		var rs byte
		writeListOnByte(&rs, testSet[index])

		if rs != tr {
			t.Error("Expected:", tr, "Got:", rs)
		}
	}
}
