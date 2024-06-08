package generrors

import "fmt"

type CodegenError struct {
	err error
	msg string
}

func (generr CodegenError) Error() string {

	if generr.err != nil {
		return fmt.Sprintf("Error: %s \n Hint: %s\n", generr.err, generr.msg)
	}
	if generr.msg == "" {
		return fmt.Sprintf("Error: %s\n", generr.err)
	}
	return fmt.Sprintf("Hint: %s\n", generr.msg)
}
func Err(msg string, err error) CodegenError {
	/* errors.Join() */
	return CodegenError{
		err: err,
		msg: msg,
	}
}
