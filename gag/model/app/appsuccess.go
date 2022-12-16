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

type SuccessPaging[T any] struct {
	Code   int          `json:"code"`
	Msg    string       `json:"msg"`
	Result T            `json:"result"`
	Paging model.Paging `json:"paging"`
}

func NewSuccessPaging[T any](result T, paging model.Paging) *SuccessPaging[T] {

	return &SuccessPaging[T]{
		Code:   1,
		Msg:    "성공",
		Result: result,
		Paging: paging,
	}
}
