package param

type OrderSeat struct {
	HotelId  string `json:"hotelId"`
	Persons  int    `json:"persons"`
	DateTime int64  `json:"datetime"`
	RoomType int    `json:"room_type"`
	Mobile   int    `json:"mobile"`
	Name     string `json:"name"`
	Sex      int    `json:"sex"`
	Message  string `json:"message"`
}
