package routes

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Real-Dev-Squad/wisee-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauth2Api "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func getUserInfoFromCode(code string, conf *oauth2.Config, ctx *gin.Context) *oauth2Api.Userinfo {
	tok, err := conf.Exchange(context.TODO(), code)

	if err != nil {
		log.Fatal(err)
	}

	oauth2Service, err := oauth2Api.NewService(ctx, option.WithTokenSource(conf.TokenSource(ctx, tok)))

	if err != nil {
		log.Fatal(err)
	}

	userInfo, err := oauth2Service.Userinfo.Get().Context(ctx).Do()

	if err != nil {
		log.Fatal(err)
	}

	return userInfo
}

func AuthRoutes(reg *gin.RouterGroup, db *bun.DB) {
	auth := reg.Group("/auth")

	googleAuth := auth.Group("/google")
	conf := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	googleAuth.GET("/login", func(ctx *gin.Context) {
		url := conf.AuthCodeURL("state")
		fmt.Printf("Visit the URL for the auth dialog: %v", url)

		ctx.Redirect(302, url)
	})

	googleAuth.GET(("/callback"), func(ctx *gin.Context) {
		code := ctx.Query("code")
		googleAccountInfo := getUserInfoFromCode(code, conf, ctx)

		user := new(models.User)

		count, _ := db.NewSelect().Model(user).Where("email = ?", googleAccountInfo.Email).ScanAndCount(ctx)

		if count != 0 {
			// create a auth token

			// set cookies here
			ctx.SetCookie("email", googleAccountInfo.Email, 3600, "/", "localhost", false, true)
			ctx.SetCookie("username", googleAccountInfo.Name, 3600, "/", "localhost", false, true)

			// redirect user to dashboard
			ctx.Redirect(302, "http://localhost:3000/dashboard")

			return
		}

		newUser := &models.User{
			Username: googleAccountInfo.Name,
			Email:    googleAccountInfo.Email,
		}

		// create account
		_, err := db.NewInsert().Model(newUser).Exec(ctx)

		if err != nil {
			// log.Fatal(err)
			ctx.JSON(500, gin.H{
				"message": "error",
			})
		}

		// create token

		// set cookie
		ctx.SetCookie("email", googleAccountInfo.Email, 3600, "/", "localhost", false, true)
		ctx.SetCookie("username", googleAccountInfo.Name, 3600, "/", "localhost", false, true)

		// redirect
		ctx.JSON(302, "http://localhost:3000/onboarding")

	})
}
