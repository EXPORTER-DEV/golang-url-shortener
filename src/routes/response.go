package routes

type Response[D any] struct {
	Error bool `json:"error"`
	Data  *D   `json:"data"`
}

type ErrorResponse struct {
	Error       bool   `json:"error"`
	Description string `json:"description"`
}

func NewResponse[D any](
	err bool,
	data *D,
) *Response[D] {

	return &Response[D]{
		Error: err,
		Data:  data,
	}

}

func NewErrorResponse(
	description string,
) *ErrorResponse {
	return &ErrorResponse{
		Error:       true,
		Description: description,
	}
}
