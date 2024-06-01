package plug

import "fmt"

// see https://github.com/hashicorp/go-plugin/blob/8d2aaa458971cba97c3bfec1b0380322e024b514/error.go#L11
type Result[T any] struct {
	Data   T
	HasErr bool
	ErrMsg string
}

func (r *Result[T]) Err() error {
	if r.HasErr {
		return fmt.Errorf(r.ErrMsg)
	}
	return nil
}

func NewResult[T any](data T, err error) Result[T] {
	if err != nil {
		return Result[T]{
			HasErr: true,
			ErrMsg: err.Error(),
		}
	}

	return Result[T]{
		Data: data,
		HasErr: false,
	}
} 