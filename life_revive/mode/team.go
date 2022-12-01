package mode

//?团购
type Team struct {
	TeamId    string `json:"teamId"`
	HotelId   string `json:"hotelId"`
	TeamName  string `json:"teamName"`
	Price     int    `json:"price"`     //*原价
	TeamPrice int    `json:"teamPrice"` //*团购价格
	SoldNum   string `json:"soldNum"`
}
