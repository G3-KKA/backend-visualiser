package generrors

import "fmt"

type CodegenError struct {
	err error
	msg string
}

func (generr CodegenError) Error() string {

	if generr.err == nil {
		return fmt.Sprintf("Hint: %s\n", generr.msg)
	}
	if generr.msg == "" {
		return fmt.Sprintf("Error: %s\n", generr.err)
	}
	return fmt.Sprintf("Error: %s \n Hint: %s\n", generr.err, generr.msg)
}
func Err(msg string, err error) CodegenError {
	return CodegenError{
		err: err,
		msg: msg,
	}
}

// TODO: Replace Errors with this
func Full(msg string, err error) CodegenError {
	return CodegenError{
		err: err,
		msg: msg,
	}
}
func Message(msg string) CodegenError {
	return CodegenError{
		err: nil,
		msg: msg,
	}
}

//	Name conflict with old Err(msg,err) function
/*
func Err(err error) CodegenError {
	return CodegenError{
		err: err,
		msg: "",
	}
}
*/
