package stacktrace_test

import (
	"github.com/houmanka/stacktrace"
)

const (
	EcodeInvalidVillain = stacktrace.ErrorCode(iota)
	EcodeNoSuchPseudo
	EcodeNotFastEnough
	EcodeTimeIsIllusion
	EcodeNotImplemented
)
