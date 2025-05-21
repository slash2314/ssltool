/*
Copyright © 2023 Dex Wood
*/
package gen

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	mathrand "math/rand"
	"strings"
	"testing"
)

type GenTestCase struct {
	csrPem, privKeyPem string
}

var unencryptedTestCase = GenTestCase{
	csrPem: `-----BEGIN CERTIFICATE REQUEST-----
MIICpjCCAY4CAQAwYTELMAkGA1UEBhMCVVMxETAPBgNVBAgTCEtlbnR1Y2t5MRYw
FAYDVQQHEw1Cb3dsaW5nIEdyZWVuMRUwEwYDVQQKEwxFeGFtcGxlIEluYy4xEDAO
BgNVBAsTB0JpbGxpbmcwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC3
8pECcoZbjZY8tTKQ81v3JrPvimXvz23qPB/Fa7GAi2jcqGmlb54ag5Ug+az+A7WN
EclH5vFuCsP9K/aW9ES5TzJ2O/hGiZ35VG4SXGm8Qk+jt2HbC/8NZtTgJOYcO4Ln
T9GX+kSRuuVImQIxhAs9m5+33DtfdslkjtO2rxpLG3NCLSE0Q9z5msRosDbiCv1p
NyOshhqmkmirN3QssOzIjPYTE9XdR+PxrUXJO3h+F7bTGM7Q4/t94BCkntBeF/Ey
SOJvppydCyQRRiX/6zDcksMBZWz8E7BRRZLs+t4KaqzGlmyojD6ALkLgPphgbQ3u
Rp38/48CbqsfdnTek8rLAgMBAAGgADANBgkqhkiG9w0BAQsFAAOCAQEAb0xYBevB
OxdeWPjtZzIkIp8y0KhRhfql1KZv9+fgvFyW5zkyKG9f66UFUw79E9sFx6OKMQxf
7biNv5522VOVcp9aM4O8QkibxYhmB1ewdQIfLeEWBTjrRbMwp3vEMaXOGynJHglG
I6iwWrN26qlM4Xn1odehq82+Kkbwxw1+2kUxJvffYCi+fZcyMCQ3EVO2xd1tG9ip
CcFRryx8vhlH2z9is/wbWG7OZPC+w/6r5MI/JqRv8WhWb6g5WQT+k9fgG+DvgBBS
1hIn1ox3oRy0F/BF8fpzlTkgcTwGgdHxq3VWM4VpgNczAAnRwtUQwz2VlA02Fq8F
nCvsvEJc42Dszg==
-----END CERTIFICATE REQUEST-----
`,
	privKeyPem: `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAt/KRAnKGW42WPLUykPNb9yaz74pl789t6jwfxWuxgIto3Khp
pW+eGoOVIPms/gO1jRHJR+bxbgrD/Sv2lvREuU8ydjv4Romd+VRuElxpvEJPo7dh
2wv/DWbU4CTmHDuC50/Rl/pEkbrlSJkCMYQLPZuft9w7X3bJZI7Ttq8aSxtzQi0h
NEPc+ZrEaLA24gr9aTcjrIYappJoqzd0LLDsyIz2ExPV3Ufj8a1FyTt4fhe20xjO
0OP7feAQpJ7QXhfxMkjib6acnQskEUYl/+sw3JLDAWVs/BOwUUWS7PreCmqsxpZs
qIw+gC5C4D6YYG0N7kad/P+PAm6rH3Z03pPKywIDAQABAoIBAErKxuk+1aBuHlMK
vcewG4mPZAQWusHaDm/6CFUGJ8MDbBykIWiRzLAtJjqlKhCSQ4RbYgLpROsgV3Ye
JQJVyYerSvJGCfnsRJ43FRZEGv3f3G/NxW0CIK87S3zjo+iAzgXtL2Ri7vRuEIIH
LJwje0qOd0+TrVRwgQxxAmY6Ji+5BhgImdtJiB91KsZPxT+S+u0Bd3irx1Xv400B
6f0pbb55bl4+M+Eplo1b+qeUqlmccFAyWjEM/YWBbjz+c13TNeSck4QiusrQ2UWf
yvMxZeClQRv6uX9ZjfaWTVkSAsVGL81lMdR9YkxTCzNk0VDrzEn7JsL1uPWBAG6x
bvKzJXECgYEA10dgjNOlcOFybrHdymTwhYKwIrsCbtpqkT3IPxdM48GLn8BQPTrH
Ti/kVpHW37SvjIbXUqFtZmT6tN4Ir+iFg5L8w1y56oL8QsQtSMDAVWJn6g3MGbEP
BeAxjESI/+cEtQNpCPXLk47r0xY1A6yqh09ZLZRUSPvrk6h3ompyMHMCgYEA2r4C
KMxgtkVRU3YkTyeqVtotsQ8mWQkBN/+4LFcjPIM+C78i1vZH+5ezZS0oiLOrCAEL
joppZ9Vgt0d1eAZzbcmLeCJmNPHuyqSi4jujZZGstXN9c1xbx6LjbxpGlGkn4GH9
93o7aO9YLCawH1qppDjswmxpQSIdFZDIOAcvnkkCgYEAifGYsmRj3GL5umiKn8fy
Pvqw9nAybOnT42FdUOATStWYLDNEtxdU4orZm4cz0oCrBDba/n2l5jjKVN75Xg52
jwq+oEhocRahof6mrbmlBJFb2KQipkvIPpGb7i90QdW8NGkNAsrNOa4Y+ld+fO5F
DxMPJ5+mKmYSW7lVf2MJ7HUCgYEAsjME81O8ngDOhh3C6rE3tdW8T3g2F2aclZA/
6+95Bz0r+MnXiiPM9IvbW9t0IBmuhbDER3U+9ZYBWo5ehk9LDe+ZLV9owE0v6epB
v+gx7vbEKnZRhv+AzZxHiCVxxkn8cHGkQk5Tw+Log99or8JeXSj6yFElViiCZSUz
12ETS/ECgYEAqwO4nRg2hNpXsCO6pH/01wnzSErkKHkrxbjUc5RxeXW4jYA879Xk
zope3LynrHYz3MNAYgDRec9OZ3dvGujKhVpnhNa69hQsUA3vHe1OGG+f6rIlGera
2eoDLWU1MJJC9cbpHr3M26igEjXGkQxSfJcfRQkeoCI3i7OywbTA9+Y=
-----END RSA PRIVATE KEY-----
`,
}

