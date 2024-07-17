package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type User struct {
	Id        int       `gorm:"type:int;primary_key;AUTO_INCREMENT" json:"id"`
	Role      string    `gorm:"type:varchar(10);not null" json:"role"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(50);not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tasks     []Task    `gorm:"constraint:OnDelete:CASCADE" json:"tasks,omitempty"` // has many
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	tx.Clauses(clause.Returning{}).Where("user_id = ?", u.Id).Delete(&Task{})
	return
}
