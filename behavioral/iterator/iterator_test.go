package iterator

import (
	"fmt"
	"testing"
)

func TestNewIterator(t *testing.T) {
	ta := &Teacher{
		Name: "A",
	}
	tb := &Teacher{
		Name: "B",
	}
	tc := &Teacher{
		Name: "C",
	}
	iterator := NewIterator()
	iterator.Add(ta)
	iterator.Add(tb)
	iterator.Add(tc)

	if len(iterator.list) != 3 {
		t.Fatal("i.list != 3")
	}

	for ; iterator.HasNext(); {
		fmt.Println(iterator.Next().GetName())
	}
}
