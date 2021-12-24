package api

import (
	"encoding/json"
)

var (
	ErrNotFound = NewStatusResponse(-1, "not found", nil, nil, "")
)

type StatusResponse struct {
	Status           int           `json:"status"`
	StatusText       string        `json:"status_text"`
	Data             interface{}   `json:"data"`
	Err              error         `json:"-"`
	DeveloperMessage string        `json:"developer_message,omitempty"`
}

func NewStatusResponse(status int, statusText string, lowLevelResp interface{}, err error, developerMessage string) *StatusResponse {
	return &StatusResponse{
		Status:           status,
		StatusText:       statusText,
		Data:             lowLevelResp,
		Err:              err,
		DeveloperMessage: developerMessage,
	}
}

func (e *StatusResponse) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func SystemError(err error) *StatusResponse {
	return NewStatusResponse(-1, "internal system error",nil, err, err.Error())
}

func (e *StatusResponse) Error() string {
	return e.Err.Error()
}
