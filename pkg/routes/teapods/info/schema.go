package info

type Item struct {
	Name        string       `json:"name"`
	Command     string       `json:"command"`
	Description string       `json:"description"`
	Teaboxes    []ItemTeabox `json:"teaboxes"`
}

type ItemTeabox struct {
	Name    string `json:"name"`
}
