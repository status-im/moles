package group

import (
	"github.com/decanus/moles"
	"github.com/decanus/moles/group/operations"
)

type State struct {
	GroupId        []byte
	Epoch          uint32
	Roster         []moles.Credential
	Tree           []moles.HPKEPublicKey
	TranscriptHash []byte
}

type Operation struct {
	MsgType moles.GroupOperationType
	Add     operations.Add
	Update  operations.Update
	Remove  operations.Remove
}
