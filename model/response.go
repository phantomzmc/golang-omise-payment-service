package modelResponse

type CommonResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}