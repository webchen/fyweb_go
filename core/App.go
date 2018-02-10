package core

// IApp ，这个其实就算得上是baseController
type IApp interface {
	Init(ctx *Context)
	Context() *Context
}

// App ，下面这几个函数，就是App实现IApp接口，相当于继承IApp接口了
type App struct {
	ctx *Context
}

// Init the Context
func (a *App) Init(ctx *Context) {
	a.ctx = ctx
}

// Context , get the Context
func (a *App) Context() (c *Context) {
	c = a.ctx
	return
}
