package domain

// 领域模型
type User struct {
	Id       uint64
	Username string
	Password string
	Email    string
}

func (u *User) IsAdmin() bool {
	return u.Username == "admin"
}
