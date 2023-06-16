/*
Copyright © 2023 Dex Wood
*/
package gen

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io"
)

type CsrInputInfo struct {
	CommonName string
	Sans       []string
	pkix.Name
	PrivKey *rsa.PrivateKey
}

type CsrOutputInfo struct {
	CsrPem, PrivateKeyPem string
}

func NewCsr(source io.Reader, csrInfo CsrInputInfo, encryptKey bool, encryptKeyPass string) (CsrOutputInfo, error) {
	cr := x509.CertificateRequest{
		Subject:  csrInfo.Name,
		DNSNames: csrInfo.Sans,
	}
	request, err := x509.CreateCertificateRequest(source, &cr, csrInfo.PrivKey)
	if err != nil {
		return CsrOutputInfo{}, err
	}
	csrPem := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: request,
	})
	pkcs1PrivKey := x509.MarshalPKCS1PrivateKey(csrInfo.PrivKey)
	var outPrivPem []byte
	if encryptKey {
		encBlock, err := x509.EncryptPEMBlock(source, "RSA PRIVATE KEY", pkcs1PrivKey, []byte(encryptKeyPass), x509.PEMCipherAES256)
		if err != nil {
			fmt.Println("Can't encrypt private key.")
		}
		outPrivPem = pem.EncodeToMemory(encBlock)
	} else {
		privPem := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: pkcs1PrivKey,
		})
		outPrivPem = privPem
	}

	return CsrOutputInfo{string(csrPem), string(outPrivPem)}, err
}

func NewCsrSecure(csrInfo CsrInputInfo, encryptKey bool, encryptKeyPass string) (CsrOutputInfo, error) {
	return NewCsr(rand.Reader, csrInfo, encryptKey, encryptKeyPass)
}
