package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func loadKey(filename string) (*ecdsa.PrivateKey, error) {
	privateKeyPEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing ECDSA private key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
