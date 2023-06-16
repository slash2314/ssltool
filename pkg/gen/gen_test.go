/*
Copyright © 2023 Dex Wood
*/
package gen

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/rand"
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
	rndReader := rand.New(rand.NewSource(1))
	block, _ := pem.Decode([]byte(unencryptedTestCase.privKeyPem))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		t.Fatalf(err.Error())
	}
	csrInfo := CsrInputInfo{
		CommonName: "test.example.com",
		Sans:       []string{},
		Name:       subj,
		PrivKey:    key,
	}

	csrOutput, err := NewCsr(rndReader, csrInfo, false, "")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !(unencryptedTestCase.csrPem == csrOutput.CsrPem) {
		t.Fatalf("Expected \n%v\nActual \n%v\n", unencryptedTestCase.csrPem, csrOutput.CsrPem)
	}
	if !(unencryptedTestCase.privKeyPem == csrOutput.PrivateKeyPem) {
		t.Fatalf("Expected \n%v\nActual \n%v\n", unencryptedTestCase.privKeyPem, csrOutput.PrivateKeyPem)
	}
	//fmt.Printf("%s %s\n", csrOutput.CsrPem, csrOutput.PrivateKeyPem)
}
