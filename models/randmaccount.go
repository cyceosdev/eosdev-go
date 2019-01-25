package models

import "crypto/rand"

const (
	StdLen  = 12
	UUIDLen = 20
)

var StdChars = []byte("abcdefghijklmnopqrstuvwxyz12345")

func NewRandomAccount() string {

	return NewLen(StdLen)
}

func NewLen(length int) string {

	return NewLenChars(length, StdChars)
}

func NewLenChars(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length for NewLenChars()")
	}

	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4))
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}
