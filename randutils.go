package randutils

import (
	"crypto/rand"
	"math/big"
)

const (
	Alphabet      = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	Numerals      = "0123456789"
	Alphanumeric  = Alphabet + Numerals
	Ascii         = Alphanumeric + "~!@#$%^&amp;*()-_+={}[]\\|&lt;,&gt;.?/\"';:`"
	HumanFriendly = "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz23456789"
)

// String returns a random alphanumeric string of the given length
func String(length int) (string, error) {
	return StringFromCharset(length, Alphanumeric)
}

// StringFromCharset returns a random string of the given langth consisting of
// characters from the charset
func StringFromCharset(length int, charset string) (string, error) {
	result := make([]byte, length) // Random string to return
	charsetlen := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		b, err := rand.Int(rand.Reader, charsetlen)
		if err != nil {
			return "", err
		}
		r := int(b.Int64())
		result[i] = charset[r]
	}

	return string(result), nil
}
