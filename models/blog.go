package models

// BlogPost defines the blog post model
type BlogPost struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Descriptions string `json:"descriptions"`
	Body         string `json:"body"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
