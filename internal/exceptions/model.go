package exceptions

import (
	"clean_arhitecture_golang/internal/domains"
	"encoding/json"
)

var (
	ErrNotFound = NewAppError(nil, "not found", "")
)

type Error struct {
	Status 			 *domains.StatusResponse `json:"status,omitempty"`
	Err              error  				`json:"-"`
	DeveloperMessage string 				`json:"developer_message,omitempty"`

}

func (e *Error) Error() string {
	return e.Err.Error()
}

func (e *Error) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func (e *Error) Unwrap() error { return e.Err }

func NewAppError(err error, message, developerMessage string) *Error {
	st := domains.NewStatusResponse(-1, message, nil)
	er := &Error{
		Status: 		  st,
		Err:              err,
		DeveloperMessage: developerMessage,
	}
	return er
}

func systemError(err error) *Error {
	return NewAppError(err, "internal system error", err.Error())
}
