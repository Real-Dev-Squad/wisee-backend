package routes

import (
	"net/http"

	"github.com/Real-Dev-Squad/wisee-backend/src/dtos"
	"github.com/Real-Dev-Squad/wisee-backend/src/models"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func FormRoutes(rg *gin.RouterGroup, db *bun.DB) {
	forms := rg.Group("/forms")

	forms.POST("", func(ctx *gin.Context) {
		var requestBody dtos.CreateFormRequestDto
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			errObj := dtos.ResponseDto{
				Message: "invalid request",
				Error: &dtos.ErrorResponse{
					Message: "invalid request body",
					Detail:  err.Error(),
				},
			}
			ctx.JSON(http.StatusBadRequest, errObj)
			return
		}

		var form = &models.Form{
			Content:     requestBody.Content,
			CreatedById: requestBody.PerformedById,
			Status:      models.DRAFT,
			OwnerId:     requestBody.PerformedById,
			// TODO : generate shareable id
			ShareableId: "123",
		}

		// Create a new form
		// ---
		if _, err := db.NewInsert().Model(form).Exec(ctx); err != nil {
			errObj := dtos.ResponseDto{
				Message: "something went wrong",
				Error: &dtos.ErrorResponse{
					Message: "error creating form",
					Detail:  err.Error(),
				},
			}
			ctx.JSON(http.StatusBadRequest, errObj)
			return
		}

		var FormMetaData = &models.FormMetaData{
			FormId: form.Id,
		}

		// Create a new entry using the form created above
		// ---
		if _, err := db.NewInsert().Model(FormMetaData).Exec(ctx); err != nil {
			errObj := dtos.ResponseDto{
				Message: "something went wrong",
				Error: &dtos.ErrorResponse{
					Message: "error creating form meta data",
					Detail:  err.Error(),
				},
			}
			ctx.JSON(http.StatusBadRequest, errObj)
			return
		}

		var resData = dtos.CreateUpdateGetFormResponseDto{
			Id:          form.Id,
			Content:     form.Content,
			OwnerId:     form.OwnerId,
			CreatedById: form.CreatedById,
			Status:      string(form.Status),
			CreatedAt:   form.CreatedAt.String(),
			UpdatedAt:   form.UpdatedAt.String(),
		}

		resObj := dtos.ResponseDto{
			Message: "form created successfully",
			Data:    resData,
		}

		ctx.JSON(http.StatusCreated, resObj)
	})

	forms.GET("", func(ctx *gin.Context) {
		var form []models.Form
		if err := db.NewSelect().Model(&form).OrderExpr("id ASC").Scan(ctx); err != nil {
			errObj := dtos.ResponseDto{
				Message: "something went wrong",
				Error: &dtos.ErrorResponse{
					Message: "error fetching forms",
					Detail:  err.Error(),
				},
			}
			ctx.JSON(http.StatusBadRequest, errObj)
			return
		}

		var resData dtos.GetFormsResponseDto
		for _, f := range form {
			resData = append(resData, dtos.CreateUpdateGetFormResponseDto{
				Id:          f.Id,
				Content:     f.Content,
				OwnerId:     f.OwnerId,
				CreatedById: f.CreatedById,
				UpdatedById: f.UpdatedById,
				Status:      string(f.Status),
				CreatedAt:   f.CreatedAt.String(),
				UpdatedAt:   f.UpdatedAt.String(),
			})
		}

		var res = dtos.ResponseDto{
			Message: "forms fetched successfully",
			Data:    resData,
		}

		ctx.JSON(http.StatusOK, res)
	})

	forms.GET("/:id", func(ctx *gin.Context) {
		var formMetaData models.FormMetaData
		var form models.Form

		query := db.NewSelect().Model(&formMetaData).Relation("Form").Where("form_id = ?", ctx.Param("id"))
		if err := query.Scan(ctx); err != nil {
			errObj := dtos.ResponseDto{
				Message: "something went wrong",
				Error: &dtos.ErrorResponse{
					Message: "error fetching form",
					Detail:  err.Error(),
				},
			}
			ctx.JSON(http.StatusBadRequest, errObj)
			return
		}

		form = *formMetaData.Form

		var resData = dtos.GetFormDetailResponseDto{
			Id:          form.Id,
			OwnerId:     form.OwnerId,
			Status:      string(form.Status),
			CreatedById: form.CreatedById,
			UpdatedById: form.UpdatedById,
			CreatedAt:   form.CreatedAt.String(),
			UpdatedAt:   form.UpdatedAt.String(),
			Content:     form.Content,
			Meta: dtos.GetFormMetaDataResponseDto{
				Id:                               formMetaData.Id,
				FormId:                           formMetaData.FormId,
				IsDeleted:                        formMetaData.IsDeleted,
				AccepctingResponses:              formMetaData.AccepctingResponses,
				AllowGuestResponses:              formMetaData.AllowGuestResponses,
				AllowMultipleRepsonses:           formMetaData.AllowMultipleRepsonses,
				SendConfirmationEmailToRespondee: formMetaData.SendConfirmationEmailToRespondee,
				SendSubmissionEmailToOwner:       formMetaData.SendSubmissionEmailToOwner,
				ValidTill:                        formMetaData.ValidTill,
				UpdatedById:                      formMetaData.UpdatedById,
				UpdatedAt:                        formMetaData.UpdatedAt,
			},
		}

		var res = dtos.ResponseDto{
			Message: "form fetched successfully",
			Data:    resData,
		}

		ctx.JSON(http.StatusOK, res)
	})

	forms.PATCH("/:id", func(ctx *gin.Context) {
		var requestBody dtos.UpdateFormRequestDto
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			errObj := dtos.ResponseDto{
				Message: "invalid request",
				Error: &dtos.ErrorResponse{
					Message: "invalid request body",
					Detail:  err.Error(),
				},
			}
			ctx.JSON(http.StatusBadRequest, errObj)
			return
		}

		var form models.Form
		if err := db.NewSelect().Model(&form).Where("id = ?", ctx.Param("id")).Scan(ctx); err != nil {
			errObj := dtos.ResponseDto{
				Message: "something went wrong",
				Error: &dtos.ErrorResponse{
					Message: "error fetching form",
					Detail:  err.Error(),
				},
			}
			ctx.JSON(http.StatusBadRequest, errObj)
			return
		}

		if requestBody.Status != string(models.DRAFT) && requestBody.Status != string(models.PUBLISHED) {
			errObj := dtos.ResponseDto{
				Message: "invalid request",
				Error: &dtos.ErrorResponse{
					Message: "invalid status",
				},
			}
			ctx.JSON(http.StatusBadRequest, errObj)
			return
		}

		form.Content = requestBody.Content
		form.Status = models.FORM_STATUS_TYPE(requestBody.Status)
		form.OwnerId = requestBody.PerformedById
		form.UpdatedById = &requestBody.PerformedById

		if _, err := db.NewUpdate().Model(&form).Where("id = ?", ctx.Param("id")).Exec(ctx); err != nil {
			errObj := dtos.ResponseDto{
				Message: "something went wrong",
				Error: &dtos.ErrorResponse{
					Message: "error updating form",
					Detail:  err.Error(),
				},
			}
			ctx.JSON(http.StatusBadRequest, errObj)
			return
		}

		var resData = dtos.CreateUpdateGetFormResponseDto{
			Id:          form.Id,
			Content:     form.Content,
			OwnerId:     form.OwnerId,
			CreatedById: form.CreatedById,
			Status:      string(form.Status),
			CreatedAt:   form.CreatedAt.String(),
			UpdatedAt:   form.UpdatedAt.String(),
			UpdatedById: form.UpdatedById,
		}

		resObj := dtos.ResponseDto{
			Message: "form updated successfully",
			Data:    resData,
		}

		ctx.JSON(http.StatusAccepted, resObj)
	})
}
