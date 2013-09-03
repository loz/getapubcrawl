package pubcrawl

import (
	"github.com/remogatto/prettytest"
	"testing"
	//"fmt"
	//"strconv"
)

type testSuite struct {
	prettytest.Suite
}

func TestRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testSuite),
	)
}
