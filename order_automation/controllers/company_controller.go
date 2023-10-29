package controllers

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gowaves/order_automaiton/configs"
	"github.com/gowaves/order_automaiton/utils"

	"net/http"
	"time"

	"github.com/gowaves/order_automaiton/models"
	"github.com/gowaves/order_automaiton/responses"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var companyCollection *mongo.Collection = configs.GetCollection(configs.DB, "company")
var validate = validator.New()

func AuthProcess(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var authRequest models.AuthRequest

	// Parse the login request
	if err := c.BodyParser(&authRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}
	var company models.Company

	// Define a filter that checks both client_id and client_secret
	filter := bson.M{
		"client_id":     authRequest.ClientId,
		"client_secret": authRequest.ClientSecret,
	}

	// Attempt to find a document matching both conditions
	err := companyCollection.FindOne(ctx, filter).Decode(&company)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responses.UserResponse{
			Status:  http.StatusUnauthorized,
			Message: "error",
			Data:    &fiber.Map{"data": "Invalid client_id or client_secret"},
		})
	}

	// Generate a JWT token\
	expirationTime := time.Now().Add(30 * time.Minute)

	// Create the Claims
	claims := jwt.MapClaims{
		"owner_name": company.OwnerName,
		"comp_code":  company.CompCode,
		"mobile":     company.Mobile,
		"exp":        expirationTime.Unix(),
	}

	// Convert the Unix timestamp to a human-readable format
	expirationTimeFormatted := expirationTime.Format("2006-01-02 15:04:05")

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := utils.GenerateSecretKey(authRequest.ClientId, authRequest.ClientSecret)

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: &fiber.Map{
			"token":           tokenString,
			"expiration_time": expirationTimeFormatted,
		},
	})
}

func CreateCompany(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	//validate the request body
	var company models.Company
	if err := c.BodyParser(&company); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&company); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	clientId, clientSecret, err := utils.GenerateClientCredentials()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	newCompany := models.Company{
		ID:           primitive.NewObjectID(),
		CompanyName:  company.CompanyName,
		Address:      company.Address,
		OwnerName:    company.OwnerName,
		Mobile:       company.Mobile,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		CompCode:     company.CompCode,
	}

	// Check if a document with the same mobile number already exists
	existingDocument := companyCollection.FindOne(ctx, bson.M{"mobile": newCompany.Mobile})
	if existingDocument.Err() == nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": "Company already exists"},
		})
	}

	result, err := companyCollection.InsertOne(ctx, newCompany)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAllCompanies(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use the Find method with no filter to get all companies
	results, err := companyCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()},
		})
	}
	defer results.Close(ctx)

	var companies []models.Company

	for results.Next(ctx) {
		var singleCompany models.Company
		if err := results.Decode(&singleCompany); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
		}
		companies = append(companies, singleCompany)
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &fiber.Map{"data": companies},
	})
}

func GetCompanyDetails(c *fiber.Ctx) error {
	mobile := c.Params("mobile")
	secretKey, err := utils.GetCompanyCredentials(mobile)

	if err != nil {
		// Handle the error, e.g., return an error response or log it
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Company not found",
		})
	}

	return c.SendString(secretKey)
}

func Stack(c *fiber.Ctx) error {
	tokenValue := c.Locals("user")

	if token, ok := tokenValue.(*jwt.Token); ok {
		claims := token.Claims.(jwt.MapClaims)
		name := claims["owner_name"]
		fmt.Println(name)
		for key, value := range claims {
			fmt.Printf("Claim: %s, Value: %v\n", key, value)
		}
	} else {
		fmt.Println("Token not found or has an invalid type in c.Locals")
	}

	return c.SendString("Welcome ")
}
