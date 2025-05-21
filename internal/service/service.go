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
	flag := true
	s = strings.TrimSpace(s)
	for _, ch := range s {		
		if !strings.ContainsRune(".- ", ch){
			flag = flag && false
		}
	}

	return flag
}



func FormatData(s string) (str string, err error){
	if len(s) == 0 {
		return s, NotInvalidString
	}
	isMorse := checkStringMorse(s)

	if isMorse {
		return morse.ToText(s), nil 
	}

	f := func(r rune) bool {
		return strings.ContainsRune(".-", r)
	}
	 
	if strings.ContainsFunc(s, f) {
		return s, NotInvalidString
	}

	return morse.ToMorse(s), nil
}