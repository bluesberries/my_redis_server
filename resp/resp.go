package resp

import (
	"errors"
	"regexp"
)

// Deserialize encoded string
func Deserialize(encoded_message []byte) ([]byte, error) {
	if isEncodedSimpleString(encoded_message) {
		return deserializeSimpleString(encoded_message), nil
	} else if isEncodedBulkString(encoded_message) {
		return deserializeBulkString(encoded_message)
	}
	return nil, errors.New("invalid response")
}

// Helper functions: Simple String
func isEncodedSimpleString(encoded_message []byte) bool {
	re := regexp.MustCompile(`^(\+)(\w)+(\r\n)$`)
	match := re.Match(encoded_message)
	return match
}

func deserializeSimpleString(encoded_message []byte) []byte {
	re := regexp.MustCompile(`\w+`)
	response_decoded := re.Find(encoded_message)
	return response_decoded
}

// Helper functions: Bulk String
func isEncodedBulkString(encoded_message []byte) bool {
	re := regexp.MustCompile(`(\$)[0-9]+(\r\n)(\w)+(\r\n)$`)
	match := re.Match(encoded_message)
	return match
}

func deserializeBulkString(encoded_message []byte) ([]byte, error) {
	re := regexp.MustCompile(`(\S)+(\w)+(\S)+`)
	response_decoded := re.Find(encoded_message)

	re = regexp.MustCompile(`[0-9]+`)
	//response_length_str := re.Find(encoded_message)
	/*
		response_length, err := strconv.Atoi(response_length_str)
		if err != nil {
			panic(err)
		}

		if len(response_decoded) != response_length {
			return response_decoded, errors.New("response decoded doesn't match expected length")
		}
	*/

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
