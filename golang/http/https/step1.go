package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ericchiang/letsencrypt"
)

func main() {
	// Create a client.
	cli, err := letsencrypt.NewClient("https://localhost:10001/")
	if err != nil {
		log.Fatal(err)
	}

	// Load the private key and register with the ACME server.
	key, err := loadKey("letsencrypt")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := cli.NewRegistration(key); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Registration successful!")
}

// loadKey reads and parses a private RSA key.
func loadKey(file string) (*rsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("pem decode: no key found")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}
