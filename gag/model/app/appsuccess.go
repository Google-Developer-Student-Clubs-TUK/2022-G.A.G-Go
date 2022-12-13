package app

type Success[T any] struct {
	Code   int
	Msg    string
	Result T
}

func NewSuccess[T any](result T) *Success[T] {

	return &Success[T]{
		Code:   0,
		Msg:    "성공",
		Result: result,
	}
}
