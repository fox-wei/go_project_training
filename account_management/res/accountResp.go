package res

type AccountResp struct {
	AccountId   string `json:"accountId"`
	AccountName string `json:"accountName"`
}

type ListReponse struct {
	Username    any            `json:"username"`
	TotalCount  uint64         `json:"totalCount"`
	AccountList []*AccountResp `json:"acountList"`
}
