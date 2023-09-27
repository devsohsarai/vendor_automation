package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gowaves/vendor_automation/internal/modules/user/requests/auth"

	UserService "github.com/gowaves/vendor_automation/internal/modules/user/services"
	"github.com/gowaves/vendor_automation/pkg/converters"
	"github.com/gowaves/vendor_automation/pkg/errors"
	"github.com/gowaves/vendor_automation/pkg/htmlparse"
	"github.com/gowaves/vendor_automation/pkg/old"
	"github.com/gowaves/vendor_automation/pkg/sessions"
)

type Controller struct {
	userService    UserService.UserServiceInterface
	companyService UserService.CompanyServiceInterface
}

func New() *Controller {
	return &Controller{
		userService:    UserService.New(),
		companyService: UserService.NewCompany(),
	}
}

/*
Create new company for grow the business.
Parameters:
  - c (*gin.Context): The Gin context object for handling HTTP requests and responses.
*/
func (controller *Controller) Company(c *gin.Context) {
	htmlparse.Render(c, http.StatusOK, "modules/user/html/company", gin.H{
		"title": "Company",
	})
}

/*
HandleCompany handles the company  process based on the provided request.
Parameters:
  - c (*gin.Context): The Gin context object representing the HTTP request and response.
*/
func (controller *Controller) HandleCompany(c *gin.Context) {
	// validate the request
	var request auth.CompanyRequest

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/company")
		return
	}

	if controller.companyService.CheckCompanyExists(request.Email) {
		errors.Init()
		errors.Add("Email", "Company Emailaddress already exists!")
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/company")
		return
	}

	//create the user
	company, err := controller.companyService.Create(request)

	//Check if there is any error on the user page
	if err != nil {
		c.Redirect(http.StatusFound, "/company")
		return
	}

	// after creating the user redirect to the home page
	log.Printf("The company created successfully with a name %s \n", company.Name)
	c.Redirect(http.StatusFound, "/")
}

/*
Register handles user registration by rendering the "register" template.
Parameters:
  - c (*gin.Context): The Gin context object for handling HTTP requests and responses.
*/
func (controller *Controller) Register(c *gin.Context) {

	htmlparse.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register",
	})
}

/*
HandleRegister handles the user registration process based on the provided request.
Parameters:
  - c (*gin.Context): The Gin context object representing the HTTP request and response.
*/
func (controller *Controller) HandleRegister(c *gin.Context) {
	// validate the request
	var request auth.RegisterRequest

	//Set the company ID
	companyIDStr := sessions.Get(c, "company_ids")
	companyIDUint64, err := strconv.ParseUint(companyIDStr, 10, 16)
	companyID := uint16(companyIDUint64)

	if err != nil {
		errors.Init()
		errors.Add("Name", err.Error())
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))
		c.Redirect(http.StatusFound, "/register")
		return
	}

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	if controller.userService.CheckUserExists(request.Email) {
		errors.Init()
		errors.Add("Email", "Email address already exists!")
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}
	request.CompanyID = companyID
	user, err := controller.userService.Create(request)

	//Check if there is any error on the user page
	if err != nil {
		fmt.Println("Hi here is the error", err)
		c.Redirect(http.StatusFound, "/register")
		return
	}

	//sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	// after creating the user redirect to the home page
	log.Printf("The user created successfully with a name %s \n", user.Name)
	c.Redirect(http.StatusFound, "/")
}

/*
Login handles the user login process by rendering the "login" template.
Parameters:
  - c (*gin.Context): The Gin context object representing the HTTP request and response.
*/
func (controller *Controller) Login(c *gin.Context) {
	htmlparse.Render(c, http.StatusOK, "modules/user/html/login", gin.H{
		"title": "Login",
	})
}

func (controller *Controller) HandleLogin(c *gin.Context) {
	// validate the request
	var request auth.LoginRequest

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := controller.userService.HandleUserLogin(request)
	if err != nil {
		errors.Init()
		errors.Add("email", err.Error())
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/login")
		return
	}
	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))
	sessions.Set(c, "company_id", strconv.Itoa(int(user.CompanyID)))

	log.Printf("The user logged in successfully with a name %s \n", user.Name)
	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) HandleLogout(c *gin.Context) {
	sessions.Remove(c, "auth")
	c.Redirect(http.StatusFound, "/")
}
