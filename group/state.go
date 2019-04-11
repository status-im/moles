package group

import "github.com/decanus/moles"

type State struct {
	GroupId        []byte
	Epoch          uint32
	Roster         []moles.Credential
	Tree           []moles.HPKEPublicKey
	TranscriptHash []byte
}
