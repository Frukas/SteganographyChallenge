package steganography

import "testing"

func TestFlipLastByte(t *testing.T) {
	test := []uint8{1, 0, 255, 254, 25, 24}
	testResult := []uint8{0, 1, 254, 255, 24, 25}
	for index, ts := range test {
		rs := flipLastByte(ts)

		if rs != testResult[index] {
			t.Error("Expected:", testResult[index], "Got:", rs, "For Test Value:", ts)
		}
	}
}

func TestGetLastByte(t *testing.T) {
	test := []uint8{1, 0, 255, 254, 25, 24}
	testResult := []uint8{1, 0, 1, 0, 1, 0}
	for index, ts := range test {
		rs := getLastByte(ts)

		if rs != testResult[index] {
			t.Error("Expected:", testResult[index], "Got:", rs, "For test value:", ts)
		}
	}
}

func TestGetIndexByteAsLastByte(t *testing.T) {
	test := []uint8{170, 255, 0, 129}
	testResult := [][]uint8{{1, 0, 1, 0, 1, 0, 1, 0}, {1, 1, 1, 1, 1, 1, 1, 1}, {0, 0, 0, 0, 0, 0, 0, 0}, {1, 0, 0, 0, 0, 0, 0, 1}}

	for index, ts := range test {
		for i := 0; i < 8; i++ {
			rs := getIndexByteAsLastByte(ts, i)

			if rs != testResult[index][i] {
				t.Error("Expected:", testResult[index][i], "Got:", rs, "For test value:t", ts)
			}
		}
	}
}

func TestGetBinaryChangeList(t *testing.T) {
	test := []uint8{170, 255, 0, 129}
	testResult := [][]uint8{{1, 0, 1, 0, 1, 0, 1, 0}, {1, 1, 1, 1, 1, 1, 1, 1}, {0, 0, 0, 0, 0, 0, 0, 0}, {1, 0, 0, 0, 0, 0, 0, 1}}

	for index, ts := range test {
		rs := getBinaryChangeList(rune(ts))

		if !isEqualSliceUINT8(rs, testResult[index]) {
			t.Error("Expected", testResult[index], "Got", rs, "For the Value", ts)
		}
	}
}

func isEqualSliceUINT8(slice1, slice2 []uint8) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for index, s := range slice1 {
		if s != slice2[index] {
			return false
		}
	}
	return true
}
