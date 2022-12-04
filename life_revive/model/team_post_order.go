package model

//*团购下单
type TeamPostOrder struct {
	TeamPostId   string `json:"teamPostOrderId"`
	TeamDetailId string `json:"teamDetailId"` //^团购详情id
	RealPrice    int    `json:"realPrice"`
	Quantity     int    `json:"quantity"` //^数量
	Mobile       string `json:"mobile"`
	Name         string `json:"name"`
	Sex          int    `json:"sex"`
	Message      string `json:"message"` //?附加信息
}
