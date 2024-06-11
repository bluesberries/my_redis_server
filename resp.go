package myredisserver

import (
	"errors"
	"regexp"
)

func Deserialize(response string) (string, error) {
	if isSimpleString(response) {
		return deserializeSimpleString(response), nil
	}
	return "", errors.New("invalid response")
}

func isSimpleString(s string) bool {
	re := regexp.MustCompile(`^(\+)(\w)+(\r\n)$`)
	match := re.MatchString(s)
	return match
}

func deserializeSimpleString(response string) string {
	re := regexp.MustCompile(`\w+`)
	reponse_decoded := re.FindString(response)
	return reponse_decoded
}
