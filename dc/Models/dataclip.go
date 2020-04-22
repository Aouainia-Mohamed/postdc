package Models

import "github.com/jinzhu/gorm"

type Dataclip struct {
	gorm.Model

	Dataclipkey string `json:"dataclipkey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Sqltext     string `json:"sqltext"`
	Nargs       int    `json:"nargs"`
	Argsdesc    string `json:"argsdesc"`
}
