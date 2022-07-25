package api

import "fmt"

func (r *Error) Error() string {
	return fmt.Sprintf("err %v", r.Err)
}
