package group

import (
	"github.com/status-im/moles"
	"github.com/status-im/moles/group/operations"
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
