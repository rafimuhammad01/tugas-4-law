package utils

import "strings"

func ErrorsParser(err error) string {
	msg := strings.Split(err.Error(), ":")

	return msg[0]
}
