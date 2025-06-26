package errs

import "fmt"

var (
	NotFound     = fmt.Errorf("resource not found")
	InvalidCreds = fmt.Errorf("invalid credentials")
	NotAllowed   = fmt.Errorf("not allowed")
	Internal     = fmt.Errorf("internal server error")
)
