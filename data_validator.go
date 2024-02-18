package mel

import (
	"errors"
	"regexp"
	"strings"
)

var (
	WITH_FORMATTING    = 1
	WITHOUT_FORMATTING = 2
)

func ValidCpfFormat(data string, typeValidation int) bool {
	if typeValidation == WITHOUT_FORMATTING {
		return len(data) == 11
	}

	re := regexp.MustCompile(`^\d{3}\.[0-9]{3}\.[0-9]{3}-\d{2}$`)
	if len(data) == 14 && re.MatchString(data) {
		return true
	}

	return false
}

func ClenCpf(data string) (string, error) {
	if len(data) < 11 {
		return "", errors.New("invalid length to be cpf")
	}

	caracteresIndesejados := ".- "

	data = strings.Map(func(r rune) rune {
		if strings.ContainsRune(caracteresIndesejados, r) {
			return -1
		}
		return r
	}, data)

	return data, nil
}
