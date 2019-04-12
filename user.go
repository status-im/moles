package moles

type ProtocolVersion uint8

type UserInitKey struct {
	UserInitKeyID     []byte
	SupportedVersions []ProtocolVersion
	CipherSuites      []CipherSuite
	InitKeys          []HPKEPublicKey
	Credential        Credential
	Signature         []byte
}
