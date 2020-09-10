package vo

type PostVo struct {
	CategoryID uint `json:"category_id"binding:"required"`
	Title string `json:"title"binding:"required,max=30"`
	HeadImg string `json:"head_img"binding:"required"`
	Content string `json:"content"binding:"required"`
}
