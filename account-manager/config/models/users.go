package models

type ReqRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RespRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RespLogin struct {
	Token string `json:"token"`
}

type ReqUpdateUserProfile struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
}

type RespCurrentUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
