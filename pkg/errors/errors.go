package errors

import "github.com/dwarvesf/gerr"

var ErrInvalidProduct = gerr.New(40001, "product is invalid")
