package request

type User struct {
	ID       int64  `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) IsValid() bool {
	return false
}
