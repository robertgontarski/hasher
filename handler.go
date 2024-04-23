package main

import (
	"encoding/json"
	"io"
)

type Handler interface {
	Handle(r io.Reader) error
}

type DefaultHandler struct {
	Store  Store
	Hasher Hasher
}

type EmailHandler struct {
	DefaultHandler
	Handler
}

func NewEmailHandler(store Store, hasher Hasher) *EmailHandler {
	return &EmailHandler{
		DefaultHandler: DefaultHandler{
			Store:  store,
			Hasher: hasher,
		},
	}
}

func (h *EmailHandler) Handle(r io.Reader) error {
	var msg EmailChangeMessage
	if err := json.NewDecoder(r).Decode(&msg); err != nil {
		return err
	}

	hash, err := h.Hasher.ProcessStringIntoHash(msg.Address, nil)
	if err != nil {
		return err
	}

	if err := h.Store.UpdateHashOnEmailByID(msg.ID, hash); err != nil {
		return err
	}

	return nil
}

type PhoneHandler struct {
	DefaultHandler
	Handler
}

func NewPhoneHandler(store Store, hasher Hasher) *PhoneHandler {
	return &PhoneHandler{
		DefaultHandler: DefaultHandler{
			Store:  store,
			Hasher: hasher,
		},
	}
}

func (h *PhoneHandler) Handle(r io.Reader) error {
	var msg PhoneChangeMessage
	if err := json.NewDecoder(r).Decode(&msg); err != nil {
		return err
	}

	hash, err := h.Hasher.ProcessStringIntoHash(msg.Number, ExtraDataHasher{
		"country_code": msg.CountryCode,
	})

	if err != nil {
		return err
	}

	if err := h.Store.UpdateHashOnPhoneByID(msg.ID, hash); err != nil {
		return err
	}

	return nil
}

type NameHandler struct {
	DefaultHandler
	Handler
}

func NewNameHandler(store Store, hasher Hasher) *NameHandler {
	return &NameHandler{
		DefaultHandler: DefaultHandler{
			Store:  store,
			Hasher: hasher,
		},
	}
}

func (h *NameHandler) Handle(r io.Reader) error {
	var msg NameChangeMessage
	if err := json.NewDecoder(r).Decode(&msg); err != nil {
		return err
	}

	hashName, err := h.Hasher.ProcessStringIntoHash(msg.Name, nil)
	if err != nil {
		return err
	}

	hashSurname, err := h.Hasher.ProcessStringIntoHash(msg.Surname, nil)
	if err != nil {
		return err
	}

	if err := h.Store.UpdateHashOnNameByID(msg.ID, hashName, hashSurname); err != nil {
		return err
	}

	return nil
}
