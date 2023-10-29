package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gowaves/order_automaiton/models"
	"github.com/gowaves/order_automaiton/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client         *mongo.Client
	userCollection *mongo.Collection
	validate       = validator.New()
	mongodbURI     = "mongodb+srv://admin:ZwHO5SI9FDcY2Zcc@testing.mo3h4vw.mongodb.net/testfiber?retryWrites=true&w=majority" // Change this to your MongoDB URI
	databaseName   = "testfiber"
	collectionName = "user"
)

func main1() {
	// Initialize Fiber
	app := fiber.New()

	// Initialize MongoDB client
	initMongoDB()

	// Define route to create a user
	app.Post("/create-user", CreateUser)

	// Start the server
	app.Listen(":6000")
}

func initMongoDB() {
	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI(mongodbURI)

	// Create MongoDB client
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	// Check if the MongoDB client is connected
	if err := client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}

	// Get a reference to the user collection
	userCollection = client.Database(databaseName).Collection(collectionName)
}

// CreateUser creates a new user in MongoDB
func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User

	// Validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// Use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newUser := models.User{
		ID:       primitive.NewObjectID(),
		Name:     user.Name,
		Location: user.Location,
		Title:    user.Title,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

// Define your User and responses structs as needed

// models.User struct example:
// type User struct {
//     ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
//     Name     string             `json:"name" validate:"required"`
//     Location string             `json:"location"`
//     Title    string             `json:"title"`
// }

// responses.UserResponse struct example:
// type UserResponse struct {
//     Status  int         `json:"status"`
//     Message string      `json:"message"`
//     Data    interface{} `json:"data"`
// }
