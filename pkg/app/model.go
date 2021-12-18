package app

type Field struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Dt      int    `json:"dt"`
	Tip     string `json:"tip"`
	Isnum   int    `json:"isnum"`
	Options interface{} `json:"options"`
	R       int    `json:"r"`
	Code    int    `json:"code,omitempty"`
}
