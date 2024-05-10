package schemas

type VideoFetchSchema struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Title string `json:"title"`
}
