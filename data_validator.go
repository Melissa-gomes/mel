package mel

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var (
	WITH_FORMATTING    = 1
	WITHOUT_FORMATTING = 2
)

func invalidCharacter(data string) bool {
	hasNonDigit := false

	for _, char := range data {
		if !unicode.IsDigit(char) {
			hasNonDigit = true
			break
		}
	}

	return hasNonDigit
}

func ValidCpfFormat(data string, typeValidation int) bool {
	if typeValidation == WITHOUT_FORMATTING {
		if invalidCharacter(data) {
			return false
		}
		return len(data) == 11
	}

	re := regexp.MustCompile(`^\d{3}\.[0-9]{3}\.[0-9]{3}-\d{2}$`)
	if len(data) == 14 && re.MatchString(data) {
		return true
	}

	return false
}

func CleanCpf(data string) (string, error) {
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

func FormatCpf(data string) (string, error) {
	if len(data) > 11 {
		return "", errors.New("invalid length to be cpf")
	}

	// if invalidCharacter(data) {
	// 	return "", errors.New("invalid character in cpf")
	// }

	re := regexp.MustCompile(`^(\d{3})(\d{3})(\d{3})(\d{2})$`)
	matches := re.FindStringSubmatch(data)
	cpfFormatted := fmt.Sprintf("%s.%s.%s-%s", matches[1], matches[2], matches[3], matches[4])

	return cpfFormatted, nil
}
