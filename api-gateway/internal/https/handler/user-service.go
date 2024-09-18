package handler

import (
	"api-gateway/genproto/userpb"
	"api-gateway/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description This method registers a new user
// @Tags AUTH
// @Accept json
// @Produce json
// @Param user body models.RegisterUserRequest true "User registration details"
// @Success 200 {object} models.RegisterUserResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/users [post]
func (h *HandlerST) RegisterUser(c *gin.Context) {
	req := userpb.CreateUserReq{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("RegisterUser: JSON binding error - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Service.RegisterUser(c, &req)
	fmt.Println(resp)
	if err != nil {
		logger.Error("RegisterUser: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// logger.Info("RegisterUser: Successfully registered user - ", resp.User.Username)
	c.JSON(200, resp)
}

// LoginUser godoc
// @Summary Login user
// @Description This method logs in a user
// @Tags AUTH
// @Accept json
// @Produce json
// @Param user body models.LoginUserRequest true "User login credentials"
// @Success 200 {object} token.Tokens
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/users/login [post]
func (h *HandlerST) LoginUser(c *gin.Context) {
	req := userpb.LoginReq{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("LoginUser: JSON binding error - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Service.LoginUser(c, &req)
	if err != nil {
		logger.Error("LoginUser: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	logger.Info("LoginUser: Successfully logged in user - ", resp.Token)
	c.JSON(200, resp)
}

// VerifyUser godoc
// @Summary Verify user
// @Description This method verifies a user
// @Tags AUTH
// @Accept json
// @Produce json
// @Param user body models.VerifyUserReq true "User verification details"
// @Success 200 {object} models.VerifyUserResp
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/users/verify [post]
func (h *HandlerST) VerifyUser(c *gin.Context) {
	req := userpb.VerifyUserReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	resp, err := h.Service.VerifyUser(c, &req)
	if err != nil {
		logger.Error("VerifyUser: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("VerifyUser: Successfully verified user - Email: ", resp.User.Email)
	c.JSON(200, resp)
}

// UpdateUser godoc
// @Summary Update user
// @Description This method updates a user's information
// @Tags USERS
// @Accept json
// @Produce json
// @Security  		BearerAuth
// @Param id path string true "User ID"
// @Param user body models.UpdateUserReq true "User update details"
// @Success 200 {object} models.UpdateUserResp
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/users/{id} [put]
func (h *HandlerST) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	req := userpb.UpdateUserReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	req.Id = id

	resp, err := h.Service.UpdateUser(c, &req)
	if err != nil {
		logger.Error("UpdateUser: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("UpdateUser: Successfully updated user - Email: ", resp.User.Email)
	c.JSON(200, resp)
}

// DeleteUser godoc
// @Summary Delete user
// @Description This method deletes a user
// @Tags USERS
// @Accept json
// @Produce json
// @Security  		BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.DeleteUserRes
// @Failure 500 {object} string
// @Router /api/v1/users/{id} [delete]
func (h *HandlerST) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	req := userpb.DeleteUserReq{}
	req.Id = id

	resp, err := h.Service.DeleteUser(c, &req)
	if err != nil {
		logger.Error("DeleteUser: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

// GetUserById godoc
// @Summary Get user by ID
// @Description This method retrieves a user by their ID
// @Tags USERS
// @Accept json
// @Produce json
// @Security  		BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.GetUserByIdResponse
// @Failure 500 {object} string
// @Router /api/v1/users/profile/{id} [get]
func (h *HandlerST) GetUserById(c *gin.Context) {
	req := userpb.GetUserByIdReq{}
	req.Id = c.Param("id")
	resp, err := h.Service.GetUserById(c, &req)
	if err != nil {
		logger.Error("GetUserById: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("GetUserById: Successfully retrieved user - ID: ", resp.User.Email)
	c.JSON(200, resp)
}
