package errors

import (
	"errors"
)

var ERR_PARAMETER = errors.New("Parameter validation failed.")
var ERR_ABNORMAL_DATA = errors.New("Internal calculation error.")
