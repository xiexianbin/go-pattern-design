package interpreter

type Expression interface {
	Interpret() int
}

type Context struct {
	value int
}

func (c *Context) GetValue() int {
	return c.value
}

type AddExpression struct {
	left  Context
	right Context
}

func (e *AddExpression) Interpret() int {
	return e.left.GetValue() + e.right.GetValue()
}

type SubExpression struct {
	left  Context
	right Context
}

func (e *SubExpression) Interpret() int {
	return e.left.GetValue() - e.right.GetValue()
}

func DoExpression(exp string, left, right Context) int {
	switch exp {
	case "add":
		e := &AddExpression{
			left:  left,
			right: right,
		}
		return e.Interpret()
	case "sub":
		e := &SubExpression{
			left:  left,
			right: right,
		}
		return e.Interpret()
	default:
		return -1
	}
}
