package param

type AccountRequest struct {
	AccountName string `form:"accountName" json:"accountName"`
	Password    string `form:"password" json:"password"`
}

type AccountList struct {
	AccountName string `json:"accountName"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
}
