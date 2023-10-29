package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/gowaves/order_automaiton/configs"
	"github.com/gowaves/order_automaiton/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var companyCollection *mongo.Collection = configs.GetCollection(configs.DB, "company")

/*
ClientIDLength defines the length of client IDs for authentication.
ClientSecretLength defines the length of client secrets for authentication.
*/
const (
	ClientIDLength     = 8
	ClientSecretLength = 16
)

/*
generateRandomString generates a random string of the specified length.

Parameters:
  - length: The desired length of the random string.

Returns:
  - The generated random string.
  - An error, if any, during the generation process.
*/
func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

/*
generateClientCredentials generates client credentials, including a client ID and client secret.
The client ID and client secret are random strings with lengths defined by ClientIDLength and ClientSecretLength constants.

Returns:
  - The generated client ID.
  - The generated client secret.
  - An error, if any, during the generation process.
*/
func GenerateClientCredentials() (string, string, error) {
	clientID, err := generateRandomString(ClientIDLength)
	if err != nil {
		return "", "", err
	}

	clientSecret, err := generateRandomString(ClientSecretLength)
	if err != nil {
		return "", "", err
	}

	return clientID, clientSecret, nil
}

/*
GenerateSecretKey generates a secret key based on a client ID and client secret.

Parameters:
  - clientID: The client ID.
  - clientSecret: The client secret.

Returns:
  - The generated secret key.
*/
func GenerateSecretKey(clientID, clientSecret string) string {
	concatenated := clientID + ":" + clientSecret
	// Encode the concatenated string in Base64.
	secretKey := base64.StdEncoding.EncodeToString([]byte(concatenated))

	return secretKey
}

// Get the company credentials just to verify the jwt
func GetCompanyCredentials(mobile string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var company models.Company

	err := companyCollection.FindOne(ctx, bson.M{"mobile": mobile}).Decode(&company)
	if err != nil {
		// Handle the error, e.g., return an error or do something else.
		return "", err
	}
	secretKey := GenerateSecretKey(company.ClientId, company.ClientSecret)

	// Assuming Company struct has fields client_id and client_secret
	return secretKey, nil
}

/*
GetCompanySecretKeys retrieves company secret keys from MongoDB.
It establishes a connection, fetches data, and returns the keys as a map.
Returns:
- A map containing company secret keys
- An error if the retrieval encounters any issues
*/
func GetCompanySecretKeys() (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Fetch data from MongoDB
	var secrets []models.CompanySecret
	cursor, err := companyCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &secrets); err != nil {
		return nil, err
	}

	// Create a map to store fetched data
	companySecretKeys := make(map[string]string)
	for _, secret := range secrets {
		companySecretKeys[secret.CompCode] = secret.Mobile
	}

	return companySecretKeys, nil
}
