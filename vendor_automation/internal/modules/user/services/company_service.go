package services

import (
	"errors"

	companyModel "github.com/gowaves/vendor_automation/internal/modules/user/models"
	CompanyRepository "github.com/gowaves/vendor_automation/internal/modules/user/repositories"
	"github.com/gowaves/vendor_automation/internal/modules/user/requests/auth"
	CompanyResponse "github.com/gowaves/vendor_automation/internal/modules/user/responses"
	"github.com/mitchellh/mapstructure"
)

type CompanyService struct {
	companyRepository CompanyRepository.CompanyRepositoryInterface
}

func NewCompany() *CompanyService {
	return &CompanyService{
		companyRepository: CompanyRepository.NewCompany(),
	}
}

func (companyService *CompanyService) Create(request auth.CompanyRequest) (CompanyResponse.Company, error) {
	var response CompanyResponse.Company
	var company companyModel.Company

	//Start the mapping code here
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &company,
		TagName: "form", // Use the appropriate struct tag (e.g., json or form)
	})

	if err != nil {
		return response, err
	}

	if err := decoder.Decode(request); err != nil {
		return response, err
	}

	newCompany := companyService.companyRepository.Create(company)
	if newCompany.ID == 0 {
		return response, errors.New("erorr on creating the company")
	}

	return CompanyResponse.ToCompany(newCompany), nil
}

func (companyService *CompanyService) CheckCompanyExists(email string) bool {
	company := companyService.companyRepository.FindByEmail(email)
	return company.ID != 0
}
