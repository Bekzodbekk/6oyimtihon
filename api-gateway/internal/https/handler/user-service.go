package handler

import (
	pb "api-gateway/genproto/userpb"
	"api-gateway/internal/https/token"
	"api-gateway/logger"

	"github.com/gin-gonic/gin"
)

// @Router /api/v1/users [post]
// @Summary REGISTER USERS
// @Security  		BearerAuth
// @Description This method registers a new user
// @Tags AUTH
// @Accept json
// @Produce json
// @Param user body models.RegisterUserRequest true "User"
// @Success 200 {object} models.RegisterUserResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) RegisterUser(c *gin.Context) {

	req := pb.RegisterUserRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("RegisterUser: JSON binding error - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Service.RegisterUser(ctx, &req)
	if err != nil {
		logger.Error("RegisterUser: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("RegisterUser: Successfully registered user - ", resp.UserName)
	c.JSON(200, resp)
}

// @Router /api/v1/users/login [post]
// @Summary LOGIN USERS
// @Security  		BearerAuth
// @Description This method login users
// @Security BearerAuth
// @Tags AUTH
// @Accept json
// @Produce json
// @Param user body models.LoginUserRequest true "User"
// @Success 200 {object} token.Tokens
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) LoginUser(c *gin.Context) {

	req := pb.LoginUserRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("LoginUser: JSON binding error - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Service.LoginUser(ctx, &req)
	if err != nil {
		logger.Error("LoginUser: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	token := token.GenereteJWTToken(resp)
	logger.Info("LoginUser: Successfully logged in user - ", resp.Email)
	c.JSON(200, token)
}

// @Router /api/v1/users/profile/{id} [get]
// @Summary GET USER BY ID
// @Security  		BearerAuth
// @Description This method retrieves a user by their ID
// @Tags USERS
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.GetUserByIdResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetUserById(c *gin.Context) {

	req := pb.GetUserByIdRequest{}
	req.UserId = c.Param("id")
	resp, err := h.Service.GetUserById(ctx, &req)
	if err != nil {
		logger.Error("GetUserById: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("GetUserById: Successfully retrieved user - ID: ", resp.User.Email)
	c.JSON(200, resp)
}

// @Router /api/v1/users/profile [get]
// @Summary GET USERS LIST
// @Security  		BearerAuth
// @Description This method retrieves a list of users
// @Tags USERS
// @Accept json
// @Produce json
// @Success 200 {object} models.GetUsersResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetUsersList(c *gin.Context) {

	resp, err := h.Service.GetUsersList(ctx, &pb.GetUsersRequest{})
	if err != nil {
		logger.Error("GetUsersList: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("GetUsersList: Successfully retrieved users list - ", len(resp.List))
	c.JSON(200, resp)
}
