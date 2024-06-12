package resp

import (
	"errors"
	"regexp"
	"strconv"
)

// Deserialize encoded string
func Deserialize(response_encoded string) (string, error) {
	if isEncodedSimpleString(response_encoded) {
		return deserializeSimpleString(response_encoded), nil
	} else if isEncodedBulkString(response_encoded) {
		return deserializeBulkString(response_encoded)
	}
	return "", errors.New("invalid response")
}

// Helper functions: Simple String
func isEncodedSimpleString(s string) bool {
	re := regexp.MustCompile(`^(\+)(\w)+(\r\n)$`)
	match := re.MatchString(s)
	return match
}

func deserializeSimpleString(response string) string {
	re := regexp.MustCompile(`\w+`)
	response_decoded := re.FindString(response)
	return response_decoded
}

// Helper functions: Bulk String
func isEncodedBulkString(s string) bool {
	re := regexp.MustCompile(`^(\$)[0-9]+(\r\n)(\w)+(\r\n)$`)
	match := re.MatchString(s)
	return match
}

func deserializeBulkString(response string) (string, error) {
	re := regexp.MustCompile(`(\S)+(\w)+(\S)+`)
	response_decoded := re.FindString(response)

	re = regexp.MustCompile(`[0-9]+`)
	response_length_str := re.FindString(response)
	response_length, err := strconv.Atoi(response_length_str)
	if err != nil {
		panic(err)
	}

	if len(response_decoded) != response_length {
		return response_decoded, errors.New("response decoded doesn't match expected length")
	}

	return response_decoded, nil
}

// Serialize string
func Serialize(response string) (string, error) {
	if isSimpleString(response) {
		return serializeSimpleString(response), nil
	}
	return "", errors.New("unable to encode")
}

func isSimpleString(s string) bool {
	re := regexp.MustCompile(`\w+`)
	match := re.MatchString(s)
	return match
}

func serializeSimpleString(response string) string {
	return "+" + response + "\r\n"
}
