package port

type BaseJsonResponse struct {
	Result interface{} `json:"result"`
	Status string      `json:"status"`
}

type BaseJsonMultipleResponse struct {
	Results interface{} `json:"results"`
	Status  string      `json:"status"`
}

type SpotOutputPort interface {
	JsonResponse(result interface{}, status string) *BaseJsonResponse
	JsonMultipleResponse(results interface{}, status string) *BaseJsonMultipleResponse
}
