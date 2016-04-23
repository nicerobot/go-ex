package ex

import (
	"fmt"

	"github.com/pkg/errors"

	nerror "github.com/nicerobot/go-err"
)

//
func ExamplePlainErrors() {
	e := nerror.NewError("ok")
	fmt.Printf("%[1]T %[1]s\n", e)
	errors.Print(e)
	fmt.Printf("%[1]T %#[1]v %[1]s\n", nerror.DoesStuff(e))
	// Output:
	// *nerror.Error ok
	// bool true %!s(bool=true)

}

//
func ExampleWrapErrors() {
	w := errors.Wrap(nerror.NewError("ok"), "wrapper")
	fmt.Printf("%[1]T %[1]s\n", w)
	errors.Print(w)
	fmt.Printf("%[1]T %#[1]v %[1]s\n", nerror.DoesStuff(w))

	// Output:
	// *errors.e wrapper: ok
	// bool true %!s(bool=true)
}
