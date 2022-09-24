package flyweight

import (
	"testing"
)

func TestBulletPool(t *testing.T) {
	bp := BulletPool{}
	for i := 0; i < 10; i++ {
		b := bp.GetBullet()
		b.Living = false
		t.Log(b.ID)
	}
}
