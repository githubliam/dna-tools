package native

import (
	"github.com/dnaproject2/dna-tool/methods/smartcontract/native/governance"
)

func RegisterNative() {
	governance.RegisterGovernance()
}
