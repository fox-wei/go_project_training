package model

import (
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	Id        uint64    `gorm:"primary_key;AUTO_INCREMENT;cloumn:id" json:"-"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"-"`
	DeletedAt time.Time `gorm:"column:deltedAt" sql:"index" json:"-"`
}

type AccountInfo struct {
	AccountId   string `json:"id"`
	AccountName string `json:"username"`
	Password    string `json:"password"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type List struct {
	Lock  *sync.RWMutex
	IdMap map[uint64]*AccountInfo
}

//*Token令牌
type Token struct {
	Token string `josn:"token"`
}

type Account struct {
	gorm.Model
	AccountId   string `json:"accountId" gorm:"column:account_id; not null"`
	AccountName string `json:"accountName" gorm:"column:account_name; not null" binding:"required" validate:"min=1,max=32"`
	Password    string `json:"password" gorm:"column:password; not null" binding:"required" validate:"min=5,max=128"`
}
