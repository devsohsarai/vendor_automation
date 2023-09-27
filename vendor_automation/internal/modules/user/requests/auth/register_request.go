package auth

type RegisterRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	Email     string `form:"email" binding:"required,email,min=3,max=100"`
	Password  string `form:"password" binding:"required,min=8"`
	Contact   string `form:"contact" binding:"required,min=10,max=20"`
	CompanyID uint16 `form:"company_id" binding:"required"`
}
