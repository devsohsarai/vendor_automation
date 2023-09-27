package seeder

import (
	"fmt"
	"log"

	articleModel "github.com/gowaves/vendor_automation/internal/modules/article/models"
	userModel "github.com/gowaves/vendor_automation/internal/modules/user/models"
	"github.com/gowaves/vendor_automation/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	// 1 FFFFF[Access key]YYYYYYYTYRTYR : "Test-DATA10001[AKIAUCJATANCNHUXX3YV]||jjjjj}}}-Test DATA"
	// 2 Secret access key : "Test-Data237877474$5785@#$$[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]784735juyhgjaf"
	// 3 ARNO :  "TETSTSSTTSTST%%$$#$#$(36383hjehfdbbkfa)[arn1234:aws1234:iam1234::279778100036:user/eSeller]||jjjjj}}}-Test DATA"
	//4 RT :Atzr|IwEBIChEgTyWiMBhVZIf2WD_HgklKztN-QsrQMkeZluv3GsArq1LaVEVbNpVlQiLxawm8p_wHfIVbD5OJeMqeYU7GyE852sPXTLoRbylV21iOh6pt692Wo41r7l9zdouEqsId12IfooPaaNskvgJqGsR12j55CZ3RquKYkLljZ-mC41-FsPxhy1b0_yS4jbavvrzoIsIi5n_OtRIm0lxp4Y4V-ROSEPQZP0mTmDWtELWDxdgCgdozGRLHumuVsiVfUyRnFvWhtPGysZmWL6dn6thN4VL6txAF8LnW3CXz9_M8x0MnnOULhhSaIYvJyPiEkS9izU
	//CID :  TETSTSSTTSTST%%$$#$#$(36383hjehfdbbkfa)[amzn1.application-oa2-client.2e97989e6029401bbb9a55976f125c2a]TETSTSSTTSTST%%$$#$#$(36383hjehfdbbkfa)
	//CSecret:  TETSTSSTTSTST%%$$#$#$(36383hjehfdbbkfa)[amzn1.oa2-cs.v1.8b9a50bf2f857558f599f8cfd4428a160f8cf7646f6eb2223d36572ac940f6bf]TETSTSSTTSTST%%$$#$#$(36383hjehfdbbkfa)
	db := database.Connection()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte("ravinesh1234"), 12)
	if err != nil {
		log.Fatal("hash password error")
		return
	}
	company := userModel.Company{
		Name:    "Shanti Digital",
		Email:   "info@shantidigital.com",
		Contact: "9930319945",
		Address: "Cypress Texas",
		Status:  true,
	}
	db.Create(&company)

	user := userModel.User{
		Name:      "Viart  Arsh",
		Email:     "virat@gmail.com",
		Password:  string(hashPassword),
		Contact:   "9930319931",
		IsAdmin:   true,
		CompanyID: company.ID,
	}

	db.Create(&user) // pass pointer of data to Create

	log.Printf("User created successfully with email address %s\n", user.Email)

	for i := 1; i <= 10; i++ {
		article := articleModel.Article{Title: fmt.Sprintf("random title %d", i), Content: fmt.Sprintf("random content %d", i), UserID: user.ID}
		db.Create(&article)
		log.Printf("Article created successfully with title %s \n", article.Title)
	}

	log.Println("Seeder done ..")
}
