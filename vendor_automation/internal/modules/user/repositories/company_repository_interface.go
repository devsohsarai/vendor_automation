package repositories

import companyModel "github.com/gowaves/vendor_automation/internal/modules/user/models"

type CompanyRepositoryInterface interface {
	Create(company companyModel.Company) companyModel.Company
	FindByEmail(email string) companyModel.Company
	FindByID(id int) companyModel.Company
}
