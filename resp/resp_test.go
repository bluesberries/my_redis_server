package resp

import (
	"regexp"
	"testing"
)

func TestDeserializeSimpleStringOK(t *testing.T) {
	response_encoded := "+OK\r\n"
	want := regexp.MustCompile("OK")
	response_decoded, err := Deserialize(response_encoded)

	if !want.MatchString(response_decoded) || err != nil {
		t.Fatalf(`Deserialize("+OK\r\n") = %q, %v, want match for %#q, nill`, response_decoded, err, want)
	}
}

func TestDeserializeSimpleStringPING(t *testing.T) {
	response_encoded := "+PING\r\n"
	want := regexp.MustCompile("PING")
	response_decoded, err := Deserialize(response_encoded)

	if !want.MatchString(response_decoded) || err != nil {
		t.Fatalf(`Deserialize("+PING\r\n") = %q, %v, want match for %#q, nill`, response_decoded, err, want)
	}
}

func TestDeserializeSimpleStringWithoutCRLF(t *testing.T) {
	response_encoded := "+PING"
	response_decoded, err := Deserialize(response_encoded)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("+PING") = %q, %v, want match for "", error`, response_decoded, err)
	}
}

func TestDeserializeSimpleStringEndingOnlyWithLF(t *testing.T) {
	response_encoded := "+PING\n"
	response_decoded, err := Deserialize(response_encoded)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("+PING\n") = %q, %v, want match for "", error`, response_decoded, err)
	}
}

func TestDeserializeSimpleStringEndingOnlyWithCR(t *testing.T) {
	response_encoded := "+PING\r"
	response_decoded, err := Deserialize(response_encoded)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("+PING\r") = %q, %v, want match for "", error`, response_decoded, err)
	}
}

func TestDeserializeSimpleStringWithoutLeadingCharacter(t *testing.T) {
	response_encoded := "PING"
	response_decoded, err := Deserialize(response_encoded)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("PING") = %q, %v, want match for "", error`, response_decoded, err)
	}
}

func TestDeserializeSimpleStringWithNonLeadingPlus(t *testing.T) {
	response_encoded := "P+ING\r\n"
	response_decoded, err := Deserialize(response_encoded)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("PI+NG") = %q, %v, want match for "", error`, response_decoded, err)
	}
}

func TestDeserializeSimpleStringWithCRLFInString(t *testing.T) {
	response_encoded := "+O\r\nK"
	response_decoded, err := Deserialize(response_encoded)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("+O\r\nK") = %q, %v, want match for "", error`, response_decoded, err)
	}
}

func TestSerializeSimpleStringOK(t *testing.T) {
	response := "OK"
	want := regexp.MustCompile(`^(\+)(OK)(\r\n)$`)
	response_encoded, err := Serialize(response)

	if !want.MatchString(response_encoded) || err != nil {
		t.Fatalf(`Serialize("OK") = %q, %v, want match for %#q, nill`, response_encoded, err, want)
	}
}

func TestSerializeSimpleStringPING(t *testing.T) {
	response := "PING"
	want := regexp.MustCompile(`^(\+)(PING)(\r\n)$`)
	response_encoded, err := Serialize(response)

	if !want.MatchString(response_encoded) || err != nil {
		t.Fatalf(`Serialize("PING") = %q, %v, want match for %#q, nill`, response_encoded, err, want)
	}
}

func TestDeserializeBulkString(t *testing.T) {
	response := "$4\r\nPING\r\n"
	want := regexp.MustCompile("PING")

	response_decoded, err := Deserialize(response)
	if !want.MatchString(response_decoded) || err != nil {
		t.Fatalf(`Deserialize("$4\r\nPING\r\n") = %q, %v, want match for %#q, nill`, response_decoded, err, want)
	}
}

func TestDeserializeBulkStringMissingCLRF1(t *testing.T) {
	response := "$4PING\r\n"
	response_decoded, err := Deserialize(response)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("$4PING\r\n") = %q, %v, want match for "", error`, response_decoded, err)
	}
}

func TestDeserializeBulkStringMissingCLRF2(t *testing.T) {
	response := "$4\r\nPING"
	response_decoded, err := Deserialize(response)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("$4\r\nPING") = %q, %v, want match for "", error`, response_decoded, err)
	}
}

func TestDeserializeBulkStringMissingLength(t *testing.T) {
	response := "$\r\nPING\r\n"
	response_decoded, err := Deserialize(response)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("$\r\nPING\r\n") = %q, %v, want match for "", error`, response_decoded, err)
	}
}

func TestDeserializeBulkStringMissingData(t *testing.T) {
	response := "$4\r\n\r\n"
	response_decoded, err := Deserialize(response)

	if response_decoded != "" || err == nil {
		t.Fatalf(`Deserialize("$\r\n\r\n") = %q, %v, want match for "", error`, response_decoded, err)
	}
}
