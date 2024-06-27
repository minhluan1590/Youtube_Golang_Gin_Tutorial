package entity

type Person struct {
	FirstName string `json:"first_name" binding:"required,min=2,max=50" validate:"is-cool"`
	LastName  string `json:"last_name" binding:"required,min=2,max=50"`
	Email     string `json:"email" binding:"required,email"`
	Age       int    `json:"age" binding:"required,min=1,max=120"`
}

type Video struct {
	Title       string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=10,max=500"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
