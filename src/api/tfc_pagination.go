package api

type Pagination struct {
	CurrentPage int  `json:"current-page"`
	NextPage    *int `json:"next-page"`
	TotalPages  int  `json:"total-pages"`
	PrevPage    *int `json:"prev-page"`
	TotalCount  int  `json:"total-count"`
	PerPage     int  `json:"per-page"`
}

type Meta struct {
	Pagination Pagination `json:"pagination"`
}
