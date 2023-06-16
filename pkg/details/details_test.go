/*
Copyright © 2023 Dex Wood
*/
package details

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
)

func TestDates(t *testing.T) {
	retrieveDates, err := RetrieveCertDetails("www.example.edu:443", false)
	if err != nil {
		t.Fatalf(err.Error())
	}
	for _, details := range retrieveDates {
		fmt.Printf("Issuer: %s Expiration Date: %v\n", details.Issuer, details.NotAfter.Format("2006-01-02"))
		if details.Cert.PublicKeyAlgorithm == x509.RSA {
			//pk := details.Cert.PublicKey.(*rsa.PublicKey)
			pemEncoded := pem.EncodeToMemory(&pem.Block{
				Type:  "CERTIFICATE",
				Bytes: details.Cert.Raw,
			})
			fmt.Println(len(pemEncoded))
			fmt.Println(string(pemEncoded))
		}

	}

}
