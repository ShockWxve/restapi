package userService

import (
	"github.com/shockwxve/restapi/internal/taskService"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string             `gorm:"uniqueIndex;not null" json:"email"`
	Password string             `gorm:"not null" json:"password"`
	Tasks    []taskService.Task `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"tasks"`
}
