package myredisserver

import (
	"regexp"
	"testing"
)

func TestSimpleStringOK(t *testing.T) {
	response_encoded := "+OK\r\n"
	want := regexp.MustCompile("OK")
	response_decoded, err := Deserialize(response_encoded)

	if !want.MatchString(response_decoded) || err != nil {
		t.Fatalf(`Deserialize("+OK\r\n") = %q, %v, want match for %#q, nill`, response_decoded, err, want)
	}
}

func TestSimpleStringPING(t *testing.T) {
	response_encoded := "+PING\r\n"
	want := regexp.MustCompile("PING")
	response_decoded, err := Deserialize(response_encoded)

	if !want.MatchString(response_decoded) || err != nil {
		t.Fatalf(`Deserialize("+PING\r\n") = %q, %v, want match for %#q, nill`, response_decoded, err, want)
	}
}

func TestSimpleStringWithoutCRLF(t *testing.T) {
	response_encoded := "+PING"
	response_decoded, err := Deserialize(response_encoded)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("+PING") = %q, %v, want match for "", error`, response_decoded, err)
	}
}

func TestSimpleStringEndingOnlyWithLF(t *testing.T) {
	response_encoded := "+PING\n"
	response_decoded, err := Deserialize(response_encoded)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("+PING\n") = %q, %v, want match for "", error`, response_decoded, err)
	}
}

func TestSimpleStringEndingOnlyWithCR(t *testing.T) {
	response_encoded := "+PING\r"
	response_decoded, err := Deserialize(response_encoded)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("+PING\r") = %q, %v, want match for "", error`, response_decoded, err)
	}
}

func TestSimpleStringWithoutLeadingCharacter(t *testing.T) {
	response_encoded := "PING"
	response_decoded, err := Deserialize(response_encoded)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("PING") = %q, %v, want match for "", error`, response_decoded, err)
	}
}

func TestSimpleStringWithNonLeadingPlus(t *testing.T) {
	response_encoded := "P+ING\r\n"
	response_decoded, err := Deserialize(response_encoded)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("PI+NG") = %q, %v, want match for "", error`, response_decoded, err)
	}
}

func TestSimpleStringWithCRLFInString(t *testing.T) {
	response_encoded := "+O\r\nK"
	response_decoded, err := Deserialize(response_encoded)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("+O\r\nK") = %q, %v, want match for "", error`, response_decoded, err)
	}
}
