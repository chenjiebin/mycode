// 测试let's encrypt免费的https服务
// https://letsencrypt.org/
// https://github.com/ericchiang/letsencrypt
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"log"

	"github.com/ericchiang/letsencrypt"
)

var supportedChallengs = []string{
	letsencrypt.ChallengeHTTP,
	letsencrypt.ChallengeTLSSNI,
}

func main() {
	cli, err := letsencrypt.NewClient("http://localhost:4000/directory")
	if err != nil {
		log.Fatal("failed to create client:", err)
	}

	// Create a private key for your account and register
	accountKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := cli.NewRegistration(accountKey); err != nil {
		log.Fatal("new registration failed:", err)
	}

	// ask for a set of challenges for a given domain
	auth, _, err := cli.NewAuthorization(accountKey, "dns", "example.org")
	if err != nil {
		log.Fatal(err)
	}
	chals := auth.Combinations(supportedChallenges...)
	if len(chals) == 0 {
		log.Fatal("no supported challenge combinations")
	}

	/*
	   review challenge combinations and complete them
	*/

	// create a certificate request for your domain
	csr, certKey, err := newCSR()
	if err != nil {
		log.Fatal(err)
	}

	// Request a certificate for your domain
	cert, err := cli.NewCertificate(accountKey, csr)
	if err != nil {
		log.Fatal(err)
	}
	// We've got a certificate. Let's Encrypt!
}

func newCSR() (*x509.CertificateRequest, *rsa.PrivateKey, error) {
	certKey, err := rsa.GenerateKey(rand.Random, 4096)
	if err != nil {
		return nil, nil, err
	}
	template := &x509.CertificateRequest{
		SignatureAlgorithm: x509.SHA256WithRSA,
		PublicKeyAlgorithm: x509.RSA,
		PublicKey:          &certKey.PublicKey,
		Subject:            pkix.Name{CommonName: "example.org"},
		DNSNames:           []string{"example.org"},
	}
	csrDER, err := x509.CreateCertificateRequest(rand.Reader, template, certKey)
	if err != nil {
		return nil, nil, err
	}
	csr, err := x509.ParseCertificateRequest(csrDER)
	if err != nil {
		return nil, nil, err
	}
	return csr, certKey, nil
}
