package handler

import (
	pb "api-gateway/genproto/budgetpb"
	"api-gateway/logger"

	"github.com/gin-gonic/gin"
)

// CreateBudget godoc
// @Summary Create a new budget
// @Description Create a new budget with the provided details
// @Tags budgets
// @Accept json
// @Produce json
// @Param budget body models.CreateBudgetRequest true "Budget details"
// @Success 200 {object} models.CreateBudgetResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /budgets [post]
// @Security BearerAuth
func (h *HandlerST) CreateBudget(c *gin.Context) {

	req := pb.CreateBudgetReq{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("CreateBudget: JSON binding error - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Service.CreateBudget(ctx, &req)
	if err != nil {
		logger.Error("CreateBudget: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("CreateBudget: Successfully created budget - ", resp.Budget.Id)
	c.JSON(200, resp)
}

// UpdateBudget godoc
// @Summary Update an existing budget
// @Description Update an existing budget with the provided ID
// @Tags budgets
// @Accept json
// @Produce json
// @Param id path string true "Budget ID"
// @Param budget body models.UpdateBudgetRequest true "Budget details"
// @Success 200 {object} models.UpdateBudgetResponse
// @Failure 500 {object} string
// @Router /budgets/{id} [put]
// @Security BearerAuth
func (h *HandlerST) UpdateBudget(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.Service.UpdateBudget(ctx, &pb.UpdateBudgetReq{
		Id: id,
	})

	if err != nil {
		logger.Error("GetListBudget: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// logger.Info("GetListBudget: Successfully retrieved list budgets - ", len((resp.List)))
	c.JSON(200, resp)
}

// DeleteBudget godoc
// @Summary Delete a budget
// @Description Delete a budget with the provided ID
// @Tags budgets
// @Accept json
// @Produce json
// @Param id path string true "Budget ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /budgets/{id} [delete]
// @Security BearerAuth
func (h *HandlerST) DeleteBudget(c *gin.Context) {

	req := pb.DeleteBudgetReq{}
	req.Id = c.Param("id")
	if err := c.BindJSON(&req); err != nil {
		logger.Error("UpdateBudget: JSON binding error - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Service.DeleteBudget(ctx, &req)
	if err != nil {
		logger.Error("UpdateBudget: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("UpdateBudhet: Successfully updated budget")
	c.JSON(200, resp)
}

// GetBudgetById godoc
// @Summary Get a budget by ID
// @Description Get details of a specific budget by its ID
// @Tags budgets
// @Accept json
// @Produce json
// @Param id path string true "Budget ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /budgets/{id} [get]
// @Security BearerAuth
func (h *HandlerST) GetBudgetById(c *gin.Context) {

	req := pb.GetBudgetByIdReq{}
	req.Id = c.Param("id")
	if err := c.BindJSON(&req); err != nil {
		logger.Error("UpdateBudget: JSON binding error - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Service.GetBudgetById(ctx, &req)
	if err != nil {
		logger.Error("UpdateBudget: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("UpdateBudhet: Successfully updated budget")
	c.JSON(200, resp)
}

// GetBudgets godoc
// @Summary Get all budgets
// @Description Get a list of all budgets
// @Tags budgets
// @Accept json
// @Produce json
// @Success 200 {object} models.GetBudgetsResp
// @Failure 500 {object} string
// @Router /budgets [get]
// @Security BearerAuth
func (h *HandlerST) GetBudgets(c *gin.Context) {

	req := pb.GetBudgetsReq{}
	resp, err := h.Service.GetBudgets(ctx, &req)
	if err != nil {
		logger.Error("UpdateBudget: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("UpdateBudhet: Successfully updated budget")
	c.JSON(200, resp)
}
