package regex

import "testing"

func TestIsIp(t *testing.T) {
	ip := "1.21.34.111"
	if !isIp(ip) {
		t.Errorf("Incorrect answer. Output given false for %s", ip)
	}
	ip = "09.00.20.02"
	if isIp(ip) {
		t.Errorf("Incorrect answer. Output given true for 09.00.20.02")

	}
}

func TestFilter(t *testing.T) {
	str := "aA1_+="
	if filter(str) != "aA1_" {
		t.Errorf("Incorrect ans, expected aA1_ , got %s", filter(str))
	}

	str = "abc--    -abc"
	if filter(str) != "abcabc" {
		t.Errorf("Incorrect answer, expected abcabc , got %s", filter(str))
	}
}
