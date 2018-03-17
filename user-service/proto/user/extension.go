package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// BeforeCreate make use of satori/go.uuid for uniq ID on user
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("Id", uuid.String())
}
