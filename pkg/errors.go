package pkg

import (
	"errors"
)

var (
	// Some Predefined Errors to make life easier.
	ErrNotFound       = errors.New("error: unable to find resource")
	ErrDatabase       = errors.New("error: Something went wrong with the database")
	ErrShortBlock     = errors.New("error: length of ciphertext is smaller than aes block size")
	ErrEncrypt        = errors.New("error: unable to encrypt data")
	ErrDecrypt        = errors.New("error: unable to decrypt given token")
	ErrMalformedToken = errors.New("error: malformed token found. please generate a new one")
	ErrExpiredToken   = errors.New("error: your token is expired. please generate a new one")
	ErrUnAuthorized   = errors.New("error: malformed token please generate new")
)
