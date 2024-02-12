package plug

// info
type Info struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// data
type Row struct {
	Id     string `json:"id"`
	Values map[string]Value `json:"values"`
}
type Value struct {
	Type    string `json:"type"` // str, int or bool
	StrVal  string `json:"strVal"`
	IntVal  int `json:"intVal"`
	BoolVal string `json:"values"`
}
