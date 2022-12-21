// Author: PIROGOM
// https://modu-print.tistory.com
// mop.pirogom@gmail.com
// MIT License
package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

type RSAHelper struct {
	PriKey *rsa.PrivateKey
	PubKey *rsa.PublicKey
}

/**
* RSA Key 생성
**/
func (r *RSAHelper) GenerateKey(bits int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		r.PriKey = nil
		r.PubKey = nil
		return err
	}

	r.PriKey = privateKey
	r.PubKey = &privateKey.PublicKey

	return nil
}

func (r *RSAHelper) GetPrivate() *rsa.PrivateKey {
	return r.PriKey
}

func (r *RSAHelper) GetPublic() *rsa.PublicKey {
	return r.PubKey
}

/**
*	PrivateToBytePEM
*	PrivateToStringPEM
**/
func (r *RSAHelper) PrivateToBytePEM() ([]byte, error) {
	if r.PriKey == nil {
		return nil, errors.New("private rsa key is nil")
	}

	privateKeyBytes, _ := x509.MarshalPKCS8PrivateKey(r.PriKey)
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKeyBytes,
		},
	)
	return privateKeyPEM, nil
}

func (r *RSAHelper) PrivateToStringPEM() (string, error) {
	pemBuf, err := r.PrivateToBytePEM()

	if err != nil {
		return "", err
	}
	return string(pemBuf), nil
}

/**
*	PrivateFromBytePEM
*	PrivateFromFilePEM
**/
func (r *RSAHelper) PrivateFromBytePEM(privateKeyPEM []byte) error {
	block, _ := pem.Decode(privateKeyPEM)

	if block == nil {
		r.PriKey = nil
		return errors.New("invalid PEM Block(private key)")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err != nil {
		r.PriKey = nil
		return err
	}

	r.PriKey = privateKey.(*rsa.PrivateKey)
	return nil
}

func (r *RSAHelper) PrivateFromStringPEM(privateKeyPEM string) error {
	return r.PrivateFromBytePEM([]byte(privateKeyPEM))
}

func (r *RSAHelper) PrivateFromFilePEM(filename string) error {
	fileBuf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = r.PrivateFromBytePEM(fileBuf)
	return err
}

/**
*	PublicToBytePEM
* 	PublicToStringPEM
**/
func (r *RSAHelper) PublicToBytePEM() ([]byte, error) {
	if r.PubKey == nil {
		return nil, errors.New("public rsa key is nil")
	}
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(r.PubKey)
	if err != nil {
		return nil, err
	}
	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publicKeyBytes,
		},
	)
	return publicKeyPEM, nil
}

func (r *RSAHelper) PublicToStringPEM() (string, error) {
	buf, err := r.PublicToBytePEM()
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

/**
*	PublicFromBytePEM
**/
func (r *RSAHelper) PublicFromBytePEM(publicKeyPEM []byte) error {
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil {
		r.PubKey = nil
		return errors.New("invalid PEM Block (public key)")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		r.PubKey = nil
		return err
	}

	switch pub := publicKey.(type) {
	case *rsa.PublicKey:
		r.PubKey = pub
		return nil
	default:
		return errors.New("invalid key type")
	}
}

func (r *RSAHelper) PublicFromStringPEM(publicKeyPEM string) error {
	return r.PublicFromBytePEM([]byte(publicKeyPEM))
}

func (r *RSAHelper) PublicFromFilePEM(filename string) error {
	fileBuf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = r.PublicFromBytePEM(fileBuf)
	return err
}

/**
*	EncryptByte
**/
func (r *RSAHelper) EncryptByte(src []byte) ([]byte, error) {

	if r.PubKey == nil {
		return nil, errors.New("public key is nil")
	}

	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, r.PubKey, src)

	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

func (r *RSAHelper) EncryptString(src string) (string, error) {
	enc, err := r.EncryptByte([]byte(src))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(enc), nil
}

/**
* DecryptByte
**/
func (r *RSAHelper) DecryptByte(src []byte) ([]byte, error) {

	if r.PriKey == nil {
		return nil, errors.New("private key is nil")
	}

	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, r.PriKey, src)

	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func (r *RSAHelper) DecryptString(src string) (string, error) {
	bsrc, err := hex.DecodeString(src)
	if err != nil {
		return "", err
	}
	dec, err := r.DecryptByte(bsrc)

	if err != nil {
		return "", err
	}

	return string(dec), nil
}
