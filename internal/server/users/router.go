package user

import (
	"context"
	"user/internal/domain/models"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Users(ctx context.Context) ([]models.User, error)
	User(ctx context.Context, id string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User, id string) error
	DeleteUser(ctx context.Context, id string) error
}

type UserRouter struct {
	router  *gin.RouterGroup
	service UserService
}

func Register(service UserService, routerGroup *gin.RouterGroup) *UserRouter {
	router := UserRouter{service: service, router: routerGroup}
	router.init()
	return &router
}

func (r *UserRouter) Users(ctx *gin.Context) {
	users, err := r.service.Users(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, users)
}

func (r *UserRouter) User(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := r.service.User(ctx, id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, user)
}

func (r *UserRouter) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := r.service.CreateUser(ctx, &user); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, user)
}

func (r *UserRouter) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := r.service.UpdateUser(ctx, &user, id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, user)

}

func (r *UserRouter) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := r.service.DeleteUser(ctx, id); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "success delete user"})
}

func (r *UserRouter) init() {
	group := r.router.Group("/users")

	group.GET("/", r.Users)
	group.GET("/:id", r.User)
	group.POST("/", r.CreateUser)
	group.PUT("/:id", r.UpdateUser)
	group.DELETE("/:id", r.DeleteUser)
}
