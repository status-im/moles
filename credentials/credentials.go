package credentials

import "github.com/decanus/moles"

type SignaturePublicKey []byte

type BasicCredential struct {
	Identity  []byte
	Algorithm moles.SignatureScheme
	PublicKey SignaturePublicKey
}
