package dto

type GetBlogTag struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type GetBlogDTO struct {
	ID         int          `json:"id"`
	Slug       string       `json:"slug"`
	Title      string       `json:"title"`
	Author     string       `json:"author"`
	ContentURL string       `json:"content_url"`
	Tags       []GetBlogTag `json:"tags"`
	CreatedAt  string       `json:"created_at"`
}
