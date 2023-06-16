/*
Copyright © 2023 Dex Wood
*/
package details

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"
)

type CertDetails struct {
	NotAfter time.Time
	Issuer   string
	DNSNames []string
	Cert     *x509.Certificate
}

func RetrieveCertDetails(address string, insecure bool) ([]CertDetails, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: insecure,
	}
	d := tls.Dialer{Config: tlsConfig}
	conn, err := d.DialContext(ctx, "tcp", address)
	cancel()
	if err != nil {
		return []CertDetails{}, err
	}
	if client, ok := conn.(*tls.Conn); ok {
		certificates := client.ConnectionState().PeerCertificates
		details := make([]CertDetails, len(certificates))
		for i, cert := range certificates {
			details[len(certificates)-i-1] = CertDetails{
				NotAfter: cert.NotAfter,
				Issuer:   cert.Issuer.String(),
				DNSNames: cert.DNSNames,
				Cert:     cert,
			}
		}
		return details, nil
	}
	return nil, errors.New("could not create TLS connect")

}
func DisplayPemCertificate(details CertDetails) {
	if details.Cert.PublicKeyAlgorithm == x509.RSA {
		pemEncoded := pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: details.Cert.Raw,
		})
		fmt.Println(string(pemEncoded))
	}
}
