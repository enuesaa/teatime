package controller

type ListResponse[T any] struct {
	Items []T `json:"items"`
}
func NewListResponse[T any]() ListResponse[T] {
	return ListResponse[T] {
		Items: make([]T, 0),
	}
}

type DescribeResponse[T any] struct {
	Data T `json:"data"`
}
func NewDescribeResponse[T any](data T) DescribeResponse[T] {
	return DescribeResponse[T] {
		Data: data,
	}
}

type IdSchema struct {
	Id string `json:"id"`
}
func NewIdSchema(id string) IdSchema {
	return IdSchema{
		Id: id,
	}
}
type EmptySchema struct {}
func NewEmptySchema() EmptySchema {
	return EmptySchema{}
}
