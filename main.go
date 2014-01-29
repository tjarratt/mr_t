package merf

import (
	"github.com/onsi/ginkgo"
)

type TestingT interface {
	Errorf(format string, args ...interface{})
}

type merf struct {}

func T() TestingT {
	return merf{}
}

func (m merf) Errorf(format string, args ...interface{}) {
	ginkgo.Fail(format, args...)
}
