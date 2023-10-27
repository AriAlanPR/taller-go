package digimon

import "gorm.io/gorm"

type Digimon struct {
	gorm.Model
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Stage     string `json:"stage"`
	Attribute string `json:"attribute"`
	Picture   string `json:"picture"`
}
