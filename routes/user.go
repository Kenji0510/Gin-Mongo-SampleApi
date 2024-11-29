package routes

import (
	"gin-api-sample/domain/users/models"
	userUsercase "gin-api-sample/domain/users/usecase"
	ErrorPesp "gin-api-sample/routes/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsercase userUsercase.UserUsecaseI
}

func New(userService userUsercase.UserUsecaseI) UserController {
	return UserController{
		UserUsercase: userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ErrorPesp.HandleErrorResponse(ctx, ErrorPesp.BadRequestError)
		return
	}
	err := uc.UserUsercase.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully Inserted Data User"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	var userName = ctx.Param("name")
	user, err := uc.UserUsercase.GetUser(ctx, &userName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserUsercase.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.UserUsercase.UpdateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully Updated Data User"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	var userName = ctx.Param("name")
	err := uc.UserUsercase.DeleteData(ctx, &userName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted Data User"})
}
