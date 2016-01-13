package context

type Context struct {
	name string
}

func NewContext(name string) *Context {
	c := Context{}
	c.name = name
	return &c
}

func (c Context) Name() string {
	return c.name
}
