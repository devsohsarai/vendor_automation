package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	ID           primitive.ObjectID `json:"id,omitempty"`
	CompanyName  string             `json:"company_name,omitempty" validate:"required"`
	Address      string             `json:"address,omitempty" validate:"required"`
	OwnerName    string             `json:"owner_name,omitempty" validate:"required"`
	Mobile       string             `json:"mobile,omitempty" validate:"required"`
	ClientId     string             `json:"client_id,omitempty" bson:"client_id,omitempty"`
	ClientSecret string             `json:"client_secret,omitempty" bson:"client_secret,omitempty"`
	CompCode     string             `json:"comp_code,omitempty" bson:"comp_code,omitempty"`
}

type AuthRequest struct {
	ClientId     string `json:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
}

// CompanySecret represents the structure of the data you want to retrieve
type CompanySecret struct {
	CompCode string `bson:"comp_code"`
	Mobile   string `bson:"mobile"`
}
