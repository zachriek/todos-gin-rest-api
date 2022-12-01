package Models

import "github.com/jinzhu/gorm"

type Todos struct {
	gorm.Model
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func (todo *Todos) TableName() string {
	return "todos"
}
