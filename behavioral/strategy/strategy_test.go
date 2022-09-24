package strategy

import "testing"

func TestContext_DoCharge(t *testing.T) {
	var ctx *Context
	ctx = &Context{
		Strategy: &St1{},
	}
	ctx.DoCharge()

	ctx = &Context{
		Strategy: &St2{},
	}
	ctx.DoCharge()
}
