package mode

//?评论表
type Comment struct {
	ComentId      string `josn:"commentId"`   //*评论ID
	Comment       string `json:"comment"`     //*评论内容
	AccountName   string `json:"accountName"` //*评论人
	AccountPic    string `json:"accountPic"`  //?头像
	Star          int    `json:"star"`        //&评分
	AveragePerson int    `json:"averagePerson"`
	IsGood        int    `json:"isGood"` //?是否优质点评
}
