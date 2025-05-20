package service

import (
	"errors"
	"strings"

	"go1fl-sprint6-final/pkg/morse"
)

var (
	NotInvalidString = errors.New("Get not invalid string")
)

func checkStringMorse(s string) bool{
	s = strings.TrimSpace(s)
	f := func(r rune) bool {
		return strings.ContainsRune(".- ", r)
	}
	return strings.ContainsFunc(s, f)
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