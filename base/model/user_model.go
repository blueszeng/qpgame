package model

import (
	"fmt"
	//	"fmt"
	"time"

	"github.com/name5566/leaf/gate"
)

type User struct {
	BaseModel
	//ID        int `gorm:"primary_key"`
	UserName  string
	IsOnline  int `gorm:"-"` // Ignore this field
	RoomID    int `gorm:"-"`
	RoomCards int
	Status    int
	Nickname  string
	Type      int
	Password  string
	CreatedAt time.Time
	Agent     gate.Agent `gorm:"-"`
	/*Name     string `gorm:"size:255"` // Default size for string is 255, reset it with this tag
	Num      int    `gorm:"AUTO_INCREMENT"`

	CreditCard CreditCard // One-To-One relationship (has one - use CreditCard's UserID as foreign key)
	Emails     []Email    // One-To-Many relationship (has many - use Email's UserID as foreign key)

	BillingAddress   Address // One-To-One relationship (belongs to - use BillingAddressID as foreign key)
	BillingAddressID sql.NullInt64

	ShippingAddress   Address // One-To-One relationship (belongs to - use ShippingAddressID as foreign key)
	ShippingAddressID int

	IgnoreMe  int        `gorm:"-"`                         // Ignore this field
	Languages []Language `gorm:"many2many:user_languages;"` // Many-To-Many relationship,*/
}

func (User) TableName() string {
	return "lf_user"
}

func (user *User) GetOne(where map[string]string) User {
	db := DB
	if len(where) > 0 {
		for k, v := range where {
			db = db.Where(k, v)
		}
	}
	db.First(user)

	fmt.Println(user.BaseModel.ID)
	return *user
}
