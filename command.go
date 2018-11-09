package goomem

type Command struct {
	Cmd  string        //命令
	Args []interface{} //参数
}

func NewCommand(cmd string, args ...interface{}) *Command {
	return &Command{
		Cmd:  cmd,
		Args: args,
	}
}
