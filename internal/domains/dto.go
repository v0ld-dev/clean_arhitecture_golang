package domains

type StatusResponse struct {
	Status int `json:"status"`
	StatusText string `json:"status_text"`
	Data interface{} `json:"data"`
}

func NewStatusResponse(status int, statusText string, lowLevelResp interface{}) *StatusResponse {
	return &StatusResponse{
		Status: status,
		StatusText: statusText,
		Data: lowLevelResp,
	}
}
