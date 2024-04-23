package main

import (
	"fmt"
	"github.com/rainycape/unidecode"
	"github.com/ttacon/libphonenumber"
	"net/mail"
	"strings"
	"unicode"
)

type Hasher interface {
	ProcessStringIntoHash(string, ExtraDataHasher) (string, error)
}

type DefaultHasher struct {
	Hasher
}

func NewDefaultHasher() *DefaultHasher {
	return &DefaultHasher{}
}

func (h *DefaultHasher) ProcessStringIntoHash(_ string, _ ExtraDataHasher) (string, error) {
	panic("implement method")
}

func (h *DefaultHasher) removeAllNonASCII(str string) string {
	return unidecode.Unidecode(str)
}

func (h *DefaultHasher) toLower(str string) string {
	return strings.ToLower(str)
}

type EmailHasher struct {
	*DefaultHasher
	Hasher
}

func NewEmailHasher() *EmailHasher {
	return &EmailHasher{
		DefaultHasher: NewDefaultHasher(),
	}
}

func (h *EmailHasher) ProcessStringIntoHash(str string, _ ExtraDataHasher) (string, error) {
	str = h.removeAllNonASCII(str)
	str = h.toLower(str)

	addr, err := mail.ParseAddress(str)
	if err == nil {
		str = EmailRegex.FindString(addr.Address)
	}

	str = MoreThanOneDotRegex.ReplaceAllString(str, ".")
	str = MoreThanOneAtRegex.ReplaceAllString(str, "@")

	addr, err = mail.ParseAddress(str)
	if err == nil {
		str = EmailRegex.FindString(addr.Address)
	}

	str = AllWhiteSpacesRegex.ReplaceAllString(str, "")

	if len(str) == 0 {
		return "", fmt.Errorf("invalid email address")
	}

	str = DisallowedCharactersEmailRegex.ReplaceAllString(str, "")

	str = strings.TrimFunc(str, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	addr, err = mail.ParseAddress(str)
	if err != nil {
		return "", fmt.Errorf("invalid email address: %v", err)
	}

	str = ExtraDataEmailRegex.ReplaceAllString(addr.Address, "@")

	return EmailRegex.FindString(str), nil
}

type PhoneHasher struct {
	*DefaultHasher
	Hasher
}

func NewPhoneHasher() *PhoneHasher {
	return &PhoneHasher{
		DefaultHasher: NewDefaultHasher(),
	}
}

func (h *PhoneHasher) ProcessStringIntoHash(str string, extra ExtraDataHasher) (string, error) {
	str = AllWhiteSpacesRegex.ReplaceAllString(str, "")
	str = AllNonNumericRegex.ReplaceAllString(str, "")

	if len(str) == 0 {
		return "", fmt.Errorf("invalid phone number")
	}

	num, err := libphonenumber.Parse(str, fmt.Sprint(extra["country_code"]))
	if err != nil {
		return "", fmt.Errorf("error parsing phone number: %v", err)
	}

	if libphonenumber.IsPossibleNumber(num) == false {
		return "", fmt.Errorf("phone number is invalid")
	}

	str = fmt.Sprintf("%d", num.GetNationalNumber())

	if len(str) > 15 {
		return "", fmt.Errorf("phone number is to long: %s", str)
	}

	return str, nil
}

type NameHasher struct {
	*DefaultHasher
	Hasher
}

func NewNameHasher() *NameHasher {
	return &NameHasher{
		DefaultHasher: NewDefaultHasher(),
	}
}

func (h *NameHasher) ProcessStringIntoHash(str string, _ ExtraDataHasher) (string, error) {
	str = AllWhiteSpacesRegex.ReplaceAllString(str, "")
	str = h.removeAllNonASCII(str)
	str = h.toLower(str)
	str = NonAlphanumericRegex.ReplaceAllString(str, "")

	if len(str) == 0 {
		return "", fmt.Errorf("str is empty")
	}

	return str, nil
}
