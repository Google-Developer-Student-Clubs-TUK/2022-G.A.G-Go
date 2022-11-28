package model

type ApiResponse[T any] struct {
	ResultCode int
	Msg        string
	Result     T
}
