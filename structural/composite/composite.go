package composite

import "fmt"

type Component interface {
	// Add(c Component)
	// Remove(c Component)
	Display(depth int)
}

type Branch struct {
	Name  string
	Chird []Component
}

func (b *Branch) Display(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf("  ")
	}
	fmt.Printf("%s\n", b.Name)
	depth++
	for _, c := range b.Chird {
		c.Display(depth)
	}
}

func (b *Branch) Add(c Component) {
	b.Chird = append(b.Chird, c)
}

func (b *Branch) Remove(c Component) {
	//for _, i := range b.Chird {
	//}
	b.Chird = append(b.Chird[:1], b.Chird[2:]...)
}

type Leaf struct {
	Name string
}

func (l *Leaf) Display(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf("  ")
	}
	fmt.Printf("%s\n", l.Name)
}
