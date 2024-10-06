package teapods

type Item struct {
	Name string `json:"name"`
}

type CreateRequestBody struct {
	Name string `json:"name"` // command name
}

type Creation struct {
	Ok bool `json:"ok"`
}
