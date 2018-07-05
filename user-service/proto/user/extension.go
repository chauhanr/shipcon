package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uid := uuid.New()
	return scope.SetColumn("Id", uid)
}
