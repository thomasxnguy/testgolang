package jwt

import (
	"time"

	"gopkg.in/square/go-jose.v2"
)

// Encrypt return a JWE
func Encrypt(key *[]byte, data *[]byte) (*string, error) {
	enc, err := jose.NewEncrypter(jose.A256GCM, jose.Recipient{Algorithm: jose.A256KW, Key: *key}, nil)
	if err != nil {
		return nil, err
	}
	object, err := enc.Encrypt(*data)
	if err != nil {
		return nil, err
	}
	serialized, err := object.CompactSerialize()
	if err != nil {
		return nil, err
	}
	return &serialized, nil
}

//Decrypt a JWE
func Decrypt(key *[]byte, data *string) (*[]byte, error) {
	object, err := jose.ParseEncrypted(*data)
	if err != nil {
		return nil, err
	}
	output, err := object.Decrypt(*key)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// IsExpired return true if current time is greater than token ttl
func IsExpired(t *time.Time, ttl int) bool {
	duration := time.Since(*t)
	if duration.Seconds() > float64(ttl) {
		return true
	}
	return false
}