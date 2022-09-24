package flyweight

import (
	"fmt"
	"github.com/google/uuid"
)

type Bullet struct {
	ID     string
	Living bool
}

type BulletPool struct {
	List []*Bullet
	size int
}

func (b *BulletPool) GetBullet() *Bullet {
	if len(b.List) == 0 {
		b.size = 5
		for i := 0; i < b.size; i++ {
			b.List = append(b.List, &Bullet{ID: uuid.New().String(), Living: true})
		}
	}

	for i := 0; i < len(b.List); i++ {
		if b.List[i].Living == true {
			return b.List[i]
		}
	}

	fmt.Println("create new Bullet...")
	nb := &Bullet{ID: uuid.New().String(), Living: true}
	b.List = append(b.List, nb)
	return nb
}
