package model

type successRes struct {
	Data any `json:"data"`
}

func NewSuccessResponse(data any) *successRes {
	return &successRes{Data: data}
}
