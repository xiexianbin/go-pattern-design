package composite

import "testing"

func TestComponent(t *testing.T) {
	root := &Branch{Name: "root"}
	leaf1 := &Leaf{Name: "l1"}
	leaf2 := &Leaf{Name: "l1"}
	branch1 := &Branch{Name: "b1", Chird: []Component{}}
	bl1 := &Leaf{Name: "bl1"}
	bl2 := &Leaf{Name: "bl1"}

	root.Add(leaf1)
	root.Add(leaf2)

	root.Add(branch1)
	branch1.Add(bl1)
	branch1.Add(bl2)

	root.Display(0)
}
