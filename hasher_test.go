package main

import (
	"fmt"
	"testing"
)

func TestEmailHasher_ProcessStringIntoHash_ValidEmail(t *testing.T) {
	h := NewEmailHasher()

	_, err := h.ProcessStringIntoHash("example@example.com", nil)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestEmailHasher_ProcessStringIntoHash_InvalidEmail(t *testing.T) {
	h := NewEmailHasher()

	_, err := h.ProcessStringIntoHash("invalid emails example .com", nil)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestEmailHasher_ProcessStringIntoHash_RemoveExtraDots(t *testing.T) {
	h := NewEmailHasher()

	v, err := h.ProcessStringIntoHash("example.......email@example..com", nil)
	expected := "example.email@example.com"

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestEmailHasher_ProcessStringIntoHash_RemoveExtraAts(t *testing.T) {
	h := NewEmailHasher()

	v, err := h.ProcessStringIntoHash("example@@example.com", nil)
	expected := "example@example.com"

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestEmailHasher_ProcessStringIntoHash_RemoveExtraSpaces(t *testing.T) {
	h := NewEmailHasher()

	v, err := h.ProcessStringIntoHash("example  @e  xample .com", nil)
	expected := "example@example.com"

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestEmailHasher_ProcessStringIntoHash_RemoveExtraNonASCII(t *testing.T) {
	h := NewEmailHasher()

	v, err := h.ProcessStringIntoHash("éxámple@exámple.com", nil)
	expected := "example@example.com"

	fmt.Println(v, err)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestEmailHasher_ProcessStringIntoHash_RemoveExtraNonAlphanumeric(t *testing.T) {
	h := NewEmailHasher()

	v, err := h.ProcessStringIntoHash("/><example@exa!mple.com", nil)
	expected := "example@example.com"

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestEmailHasher_ProcessStringIntoHash_ExtraData(t *testing.T) {
	h := NewEmailHasher()

	v, err := h.ProcessStringIntoHash("example+extra@example.com", nil)
	expected := "example@example.com"

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestEmailHasher_ProcessStringIntoHash_FindEmailWithoutNameOrNameAndSurname(t *testing.T) {
	h := NewEmailHasher()

	v, err := h.ProcessStringIntoHash("example example <example@example.com>", nil)
	expected := "example@example.com"

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestPhoneHasher_ProcessStringIntoHash_ValidPhone(t *testing.T) {
	h := NewPhoneHasher()

	_, err := h.ProcessStringIntoHash("123456789", ExtraDataHasher{
		"country_code": "PL",
	})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestPhoneHasher_ProcessStringIntoHash_InvalidPhone(t *testing.T) {
	h := NewPhoneHasher()

	_, err := h.ProcessStringIntoHash("invalid phone number", ExtraDataHasher{
		"country_code": "PL",
	})

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestPhoneHasher_ProcessStringIntoHash_RemoveExtraSpaces(t *testing.T) {
	h := NewPhoneHasher()

	v, err := h.ProcessStringIntoHash(" 123 45 6789 ", ExtraDataHasher{
		"country_code": "PL",
	})
	expected := "123456789"

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestPhoneHasher_ProcessStringIntoHash_RemoveExtraNonNumeric(t *testing.T) {
	h := NewPhoneHasher()

	v, err := h.ProcessStringIntoHash("123-45-6789", ExtraDataHasher{
		"country_code": "PL",
	})
	expected := "123456789"

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestPhoneHasher_ProcessStringIntoHash_ExtraData(t *testing.T) {
	h := NewPhoneHasher()

	v, err := h.ProcessStringIntoHash("+48 123456789", ExtraDataHasher{
		"country_code": "PL",
	})
	expected := "123456789"

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestPhoneHasher_ProcessStringIntoHash_InvalidCountry(t *testing.T) {
	h := NewPhoneHasher()

	_, err := h.ProcessStringIntoHash("+1234567890", ExtraDataHasher{
		"country_code": "INVALID",
	})

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestPhoneHasher_ProcessStringIntoHash_CheckByCountry(t *testing.T) {
	h := NewPhoneHasher()

	_, err := h.ProcessStringIntoHash("+49 123456789", ExtraDataHasher{
		"country_code": "PL",
	})

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestNameHasher_ProcessStringIntoHash_ValidName(t *testing.T) {
	h := NewNameHasher()

	_, err := h.ProcessStringIntoHash("John", nil)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestNameHasher_ProcessStringIntoHash_InvalidName(t *testing.T) {
	h := NewNameHasher()

	_, err := h.ProcessStringIntoHash("", nil)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestNameHasher_ProcessStringIntoHash_RemoveExtraSpaces(t *testing.T) {
	h := NewNameHasher()

	v, err := h.ProcessStringIntoHash("  J ohn  ", nil)
	expected := "john"

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestNameHasher_ProcessStringIntoHash_RemoveExtraNonASCII(t *testing.T) {
	h := NewNameHasher()

	v, err := h.ProcessStringIntoHash("Jóhn", nil)
	expected := "john"

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestNameHasher_ProcessStringIntoHash_RemoveExtraNonAlphanumeric(t *testing.T) {
	h := NewNameHasher()

	v, err := h.ProcessStringIntoHash("Joh!n", nil)
	expected := "john"

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if v != expected {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}
