package services

import (
	"github.com/gowaves/vendor_automation/internal/modules/user/requests/auth"
	UserResponse "github.com/gowaves/vendor_automation/internal/modules/user/responses"
)

type CompanyServiceInterface interface {
	Create(request auth.CompanyRequest) (UserResponse.Company, error)
	CheckCompanyExists(email string) bool
}
