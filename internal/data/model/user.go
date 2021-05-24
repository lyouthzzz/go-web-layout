package model

type User struct {
	ID       uint64 `gorm:"column:id;PRIMARY_KEY;AUTO_INCREMENT"`
	Username string `gorm:"column:username;type:char(60);NOT NULL"`
	Password string `gorm:"column:password;type:char(60);NOT NULL"`
	Email    string `gorm:"column:email;type:char(60);NOT NULL"`
}

func (User) TableName() string {
	return "user"
}
