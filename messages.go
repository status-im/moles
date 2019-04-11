package moles

import "github.com/decanus/moles/group"

type ApplicationMessageContent struct {
	Content   []byte
	Signature []byte
	Zeros     []uint8
}

type ApplicationMessage struct {
	Group            [32]uint8
	Epoch            uint32
	Generation       uint32
	Sender           uint32
	EncryptedContent []byte
}

type Handshake struct {
	PriorEpoch   uint32
	Operation    group.Operation
	SignerIndex  uint32
	Signature    []byte
	Confirmation []byte
}
