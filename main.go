package merf

import (
	"fmt"
	"github.com/onsi/ginkgo"
)

type TestingT interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Parallel()
	Skip(args ...interface{})
	Skipf(format string, args ...interface{})
	SkipNow(args ...interface{})
	Skipped() bool
}

type merf struct{}

func T() TestingT {
	return merf{}
}

func (m merf) Error(args ...interface{}) {
	m.Log(args)
	ginkgo.Fail("failed")
}

func (m merf) Errorf(format string, args ...interface{}) {
	ginkgo.Fail(fmt.Sprintf(format, args...))
}

func (m merf) Fail() {
	ginkgo.Fail("failed")
}

func (m merf) FailNow() {
	ginkgo.Fail("failed")
}

func (m merf) Failed() bool {
	return false
}

func (m merf) Fatal(args ...interface{}) {
	m.Log(args)
	m.FailNow()
}

func (m merf) Fatalf(format string, args ...interface{}) {
	m.Logf(format, args...)
	m.FailNow()
}

func (m merf) Log(args ...interface{}) {
	for _, log := range args {
		println(log)
	}
}

func (m merf) Logf(format string, args ...interface{}) {
	println(fmt.Sprintf(format, args...))
}

func (m merf) Parallel() {
	return
}

func (m merf) Skip(args ...interface{}) {
	m.Log(args...)
}

func (m merf) Skipf(format string, args ...interface{}) {
	m.Logf(format, args...)
}

func (m merf) SkipNow(args ...interface{}) {
	m.Log(args...)
	return
}

func (m merf) Skipped() bool {
	return false
}
