package moles

type SignaturePublicKey []byte

type BasicCredential struct {
	Identity  []byte
	Algorithm SignatureScheme
	PublicKey SignaturePublicKey
}

type Credential struct {
	Type CredentialType
	Basic BasicCredential
	X509 struct {
		CertData []byte
	}
}
