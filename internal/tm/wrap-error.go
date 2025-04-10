package tm

import (
	"errors"
	"fmt"
)

// wrapErr returns an error that wraps the original error with additional context.
// If the original error is of type UserErr, it does not wrap and returns it as is.
func wrapErr(err error, errContext string) error {
	var userErr UserErr
	if errors.As(err, &userErr) {
		return err
	}
	return fmt.Errorf("%s: %w", errContext, err)
}
