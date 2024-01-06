package biz

type (
	// User 用户信息
	User struct {
		Uid      int64
		Password string
		Username string
		Email    string
	}
	// Users 用户列表
	Users []*User
	// UserPage 用户分页
	UserPage struct {
		Total int64
		Items Users
	}
)
