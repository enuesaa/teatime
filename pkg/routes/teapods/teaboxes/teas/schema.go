package teas

type CreateRequestBody struct {
	Name string `json:"name"` // command name
}

type Creation struct {
	Id string `json:"id"`
}

type Item struct {
	Id   string                 `json:"id"`
	CreatedAt string            `json:"createdAt"`
	UpdatedAt string            `bson:"updatedAt"`
	Data map[string]interface{} `json:"data"`
}

type DeleteResponse struct{}
