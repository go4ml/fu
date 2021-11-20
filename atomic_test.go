package fu

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Test_Atomic1(t *testing.T) {
	f := AtomicFlag{1}
	assert.Assert(t, f.State() == true)
	f.Clear()
	assert.Assert(t, f.State() == false)
	f.Set()
	assert.Assert(t, f.State() == true)
	f.Clear()
	assert.Assert(t, f.State() == false)

	f = AtomicFlag{0}
	assert.Assert(t, f.State() == false)
	f.Clear()
	assert.Assert(t, f.State() == false)
	f.Set()
	assert.Assert(t, f.State() == true)
}
