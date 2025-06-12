package common

type successResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

// NewSuccessResponse returns detailed response
func NewSuccessResponse(data, paging, filter interface{}) *successResponse {
	return &successResponse{Data: data, Paging: paging, Filter: filter}
}

// SimpleSuccessResponse is a simple response with just data
func SimpleSuccessResponse(data interface{}) *successResponse {
	return &successResponse{Data: data}
}
