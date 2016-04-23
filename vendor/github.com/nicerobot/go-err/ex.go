// http://dave.cheney.net/paste/gocon-spring-2016.pdf
package nerror

import (
	"fmt"

	"github.com/pkg/errors"
)

//
type Error struct {
	Stuff string
}

//
type stuff interface {
	DoesStuff() bool
}

//
func DoesStuff(err error) bool {
	te, ok := errors.Cause(err).(stuff)
	return ok && te.DoesStuff()
}

//
func (e *Error) Error() string {
	return fmt.Sprintf("%s", e.Stuff)
}

//
func (e *Error) String() string {
	return fmt.Sprintf("%s", e.Error())
}

//
func (e *Error) DoesStuff() bool {
	// The result might/probably depends on the Error's context.
	return true
}

//
func NewError(stuff string) *Error {
	return &Error{stuff}
}
