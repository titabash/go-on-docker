package presenter

import "app/usecase/port"

type BaseJsonPresenter struct {
}

func NewBaseJsonPresenter() *BaseJsonPresenter {
	return &BaseJsonPresenter{}
}

func (bjp *BaseJsonPresenter) JsonResponse(result interface{}, status string) *port.BaseJsonResponse {
	response := &port.BaseJsonResponse{
		Result: result,
		Status: status,
	}
	return response
}

func (bjp *BaseJsonPresenter) JsonMultipleResponse(results interface{}, status string) *port.BaseJsonMultipleResponse {
	response := &port.BaseJsonMultipleResponse{
		Results: results,
		Status:  status,
	}
	return response
}
