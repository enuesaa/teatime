package info

type Item struct {
	Name        string       `json:"name"`
	Command     string       `json:"command"`
	Description string       `json:"description"`
	Teaboxes    []ItemTeabox `json:"teaboxes"`
	Actions     []ItemAction `json:"actions"`
}

type ItemTeabox struct {
	Name        string        `json:"name"`
	Placeholder string        `json:"placeholder"`
	Inputs      []TeaboxInput `json:"inputs"`
}

type TeaboxInput struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ItemAction struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}
