package details

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"testing"
	"time"
)

// startTestTLSServer starts a localhost TLS server with a self-signed cert.
// Returns its address, the parsed certificate, and a cleanup function.
func startTestTLSServer(t *testing.T) (addr string, cert *x509.Certificate, closeFunc func() error) {
	t.Helper()
	// generate RSA key
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}
	// build self-signed cert template
	now := time.Now()
	template := &x509.Certificate{
		SerialNumber: big.NewInt(42),
		Subject:      pkix.Name{CommonName: "localhost"},
		NotBefore:    now.Add(-time.Hour),
		NotAfter:     now.Add(time.Hour),
		DNSNames:     []string{"localhost"},
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, template, template, &priv.PublicKey, priv)
	if err != nil {
		t.Fatalf("failed to create certificate: %v", err)
	}
	cert, err = x509.ParseCertificate(derBytes)
	if err != nil {
		t.Fatalf("failed to parse certificate: %v", err)
	}
	// PEM encode cert and key
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		t.Fatalf("failed to build tls.Certificate: %v", err)
	}
	// listen on random port
	ln, err := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{Certificates: []tls.Certificate{tlsCert}})
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	// accept & handshake, then close
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}
			tlsConn := conn.(*tls.Conn)
			_ = tlsConn.Handshake()
			conn.Close()
		}
	}()
	return ln.Addr().String(), cert, ln.Close
}

func TestRetrieveCertDetails_LocalTLS(t *testing.T) {
	addr, cert, cleanup := startTestTLSServer(t)
	defer cleanup()

	// connecting with strict verification should fail
	if _, err := RetrieveCertDetails(addr, false); err == nil {
		t.Error("expected error with invalid/self-signed cert and insecure=false")
	}

	// insecure=true should succeed
	detailsList, err := RetrieveCertDetails(addr, true)
	if err != nil {
		t.Fatalf("expected success with insecure=true, got %v", err)
	}
	if len(detailsList) != 1 {
		t.Fatalf("expected 1 cert, got %d", len(detailsList))
	}
	d := detailsList[0]

	// check that the returned Cert matches our self-signed cert
	if !d.Cert.NotBefore.Equal(cert.NotBefore) || !d.Cert.NotAfter.Equal(cert.NotAfter) {
		t.Errorf("certificate validity mismatch; got [%v - %v], want [%v - %v]",
			d.Cert.NotBefore, d.Cert.NotAfter, cert.NotBefore, cert.NotAfter)
	}
	if d.Cert.SerialNumber.Cmp(cert.SerialNumber) != 0 {
		t.Errorf("serial number mismatch; got %v, want %v", d.Cert.SerialNumber, cert.SerialNumber)
	}
	// issuer should equal the template Subject (self-signed)
	if d.Issuer != cert.Issuer.String() {
		t.Errorf("issuer mismatch; got %q, want %q", d.Issuer, cert.Issuer.String())
	}
	// DNSNames should include "localhost"
	if len(d.DNSNames) != 1 || d.DNSNames[0] != "localhost" {
		t.Errorf("DNSNames mismatch; got %v, want [\"localhost\"]", d.DNSNames)
	}
}

func TestRetrieveCertDetails_InvalidHost(t *testing.T) {
	// connect to a non-listening address
	_, err := RetrieveCertDetails("127.0.0.1:0", true)
	if err == nil {
		t.Fatal("expected error when connecting to invalid host")
	}
}
