package generrors

import "fmt"

type CodegenError struct {
	err error
	msg string
}

func (vce CodegenError) Error() string {

	if vce.err != nil {
		return fmt.Sprintf("Error: %s \n Hint: %s\n", vce.err, vce.msg)
	}
	if vce.msg == "" {
		return fmt.Sprintf("Error: %s\n", vce.err)
	}
	return fmt.Sprintf("Hint: %s\n", vce.msg)
}
func Err(msg string, err error) CodegenError {
	return CodegenError{
		err: err,
		msg: msg,
	}
}
