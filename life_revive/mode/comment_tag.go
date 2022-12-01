package mode

type CommentTag struct {
	TagId   string `json:"tagId"`
	TagName string `json:"tagName"`
	TagNum  int    `json:"tagNum"` //?tag次数
	HoteId  string `json:"hotelId"`
}
