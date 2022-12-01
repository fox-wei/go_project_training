package mode

type OrderSeat struct {
	OrderSeatId string `json:"orderId"`
	HotelId     string `json:"hotelId"`
	Persons     int    `json:"persons"`   //?订座人数
	DateTime    int64  `json:"datetime"`  //*预订时间
	Mobile      int    `json:"mobile"`    //*手机号
	RoomType    int    `json:"room_Type"` //*房间类型
	Name        string `json:"name"`      //*联系人名字
	Sex         int    `json:"sex"`
	Message     string `json:"message"` //?留言
}
