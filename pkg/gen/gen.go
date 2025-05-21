/*
Copyright © 2023 Dex Wood
*/
package gen

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
)

type CsrInputInfo struct {
	CommonName string
	Sans       []string
	pkix.Name
	PrivKey crypto.PrivateKey
}

type CsrOutputInfo struct {
	CsrPem, PrivateKeyPem string
}

func NewCsr(source io.Reader, csrInfo CsrInputInfo) (CsrOutputInfo, error) {
	if csrInfo.CommonName == "" && len(csrInfo.Sans) == 0 {
		return CsrOutputInfo{}, errors.New("at least one of CommonName or SANs must be provided")
	}
	// Adding the CommonName to the Sans
	if csrInfo.CommonName != "" {
		csrInfo.Name.CommonName = csrInfo.CommonName
	}

	cr := x509.CertificateRequest{
		Subject:  csrInfo.Name,
		DNSNames: csrInfo.Sans,
	}
	request, err := x509.CreateCertificateRequest(source, &cr, csrInfo.PrivKey)
	if err != nil {
		return CsrOutputInfo{}, fmt.Errorf("failed to create certificate request: %w", err)
	}
	csrPem := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: request,
	})

	var privKeyBytes []byte
	var pemType string
	switch key := csrInfo.PrivKey.(type) {
	case *rsa.PrivateKey:
		privKeyBytes = x509.MarshalPKCS1PrivateKey(key)
		pemType = "RSA PRIVATE KEY"
	case *ecdsa.PrivateKey:
		privKeyBytes, err = x509.MarshalECPrivateKey(key)
		if err != nil {
			return CsrOutputInfo{}, fmt.Errorf("failed to marshal EC private key: %w", err)
		}
		pemType = "EC PRIVATE KEY"
	case ed25519.PrivateKey:
		privKeyBytes, err = x509.MarshalPKCS8PrivateKey(key)
		if err != nil {
			return CsrOutputInfo{}, fmt.Errorf("failed to marshal Ed25519 private key: %w", err)
		}
		pemType = "PRIVATE KEY"
	default:
		return CsrOutputInfo{}, errors.New("unsupported private key type")
	}

	var outPrivPem []byte
	privPem := pem.EncodeToMemory(&pem.Block{
		Type:  pemType,
		Bytes: privKeyBytes,
	})
	outPrivPem = privPem

	return CsrOutputInfo{string(csrPem), string(outPrivPem)}, nil
}

func NewCsrSecure(csrInfo CsrInputInfo) (CsrOutputInfo, error) {
	return NewCsr(rand.Reader, csrInfo)
}
