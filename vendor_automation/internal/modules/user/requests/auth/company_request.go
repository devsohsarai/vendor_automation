package auth

type CompanyRequest struct {
	Name    string `form:"name" binding:"required,min=3,max=100"`
	Email   string `form:"email" binding:"required,email,min=3,max=320"`
	Contact string `form:"contact" binding:"required,min=10,max=20"`
	Address string `form:"address" binding:"required,min=8"`
}
