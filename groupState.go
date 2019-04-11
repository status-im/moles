package moles

type HPKEPublicKey []byte

type GroupState struct {
	GroupId        []byte
	Epoch          uint32
	Roster         []Credential
	Tree           []HPKEPublicKey
	TranscriptHash []byte
}
