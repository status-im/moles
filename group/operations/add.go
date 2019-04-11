package operations

import "github.com/status-im/moles"

type WelcomeInfo struct {
	Version        moles.ProtocolVersion
	GroupId        []byte
	Epoch          uint32
	Roster         []moles.Credential
	Tree           []moles.HPKEPublicKey
	TranscriptHash []byte
	InitSecret     []byte
}

type Welcome struct {
	UserInitKeyId        []byte
	CipherSuite          moles.CipherSuite
	EncryptedWelcomeInfo moles.HPKEPublicKey
}

type Add struct {
	Index           uint32
	InitKey         moles.UserInitKey
	WelcomeInfoHash []byte
}
