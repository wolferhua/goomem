package goomem

const (
	CmdGet = "GET"
	CmdSet = "SET"
)

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

func (cmd *Command) GetArg(i int) interface{} {
	l := len(cmd.Args)
	if i < l && i >= 0 {
		return cmd.Args[i]
	}
	return nil
}
