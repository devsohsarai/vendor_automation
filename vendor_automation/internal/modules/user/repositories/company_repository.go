package repositories

import (
	companyModel "github.com/gowaves/vendor_automation/internal/modules/user/models"
	"github.com/gowaves/vendor_automation/pkg/database"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	DB *gorm.DB
}

func NewCompany() *CompanyRepository {
	return &CompanyRepository{
		DB: database.Connection(),
	}
}

func (companyRepository *CompanyRepository) Create(user companyModel.Company) companyModel.Company {
	var newCompany companyModel.Company
	companyRepository.DB.Create(&user).Scan(&newCompany)
	return newCompany
}

func (companyRepository *CompanyRepository) FindByEmail(email string) companyModel.Company {
	var company companyModel.Company
	companyRepository.DB.First(&company, "email = ?", email)
	return company
}

func (companyRepository *CompanyRepository) FindByID(id int) companyModel.Company {
	var company companyModel.Company
	companyRepository.DB.First(&company, "id = ?", id)
	return company
}
