package builder

type Person struct {
	Name   string
	Age    int
	Weight float32
	Sex    string
	Email  string
	Addr   string
}

type IPersonBuilder interface {
	SetName(name string) IPersonBuilder
	SetAge(age int) IPersonBuilder
	SetWeight(weight float32) IPersonBuilder
	SetSex(sex string) IPersonBuilder
	SetEmail(email string) IPersonBuilder
	SetAddr(addr string) IPersonBuilder
	Build() Person
}

type PersonBuilder struct {
	person Person
}

func (p *PersonBuilder) SetName(name string) IPersonBuilder {
	p.person.Name = name
	return p
}

func (p *PersonBuilder) SetAge(age int) IPersonBuilder {
	p.person.Age = age
	return p
}

func (p *PersonBuilder) SetWeight(weight float32) IPersonBuilder {
	p.person.Weight = weight
	return p
}

func (p *PersonBuilder) SetSex(sex string) IPersonBuilder {
	p.person.Sex = sex
	return p
}

func (p *PersonBuilder) SetEmail(email string) IPersonBuilder {
	p.person.Email = email
	return p
}

func (p *PersonBuilder) SetAddr(addr string) IPersonBuilder {
	p.person.Addr = addr
	return p
}

func (p *PersonBuilder) Build() Person {
	return p.person
}
