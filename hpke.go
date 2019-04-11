package moles

type HPKEPublicKey []byte

type HPKECipherText struct {
	EphemeralKey HPKEPublicKey
	CipherText   []byte
}

type RatchetNode struct {
	PublicKey   HPKEPublicKey
	NodeSecrets []HPKECipherText
}

type DirectPath struct {
	nodes []RatchetNode
}
