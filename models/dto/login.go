package dto

type LoginResp struct {
	Status string `json:"status"`
	CurrentAuthority string `json:"currentAuthority"`
	Type string `json:"type"`
}

type LoginRequest struct {
	AutoLogin bool `json:"autoLogin"`
	Password string `json:"password"`
	Type string `json:"type"`
	Username string `json:"username"`
}