func TestGen(t *testing.T) {
	subj := pkix.Name{
		Country:            []string{"US"},
		Organization:       []string{"Example Inc."},
		OrganizationalUnit: []string{"Billing"},
		Locality:           []string{"Bowling Green"},
		Province:           []string{"Kentucky"},
	}
	rndReader := mathrand.New(mathrand.NewSource(0))
	block, _ := pem.Decode([]byte(unencryptedTestCase.privKeyPem))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		t.Fatal(err.Error())
	}
	csrInfo := CsrInputInfo{
		CommonName: "test.example.com",
		Sans:       []string{},
		Name:       subj,
		PrivKey:    key,
	}

	csrOutput, err := NewCsr(rndReader, csrInfo)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Validate CSR structure
	block, _ = pem.Decode([]byte(csrOutput.CsrPem))
	if block == nil {
		t.Fatal("Failed to decode generated CSR PEM")
	}
	csr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		t.Fatalf("Failed to parse generated CSR: %v", err)
	}

	// Verify subject components
	if csr.Subject.CommonName != "test.example.com" {
		t.Errorf("CommonName mismatch. Expected 'test.example.com', got '%s'", csr.Subject.CommonName)
	}
	if !stringSliceEqual(csr.Subject.Country, []string{"US"}) ||
		!stringSliceEqual(csr.Subject.Organization, []string{"Example Inc."}) ||
		!stringSliceEqual(csr.Subject.OrganizationalUnit, []string{"Billing"}) ||
		!stringSliceEqual(csr.Subject.Locality, []string{"Bowling Green"}) ||
		!stringSliceEqual(csr.Subject.Province, []string{"Kentucky"}) {
		t.Error("Subject components do not match expected values")
	}

	// Verify private key matches
	if strings.TrimSpace(csrOutput.PrivateKeyPem) != strings.TrimSpace(unencryptedTestCase.privKeyPem) {
		t.Error("Private key PEM doesn't match expected value")
	}
}

