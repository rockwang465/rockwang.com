package vo

type PostVo struct {
	CategoryID uint   `json:"category_id" binding:"required"`
	Title      string `json:"title" binding:"required,max=30"`
	HeadImage  string `json:"head_image"`
	Content    string `json:"content" binding:"required"`
}
