package middleware

import (
	"fmt"
	"net/http"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/gowaves/order_automaiton/responses"
	"github.com/gowaves/order_automaiton/utils"
)

// AuthMiddleWare is a middleware function that authenticates JWT tokens using the provided secret key.
// It returns a Fiber handler which verifies JWT tokens with the given secret key.
// Parameters:
//   - secret: The secret key used for JWT token validation.
//
// Returns:
//   - fiber.Handler: The JWT middleware handler configured with the provided secret key.
func AuthMiddleWare(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
	})
}

// CompanyCodeMiddleware generates a Fiber middleware function that extracts the company code from the query parameter.
// It validates the company code against the provided companySecretKeys map and sets the secret key in the context's locals.
// Parameters:
//   - companySecretKeys: A map containing company codes as keys and their corresponding mobile codes.
//
// Returns:
//   - func(*fiber.Ctx) error: The Fiber middleware function that processes the company code validation and sets the secret key.
func CompanyCodeMiddleware(companySecretKeys map[string]string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		compCode := c.Query("comp_code")
		mobileCode, ok := companySecretKeys[compCode]
		if !ok {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid company code")
		}
		secretKey, _ := utils.GetCompanyCredentials(mobileCode)
		c.Locals("secretKey", secretKey)
		return c.Next()
	}
}

// JwtMiddleware generates a Fiber middleware function for handling JSON Web Token (JWT) authentication.
// It retrieves the secret key from Fiber context locals and uses it to validate and process the JWT token.
// Returns:
//   - func(*fiber.Ctx) error: The Fiber middleware function responsible for JWT authentication and validation.
func JwtMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		secretKey := c.Locals("secretKey").(string)
		return AuthMiddleWare(secretKey)(c)
	}
}

/**
 * AuthorizationMiddleware returns a Fiber middleware function that checks user authorization based on 'comp_code' in the JWT token.
 */
func AuthorizationMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenValue := c.Locals("user")

		if token, ok := tokenValue.(*jwt.Token); ok {
			fmt.Println(token)
			fmt.Println(ok)
			compCode := c.Query("comp_code")
			claims := token.Claims.(jwt.MapClaims)
			if comp_code := claims["comp_code"]; comp_code != compCode {
				return c.Status(http.StatusUnauthorized).JSON(responses.AppResponse{
					Status:  http.StatusUnauthorized,
					Message: "Error",
					Data:    &fiber.Map{"data": "You are not authorized to access the other vendor details"},
				})
			}
		}

		/*claims := token.Claims.(jwt.MapClaims)
		compCodeFromToken, exists := claims["comp_code"]
		if !exists {
			return c.Status(http.StatusUnauthorized).JSON(responses.AppResponse{
				Status:  http.StatusUnauthorized,
				Message: "Error",
				Data:    &fiber.Map{"data": "Claim 'comp_code' not found in the token"},
			})
		}

		compCodeFromParam := c.Params("comp_code")
		if compCodeFromToken != compCodeFromParam {
			return c.Status(http.StatusUnauthorized).JSON(responses.AppResponse{
				Status:  http.StatusUnauthorized,
				Message: "Error",
				Data:    &fiber.Map{"data": "You are not authorized to access the other vendor details"},
			})
		}*/

		return c.Next()
	}
}