func TestECDSAGeneration(t *testing.T) {
	subj := pkix.Name{
		Country:            []string{"US"},
		Organization:       []string{"Test Org"},
		OrganizationalUnit: []string{"Test OU"},
		Locality:           []string{"Test City"},
		Province:           []string{"Test State"},
	}

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate ECDSA key: %v", err)
	}

	csrInfo := CsrInputInfo{
		CommonName: "test-ecdsa.example.com",
		Sans:       []string{"test-ecdsa.example.com", "www.test-ecdsa.example.com"},
		Name:       subj,
		PrivKey:    key,
	}

	csrOutput, err := NewCsrSecure(csrInfo)
	if err != nil {
		t.Fatalf("Failed to generate CSR: %v", err)
	}

	// Validate CSR
	block, _ := pem.Decode([]byte(csrOutput.CsrPem))
	if block == nil {
		t.Fatal("Failed to decode CSR PEM")
	}
	csr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		t.Fatalf("Failed to parse CSR: %v", err)
	}

	if csr.Subject.CommonName != "test-ecdsa.example.com" {
		t.Errorf("CommonName mismatch: got %s", csr.Subject.CommonName)
	}
	if !stringSliceEqual(csr.DNSNames, []string{"test-ecdsa.example.com", "www.test-ecdsa.example.com"}) {
		t.Errorf("SANs mismatch: got %v", csr.DNSNames)
	}

	// Validate private key
	privBlock, _ := pem.Decode([]byte(csrOutput.PrivateKeyPem))
	if privBlock == nil {
		t.Fatal("Failed to decode private key PEM")
	}
	if privBlock.Type != "EC PRIVATE KEY" {
		t.Errorf("Expected EC PRIVATE KEY, got %s", privBlock.Type)
	}
	_, err = x509.ParseECPrivateKey(privBlock.Bytes)
	if err != nil {
		t.Errorf("Failed to parse EC private key: %v", err)
	}
}

func TestEd25519Generation(t *testing.T) {
	subj := pkix.Name{
		Country:            []string{"US"},
		Organization:       []string{"Test Org"},
		OrganizationalUnit: []string{"Test OU"},
		Locality:           []string{"Test City"},
		Province:           []string{"Test State"},
	}

	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate Ed25519 key: %v", err)
	}

	csrInfo := CsrInputInfo{
		CommonName: "test-ed25519.example.com",
		Sans:       []string{"test-ed25519.example.com"},
		Name:       subj,
		PrivKey:    privKey,
	}

	csrOutput, err := NewCsrSecure(csrInfo)
	if err != nil {
		t.Fatalf("Failed to generate CSR: %v", err)
	}

	// Validate CSR
	block, _ := pem.Decode([]byte(csrOutput.CsrPem))
	if block == nil {
		t.Fatal("Failed to decode CSR PEM")
	}
	csr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		t.Fatalf("Failed to parse CSR: %v", err)
	}

	if csr.Subject.CommonName != "test-ed25519.example.com" {
		t.Errorf("CommonName mismatch: got %s", csr.Subject.CommonName)
	}

	// Validate private key
	privBlock, _ := pem.Decode([]byte(csrOutput.PrivateKeyPem))
	if privBlock == nil {
		t.Fatal("Failed to decode private key PEM")
	}
	if privBlock.Type != "PRIVATE KEY" {
		t.Errorf("Expected PRIVATE KEY, got %s", privBlock.Type)
	}
	_, err = x509.ParsePKCS8PrivateKey(privBlock.Bytes)
	if err != nil {
		t.Errorf("Failed to parse Ed25519 private key: %v", err)
	}

	if !pubKey.Equal(csr.PublicKey) {
		t.Error("CSR public key does not match generated Ed25519 key")
	}
}

func TestMissingCommonNameAndSans(t *testing.T) {
	subj := pkix.Name{
		Country: []string{"US"},
	}

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key: %v", err)
	}

	csrInfo := CsrInputInfo{
		CommonName: "",
		Sans:       []string{},
		Name:       subj,
		PrivKey:    key,
	}

	_, err = NewCsrSecure(csrInfo)
	if err == nil {
		t.Fatal("Expected error when CommonName and SANs are empty")
	}
	expectedErr := "at least one of CommonName or SANs must be provided"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%v'", expectedErr, err)
	}
}

func TestOnlySans(t *testing.T) {
	subj := pkix.Name{
		Country: []string{"US"},
	}

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key: %v", err)
	}

	csrInfo := CsrInputInfo{
		CommonName: "",
		Sans:       []string{"sans.example.com"},
		Name:       subj,
		PrivKey:    key,
	}

	csrOutput, err := NewCsrSecure(csrInfo)
	if err != nil {
		t.Fatalf("Failed to generate CSR: %v", err)
	}

	block, _ := pem.Decode([]byte(csrOutput.CsrPem))
	if block == nil {
		t.Fatal("Failed to decode CSR PEM")
	}
	csr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		t.Fatalf("Failed to parse CSR: %v", err)
	}

	if len(csr.DNSNames) != 1 || csr.DNSNames[0] != "sans.example.com" {
		t.Errorf("Expected SANs [sans.example.com], got %v", csr.DNSNames)
	}
}

func stringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
