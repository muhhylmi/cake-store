package web

type CakeGetRequest struct {
	CakeId int `params:"cakeId"`
}

type CakeListRequest struct {
	Keyword string `query:"q"`
	Size    int    `query:"size"`
	Page    int    `query:"page"`
}
