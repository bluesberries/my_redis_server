package myredisserver

import (
	"errors"
	"regexp"
)

func Deserialize(response_encoded string) (string, error) {
	if isEncodedSimpleString(response_encoded) {
		return deserializeSimpleString(response_encoded), nil
	}
	return "", errors.New("invalid response")
}

func isEncodedSimpleString(s string) bool {
	re := regexp.MustCompile(`^(\+)(\w)+(\r\n)$`)
	match := re.MatchString(s)
	return match
}

func deserializeSimpleString(response string) string {
	re := regexp.MustCompile(`\w+`)
	reponse_decoded := re.FindString(response)
	return reponse_decoded
}

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
