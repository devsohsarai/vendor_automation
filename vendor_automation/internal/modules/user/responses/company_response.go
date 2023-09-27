package responses

import (
	UserModel "github.com/gowaves/vendor_automation/internal/modules/user/models"
)

type Company struct {
	ID      uint16
	Name    string
	Email   string
	Contact string
	Address string
	Status  bool
}

type Companies struct {
	Data []Company
}

func ToCompany(company UserModel.Company) Company {
	return Company{
		ID:      company.ID,
		Name:    company.Name,
		Email:   company.Email,
		Contact: company.Contact,
		Address: company.Address,
		Status:  company.Status,
	}
}
