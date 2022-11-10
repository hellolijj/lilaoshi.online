package dto

type UserInfoReq struct {
}

type CurrentUserResp struct {
	Success      bool   `json:"success"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`

	Data CurrentUserData `json:"data"`
}

type CurrentUserData struct {
	IsLogin   bool   `json:"isLogin"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	UserId    string `json:"userid"`
	Email     string `json:"email"`
	Signature string `json:"signature"`
	Title     string `json:"title"`
	Group     string `json:"group"`
	Tags      []struct {
		Key   string `json:"key"`
		Label string `json:"label"`
	} `json:"tags"`
	NotifyCount int    `json:"notifyCount"`
	UnreadCount int    `json:"unreadCount"`
	Country     string `json:"country"`
	Access      string `json:"access"`
	Geographic  struct {
		Province struct {
			Label string `json:"label"`
			Key   string `json:"key"`
		} `json:"province"`
		City struct {
			Label string `json:"label"`
			Key   string `json:"key"`
		} `json:"city"`
	} `json:"geographic"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name string `json:"name"`
	Mobile string `json:"mobile"`
	AsAdmin bool `json:"asAdmin"`
}

type RegisterResp struct {
	Status string `json:"status"`
	ErrMessage string `json:"errMessage"`
}

type ListResp struct {
	Current int64 `json:"current"`
	PageSize int64 `json:"pageSize"`
	Total int `json:"total"`
	Success bool `json:"success"`
	Data interface{} `json:"data"`
}
