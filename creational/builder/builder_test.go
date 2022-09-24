package builder

import "testing"

func TestPersonBuilder_Build(t *testing.T) {
	pb := &PersonBuilder{}
	p := pb.SetName("xie").
		SetAge(18).
		SetEmail("me@xiexianbin.cn").Build()
	t.Logf("%#v", p)
}
