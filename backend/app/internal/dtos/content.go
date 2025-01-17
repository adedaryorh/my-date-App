package dtos

type PagedContent struct {
	Page  int       `json:"page"`
	Limit int       `json:"limit"`
	Posts []Content `json:"contnet"`
}

type Content struct {
	Type         string        `json:"type"`
	Title        string        `json:"title"`
	Post         *Post         `json:"post"`
	Celebration  *Celebration  `json:"celebration"`
	Celebrations []Celebration `json:"celebrations"`
}
