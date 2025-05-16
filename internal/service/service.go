package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var (
	NotInvalidString = errors.New("Get not invalid string")
)

func checkStringMorse(s string) bool{
	s = strings.TrimSpace(s)
	for _, ch := range s {
		if !strings.ContainsRune(".- ", ch){
			return false
		}
	}

	return true
}

func FormatData(s string) (str string, err error){
	if len(s) == 0 {
		return s, NotInvalidString
	}
	isMorse := checkStringMorse(s)

	if isMorse {
		return morse.ToText(s), nil 
	}

	return morse.ToMorse(s), nil
}