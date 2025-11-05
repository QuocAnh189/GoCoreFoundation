package login

type LoginReq struct {
	LoginName   string `json:"login_name"`
	RawPassword string `json:"password"`
}

type LoginRes struct {
	AuthToken string `json:"auth_token"`
}
