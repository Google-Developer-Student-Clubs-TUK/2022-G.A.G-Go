package model

type ApiResponse[T any] struct {
	Code   int
	Msg    string
	Result T
}

type Result struct {
	Success string
}
