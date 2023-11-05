package plug

// info
type Info struct {
	Name        string `json:"info"`
	Description string `json:"description"`
	Cards       []string `json:"cards"`
	PanelMap    map[string][]string `json:"panelMap"` // per card name
}

// ui
type Card struct {
	Enable   bool `json:"enable"`
	Layout   string `json:"layout"` // textCard
	TextCard TextCardConfig `json:"textCard"`
}
type TextCardConfig struct {
	Heading string `json:"heading"`
	Center  bool `json:"center"`
	Text    string `json:"text"`
}
type Panel struct {
	Enable     bool `json:"enable"`
	Layout     string `json:"layout"` // tablePanel
	TablePanel TablePanelConfig `json:"tablePanel"`
}
type TablePanelConfig struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Head        []string `json:"head"`
	Records     []TablePanelRecord `json:"records"` // パスを格納しているだけ
	Creation    TablePanelCreation `json:"creation"`
}
type TablePanelRecord struct {
	Model  string `json:"model"` // like `notes`
	Name   string `json:"name"` // like `main`
	Values []TablePanelRecordValue `json:"values"`
}
type TablePanelRecordValue struct {
	Readonly bool `json:"readonly"`
	Key      string `json:"key"`
}
type TablePanelCreation struct {
	Enable bool `json:"enable"`
	ModelName string `json:"modelName"`
}

// data
type Record struct {
	Enable bool `json:"enable"`
	Name   string `json:"name"`
	Values map[string]Value `json:"values"`
}
type Value struct {
	Type    string `json:"type"` // str, int or bool
	StrVal  string `json:"strVal"`
	IntVal  int `json:"intVal"`
	BoolVal string `json:"values"`
}
