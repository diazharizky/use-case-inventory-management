package app

type Context struct {
	UserRepository IUserRepository
}

func NewContext() *Context {
	return &Context{}
}
