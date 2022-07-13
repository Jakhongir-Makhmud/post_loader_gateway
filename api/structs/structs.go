package structs

type LoadPostsReq struct {
	Pages int64 `json:"pages"`
}

type GetPostReq struct {
	Id int64 `json:"id"`
}

type UpdatePostReq struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type GetListPostsReq struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}
