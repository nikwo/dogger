package errors

import "errors"

var ErrInvalidLevelString = errors.New("invalid log level string, expect: trace, debug, info, warn, error")
