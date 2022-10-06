package context

func NewContext() Context {
	return Context{}
}

// used to pass context into controllers
type Context struct {
}
