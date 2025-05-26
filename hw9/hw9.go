package hw9

import (
	"errors"
)

// SpecificError is a predefined error for specific cases.
var SpecificError = errors.New("specific error")

func someError(i int) error {
	switch i {
	case 1:
		return nil
	case 2:
		return errors.New("error")
	case 3:
		return errors.New("error")
	case 4:
		return SpecificError // Return the predefined error instance
	case 5:
		return nil
	default:
		return nil
	}
}
