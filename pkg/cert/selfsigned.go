// SPDX-License-Identifier: BSD-3-Clause

package cert

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"time"
)

// GenerateSelfsigned creates a self-signed certificate for the given hostname.
// It returns the certificate and private key as byte slices, or an error if any step of the generation fails.
func GenerateSelfsigned(hostname string) ([]byte, []byte, error) {
	// Generate a new public/private key pair
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// Generate a random serial number for the certificate
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	sn, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, nil, err
	}

	// Create a template for the certificate
	template := &x509.Certificate{
		Subject: pkix.Name{
			CommonName:   hostname,
			Organization: []string{"u-bmc"},
		},
		DNSNames: []string{hostname},
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageServerAuth,
		},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageKeyAgreement | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		SerialNumber:          sn,
		NotBefore:             time.Now().Add(-30 * time.Second),
		NotAfter:              time.Now().Add(262980 * time.Hour),
		IsCA:                  true,
	}

	// Create the certificate using the template and the key pair
	derBytes, err := x509.CreateCertificate(rand.Reader, template, template, pub, priv)
	if err != nil {
		return nil, nil, err
	}

	// Encode the certificate to PEM format
	var certPem bytes.Buffer
	if err := pem.Encode(&certPem, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return nil, nil, err
	}

	// Marshal the private key to PKCS8 format
	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return nil, nil, err
	}

	// Encode the private key to PEM format
	var keyPem bytes.Buffer
	if err := pem.Encode(&keyPem, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		return nil, nil, err
	}

	// Return the certificate and private key
	return certPem.Bytes(), keyPem.Bytes(), nil
}
