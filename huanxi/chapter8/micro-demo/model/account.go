package model

import (
	"github.com/rey/micro-demo/proto/account"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Mobile   string `gorm:"index:idx_mobile;unique;varchar(11;not null)"`
	Password string `gorm:"type:varchar(64);not null"`
	UserName string `gorm:"type:varchar(32)"`
	Gender   string `gorm:"type:varchar(6);default:male"`
	Salt     string `gorm:"type:varchar(64)"`
	Role     int    `gorm:"type:tinyint;defaut:1;comment:'1:user 0:admin'"`
}

// 更改gorm自动迁移时的表名
func (Account) TableName() string {
	return "micro_account"
}

// Account -> pb.AccountInfo
func (a Account) ToPBAccountInfo() (pa *account.AccountInfo) {
	pa = &account.AccountInfo{}
	pa.Id = int64(a.ID)
	pa.Mobile = a.Mobile
	pa.Password = a.Password
	pa.Username = a.UserName
	pa.Gender = a.Gender

	return
}
