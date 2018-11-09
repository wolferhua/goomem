package goomem

func Get(key string) *Command {
	return NewCommand("get",key)
}
