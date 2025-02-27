package domain

import "bank-poc/src/internal/utils"

type ApiResponse struct {
	Successful bool
	Message    string
	Data       any
	Code       string
}

func (res ApiResponse) Failure(msg string) ApiResponse {
	res.Code = utils.FailureResponseCode
	res.Data = nil
	res.Message = msg
	res.Successful = false

	return res
}

func (res ApiResponse) Success(msg string, data any) ApiResponse {
	res.Code = utils.SuccessResponseCode
	res.Data = data
	res.Message = msg
	res.Successful = true

	return res
}
