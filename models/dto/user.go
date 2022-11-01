package dto

type CurrentUserReq struct {
}

type CurrentUserResp struct {
	Success      bool   `json:"success"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`

	Data CurrentUserResData `json:"data"`
}

type CurrentUserResData struct {
	IsLogin   bool   `json:"isLogin"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Userid    string `json:"userid"`
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
