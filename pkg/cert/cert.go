// SPDX-License-Identifier: BSD-3-Clause

// Package cert provides functionality for creating self-signed and signed certificates.
package cert

// Constants for certificate creation operations.
const (
	// CertCreateSelfsigned represents the operation of creating a self-signed certificate.
	CertCreateSelfsigned = "create_cert"
	// CertCreateSigned represents the operation of creating a signed certificate.
	CertCreateSigned = "create_signed_cert"
	// StoreCert represents the operation of storing a certificate.
	StoreCert = "store_cert"
	// StoreKey represents the operation of storing a key.
	StoreKey = "store_key"
	// GetCert represents the operation of retrieving a certificate.
	GetCert = "get_cert"
	// GetKey represents the operation of retrieving a key.
	GetKey = "get_key"
)
