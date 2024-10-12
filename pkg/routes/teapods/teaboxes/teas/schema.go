package teas

type CreateRequestBody struct {
	Name string `json:"name"` // command name
}

type Creation struct {
	Id string `json:"id"`
}
