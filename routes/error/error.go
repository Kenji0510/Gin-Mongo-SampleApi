package error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status       string  `json:"status"`
	ErrorMessage *string `json:"error_message,omitempty"`
}

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	DatabaseError       = &AppError{Code: http.StatusInternalServerError, Message: "Database error"}
	NotFoundError       = &AppError{Code: http.StatusNotFound, Message: "Resource not found"}
	BadRequestError     = &AppError{Code: http.StatusBadRequest, Message: "Bad request"}
	InternalServerError = &AppError{Code: http.StatusInternalServerError, Message: "Database error"}
	ConflictError       = &AppError{Code: http.StatusConflict, Message: "Conflict error"}
)

func HandleErrorResponse(ctx *gin.Context, appErr *AppError) {
	errorResponse := ErrorResponse{
		Status:       "Error",
		ErrorMessage: &appErr.Message,
	}
	ctx.JSON(appErr.Code, errorResponse)
}
