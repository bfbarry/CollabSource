package errors

import (
	native "errors"
	"net/http"
	"strings"
	"fmt"
)

type Op string
type Error struct {
	err error
	code int
	op Op
	message string
}

func E(	err error, code int, op Op, message string) *Error {
	e := &Error{
		err: err,
		op: op,
		code: http.StatusInternalServerError,
		message: message,
	}
	return e
}

//Error as string format for debugging (server)
func (this *Error) Error() string {
	b := new(strings.Builder)
	b.WriteString(fmt.Sprintf("Error in %s: ", string(this.op)))

	if this.err != nil {
		b.WriteString(this.err.Error())
	}

	return b.String()
}

//used for json returned to client
func (this *Error) ClientMessage() string {

	// if this.code >= http.StatusInternalServerError {
	// 	return &M{msg: "Internal server error"}
	// }
	if this.message == "" {
		switch this.code {
		case http.StatusBadRequest:
		case http.StatusUnauthorized:
			return "Unauthorized"
		case http.StatusForbidden:
			return "You do not have permission to perform this action"
		case http.StatusNotFound:
			return "The requested resource was not found"
		case http.StatusUnsupportedMediaType:
			return "Unsupported content-type"
		default:
			return "Something went wrong"
		}
	}
	return this.message

}

func (this *Error) Status() int {
	if this.code >= http.StatusBadRequest {
		return this.code
	}
	return http.StatusInternalServerError
}

////
//native error overrides
////

func (this *Error) Unwrap() error {
	return this.err
}

func (this *Error) Newf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

func New(s string) error {
	return native.New(s)
}

func Is(err, target error) bool {
	return native.Is(err, target)
}

func As(err error, target interface{}) bool {
	return native.As(err, target)
}


