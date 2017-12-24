// Package errors provides ways to add information to Go errors.
package errors

import "fmt"

const (
	GRPCSpace = "google.golang.org/grpc/codes.Code"
	HTTPSpace = "net/http"
	UnixSpace = "syscall.Errno"
)

// Error establishes standard names for common error functionality.
type Error interface {
	error

	// ErrorCode returns a numeric code classifying the error, as well as a space
	// that establishes how to interpret the code. The space must be globally unique.
	// It is recommended to use the fully-qualified name of the Go type defining the
	// codes, or the name of the package where the codes are defined if they are not
	// a separate type. See the constants in this package for common spaces.
	ErrorCode() (space string, code int)

	// ErrorSource returns the error that led to the creation of this error.
	ErrorSource() error

	// ErrorDetails provides more information about the error.
	ErrorDetails() interface{}
}

// Err implements Error.
type Err struct {
	Message string
	Code    int
	Space   string
	Source  error
	Details interface{}
}

func (e Err) Error() string             { return e.Message }
func (e Err) ErrorCode() (string, int)  { return e.Space, e.Code }
func (e Err) ErrorSource() error        { return e.Source }
func (e Err) ErrorDetails() interface{} { return e.Details }

// Printf returns an Err that is identical to e, except that its Message field is
// formatted string.
func (e Err) Printf(format string, args ...interface{}) Err {
	e.Message = fmt.Sprintf(format, args...)
	return e
}

// Source returns the source of the error, if it implements the ErrorSource method.
// Otherwise, Source returns nil.
func Source(err error) error {
	if es, ok := err.(interface {
		ErrorSource() error
	}); ok {
		return es.ErrorSource()
	}
	return nil
}

// RootSource repeatedly calls Source on err. It returns the error for which Source
// returns nil.
func RootSource(err error) error {
	var prev error
	for err != nil {
		prev = err
		err = Source(err)
	}
	return prev
}

// Code returns the result of calling ErrorCode on err, if it implements that method.
// Otherwise, Code returns ("", -1), unless err == nil, in which
// case it returns ("", 0).
func Code(err error) (space string, code int) {
	if err == nil {
		return "", 0
	}
	if ec, ok := err.(interface {
		ErrorCode() (string, int)
	}); ok {
		return ec.ErrorCode()
	}
	return "", -1
}

// Details returns the result of calling ErrorDetails on err, if it implements that
// method. Otherwise, Details returns nil.
func Details(err error) interface{} {
	if ed, ok := err.(interface {
		ErrorDetails() interface{}
	}); ok {
		return ed.ErrorDetails()
	}
	return nil
}
