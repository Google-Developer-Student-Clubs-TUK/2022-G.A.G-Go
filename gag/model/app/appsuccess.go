package app

import (
	"gag.com/model"
)

type Success[T any] struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result T      `json:"result"`
}

func NewSuccess[T any](result T) *Success[T] {

	return &Success[T]{
		Code:   1,
		Msg:    "성공",
		Result: result,
	}
}

type SuccessPagination[T any] struct {
	Code       int              `json:"code"`
	Msg        string           `json:"msg"`
	Result     T                `json:"result"`
	Pagination model.Pagination `json:"pagination"`
}

func NewSuccessPagination[T any](result T, pagination model.Pagination) *SuccessPagination[T] {

	return &SuccessPagination[T]{
		Code:       1,
		Msg:        "성공",
		Result:     result,
		Pagination: pagination,
	}
}
