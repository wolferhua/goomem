package goomem

import "fmt"

const (
	_ = iota
	ERR_CODE_NOT_FOUND
	ERR_CODE_EXPIRED
)

const (
	ERR_MSG_NOT_FOUND = "key not found"
	ERR_MSG_EXPIRED   = "key not found"
)

type gm_error struct {
	Msg  string
	Code int
}

func (this *gm_error) Error() string {
	return fmt.Sprintf("[%d]%s", this.Code, this.Msg)
}
func (this *gm_error) ErrorCode() int {
	return this.Code
}

func newError(msg string, code int) *gm_error {
	return &gm_error{
		Msg:  msg,
		Code: code,
	}
}
