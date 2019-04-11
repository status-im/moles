package operations

import "github.com/status-im/moles"

type Remove struct {
	Removed uint32
	Path    moles.DirectPath
}
