package folhacerta

type BaseResponse struct {
	Success   bool   `json:"Success"`
	Error     string `json:"Error"`
	ErrorCode int64  `json:"ErrorCode"`
}
