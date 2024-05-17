package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavohroos/gopportunities/schemas"
)

// @BasePath		/api/v1
// @Summary		Delete opening
// @Description	Delete a job opening
// @Tags			Openings
// @Accept			json
// @Produce		json
// @Param			id	query		string	true	"Opening ID"
// @Success		200	{object}	DeleteOpeningResponse
// @Failure		400	{object}	ErrorResponse
// @Failure		404	{object}	ErrorResponse
// @Failure		500	{object}	ErrorResponse
// @Router			/opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error deleting opening")
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}
