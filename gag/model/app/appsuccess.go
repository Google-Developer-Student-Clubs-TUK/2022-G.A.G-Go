package app

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
