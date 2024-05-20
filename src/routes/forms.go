package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Real-Dev-Squad/wisee-backend/src/dtos"
	"github.com/Real-Dev-Squad/wisee-backend/src/models"
	"github.com/Real-Dev-Squad/wisee-backend/src/utils"
	"github.com/Real-Dev-Squad/wisee-backend/src/utils/logger"
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

		/**
		* Stringify the content to generate a hash
		* ---
		 */
		stringifiedJson, marshalErr := json.Marshal(requestBody.Content)

		if marshalErr != nil {
			errObj := dtos.ResponseDto{
				Message: "invalid request",
				Error: &dtos.ErrorResponse{
					Message: "invalid request body",
					Detail:  "error reading the content",
				},
			}

			ctx.JSON(http.StatusBadRequest, errObj)
			return
		}

		// create a key from the content and the id of the user who created the form
		shareableIdKey := string(stringifiedJson) + strconv.FormatInt(requestBody.PerformedById, 10)
		shareableId, hashErr := utils.GenerateHash(shareableIdKey, 5)

		if hashErr != nil {
			errObj := dtos.ResponseDto{
				Message: "invalid request",
				Error: &dtos.ErrorResponse{
					Message: "invalid request body",
					Detail:  "error generating shareable id",
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
			ShareableId: shareableId,
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
			ShareableId: form.ShareableId,
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
			ShareableId: form.ShareableId,
			Meta: dtos.GetFormMetaDataResponseDto{
				Id:                               formMetaData.Id,
				FormId:                           formMetaData.FormId,
				IsDeleted:                        formMetaData.IsDeleted,
				AcceptingResponses:               formMetaData.AcceptingResponses,
				AllowGuestResponses:              formMetaData.AllowGuestResponses,
				AllowMultipleResponses:           formMetaData.AllowMultipleResponses,
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

	forms.GET("/share/:shareableId", func(ctx *gin.Context) {
		var formMetaData models.FormMetaData
		var form models.Form
		shareable_id := ctx.Param("shareableId")

		query := db.NewSelect().Model(&formMetaData).Relation("Form").Where("shareable_id = ?", shareable_id)
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

		// TODO - @yesyash : trim down the response size (createdAt, updatedAt), we don't need all the fields
		var resData = dtos.GetFormDetailResponseDto{
			Id:          form.Id,
			OwnerId:     form.OwnerId,
			Status:      string(form.Status),
			CreatedById: form.CreatedById,
			UpdatedById: form.UpdatedById,
			CreatedAt:   form.CreatedAt.String(),
			UpdatedAt:   form.UpdatedAt.String(),
			Content:     form.Content,
			ShareableId: form.ShareableId,
			Meta: dtos.GetFormMetaDataResponseDto{
				Id:                               formMetaData.Id,
				FormId:                           formMetaData.FormId,
				IsDeleted:                        formMetaData.IsDeleted,
				AcceptingResponses:               formMetaData.AcceptingResponses,
				AllowGuestResponses:              formMetaData.AllowGuestResponses,
				AllowMultipleResponses:           formMetaData.AllowMultipleResponses,
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

	forms.POST("/:shareableId/respond", func(ctx *gin.Context) {
		shareableId := ctx.Param("shareableId")
		var requestBody dtos.CreateFormSubmissionRequestDto

		var formData struct {
			ID                 int64  `bun:"id"`
			ShareableID        string `bun:"shareable_id"`
			AcceptingResponses bool   `bun:"accepting_responses"`
		}

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

		query := db.NewRaw("SELECT form.id, forms.form.shareable_id, forms.metadata.accepting_responses FROM forms.form JOIN forms.metadata ON form.id = metadata.form_id WHERE form.shareable_id = ?", shareableId)

		if err := query.Scan(ctx, &formData); err != nil {
			logger.Warn("error fetching form: ", err.Error())

			errObj := dtos.ResponseDto{
				Message: "something went wrong",
				Error: &dtos.ErrorResponse{
					Message: "error fetching form",
				},
			}

			ctx.JSON(http.StatusBadRequest, errObj)
			return
		}

		if !formData.AcceptingResponses {
			errObj := dtos.ResponseDto{
				Message: "invalid request",
				Error: &dtos.ErrorResponse{
					Message: "form is not accepting responses",
				},
			}

			ctx.JSON(http.StatusBadRequest, errObj)
			return
		}

		formResponse := &models.FormResponse{
			ResponseByID: requestBody.ResponseById,
			FormID:       formData.ID,
			Content:      requestBody.Content,
		}

		// TODO - @yesyash : if multiple responses are not allowed, throw an error if a response already exists from a responseById
		if _, err := db.NewInsert().Model(formResponse).Exec(ctx); err != nil {
			errObj := dtos.ResponseDto{
				Message: "something went wrong",
				Error: &dtos.ErrorResponse{
					Message: "error creating form response",
					Detail:  err.Error(),
				},
			}
			ctx.JSON(http.StatusBadRequest, errObj)
			return
		}

		resObj := dtos.ResponseDto{
			Message: "form response created successfully",
		}

		ctx.JSON(http.StatusCreated, resObj)
	})
}
