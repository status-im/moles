package moles

type CredentialType uint8

const (
	BASIC CredentialType = 0
	X509  CredentialType = 1
)

type GroupOperationType uint8

const (
	INIT   GroupOperationType = 0
	ADD    GroupOperationType = 1
	UPDATE GroupOperationType = 2
	REMOVE GroupOperationType = 3
)

type CipherSuite uint8

const (
	P256_SHA256_AES128GCM   CipherSuite = 0x0000
	X25519_SHA256_AES128GCM CipherSuite = 0x0001
)

type SignatureScheme uint8

const (
	ECSDA_SECP256R1_SHA256 SignatureScheme = 0x0403
	ED25519                SignatureScheme = 0x0807
)
