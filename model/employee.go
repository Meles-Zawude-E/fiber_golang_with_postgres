package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Phone    string    `json:"phone"`
	Photo    string    `json:"photo"`
	Role     string    `json:"role"` // Add role field
}
type Employees struct {
	Employees []Employee `json:"employees"`
}

func (employee *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	employee.ID = uuid.New()
	return
}