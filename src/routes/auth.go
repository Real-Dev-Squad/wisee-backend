package routes

import (
	"context"

	"github.com/Real-Dev-Squad/wisee-backend/src/config"
	"github.com/Real-Dev-Squad/wisee-backend/src/models"
	"github.com/Real-Dev-Squad/wisee-backend/src/utils"
	"github.com/Real-Dev-Squad/wisee-backend/src/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauth2Api "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func getUserInfoFromCode(code string, conf *oauth2.Config, ctx *gin.Context) (*oauth2Api.Userinfo, error) {
	tok, exchangeErr := conf.Exchange(context.TODO(), code)

	if exchangeErr != nil {
		return nil, exchangeErr
	}

	oauth2Service, serviceError := oauth2Api.NewService(ctx, option.WithTokenSource(conf.TokenSource(ctx, tok)))

	if serviceError != nil {
		return nil, serviceError
	}

	userInfo, getInfoError := oauth2Service.Userinfo.Get().Context(ctx).Do()

	if getInfoError != nil {
		return nil, getInfoError
	}

	return userInfo, nil
}

func AuthRoutes(reg *gin.RouterGroup, db *bun.DB) {
	auth := reg.Group("/auth")
	googleAuth := auth.Group("/google")

	conf := &oauth2.Config{
		ClientID:     config.GoogleClientId,
		ClientSecret: config.GoogleClientSecret,
		RedirectURL:  config.GoogleRedirectUrl,

		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	googleAuth.GET("/login", func(ctx *gin.Context) {
		url := conf.AuthCodeURL("state")
		ctx.Redirect(302, url)
	})

	googleAuth.GET(("/callback"), func(ctx *gin.Context) {
		code := ctx.Query("code")
		domain := config.Domain
		authRedirectUrl := config.AuthRedirectUrl

		user := new(models.User)
		googleAccountInfo, getInfoError := getUserInfoFromCode(code, conf, ctx)

		if getInfoError != nil {
			logger.Fatal(getInfoError)
			ctx.JSON(500, gin.H{
				"message": "error",
			})
		}

		count, _ := db.NewSelect().Model(user).Where("email = ?", googleAccountInfo.Email).ScanAndCount(ctx)

		// User does not exist, create a new user
		if count == 0 {
			newUser := &models.User{
				Username: googleAccountInfo.Name,
				Email:    googleAccountInfo.Email,
			}

			_, err := db.NewInsert().Model(newUser).Exec(ctx)

			if err != nil {
				logger.Fatal(err)
				ctx.JSON(500, gin.H{
					"message": "error",
				})
			}

			user = newUser
		}

		token, err := utils.GenerateToken(user)

		if err != nil {
			logger.Fatal(err)
			ctx.JSON(500, gin.H{
				"message": "error",
			})
		}

		// set cookie and redirect
		ctx.SetCookie("token", token, 3600, "/", domain, true, true)
		ctx.Redirect(302, authRedirectUrl)
	})
}
