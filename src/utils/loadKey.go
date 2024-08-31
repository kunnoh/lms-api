package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func loadPrivateKey(filename string) (*ecdsa.PrivateKey, error) {
	// load private key
	privateKeyPEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Decode the PEM block
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing ECDSA private key")
	}

	// Parse the private key
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func loadPublicKey(filename string) (*ecdsa.PublicKey, error) {
	// Read the public key file
	publicKeyPEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Decode the PEM block
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing ECDSA public key")
	}

	// Parse the public key
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Type assert to *ecdsa.PublicKey
	publicKey, ok := publicKeyInterface.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("not an ECDSA public key")
	}

	return publicKey, nil
}
