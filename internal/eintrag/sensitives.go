package eintrag

import (
	"encoding/base64"
	"errors"
	"github.com/jamesruan/sodium"
)

var masterKey sodium.MasterKey

const SaltSize = 32
const SubKeySize = 32

type ISensitiveJar interface {
	GetIdentifier() string
	SetNonce(*sodium.SecretBoxNonce) error
	ToJSON() error
}

func SetMasterKey(m string) error {
	if "" == m {
		return errors.New("master key cannot be empty")
	}

	keyBytes := make([]byte, base64.StdEncoding.DecodedLen(len(m)))
	if _, err := base64.StdEncoding.Decode(keyBytes, []byte(m)); nil != err {
		return err
	}

	masterKey = sodium.MasterKey{
		Bytes: keyBytes,
	}

	return nil
}

func IdentifierToKeyId(identifier string) (uint64, error) {
	if "" == identifier {
		return 0, errors.New("identifier cannot be empty")
	}

	id := uint64(0)
	for _, c := range []byte(identifier) {
		id += uint64(c)
	}

	return id, nil
}

func DeriveKey(identifier string) (*sodium.SubKey, error) {
	ctx := sodium.MakeKeyContext(identifier)
	keyId, err := IdentifierToKeyId(identifier)
	if nil != err {
		return nil, err
	}

	subKey := masterKey.Derive(sodium.CryptoKDFBytesMin, keyId, ctx)

	return &subKey, err
}

// TODO: Chacha20-Poly1305 AEAD
func Encrypt(jar ISensitiveJar) ([]byte, error) {
	subKey, err := DeriveKey(jar.GetIdentifier())
	if nil != err {
		return nil, err
	}

	nonce := sodium.SecretBoxNonce{}
	nonce.Next()

	err = jar.SetNonce(&nonce)

	box := sodium.SecretBox

	payload, err := jar.ToJSON()
	if nil != err {
		return nil, err
	}

	subKey.SecretBoxOpen()

	data, err := subKey.Encrypt(payload)

	return subKey.Encrypt(data)
}

// TODO Data structures and functions for handling sensitive data
