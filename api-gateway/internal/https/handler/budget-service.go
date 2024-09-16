package handler

import (
	pb "api-gateway/genproto/budgetpb"
	"api-gateway/logger"

	"github.com/gin-gonic/gin"
)

// @Router /api/v1/budgets [post]
// @Summary CREATE BUDGETS
// @Security  		BearerAuth
// @Description This method creates budget
// @Tags BUDGETS
// @Accept json
// @Produce json
// @Param budget body models.CreateBudgetRequest true "Budget"
// @Success 200 {object} models.CreateBudgetResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) CreateBudget(c *gin.Context) {

	req := pb.CreateBudgetRequest{}
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
	logger.Info("CreateBudget: Successfully created budget - ", resp.BudgetId)
	c.JSON(200, resp)
}

// @Router /api/v1/budgets [get]
// @Summary GET BUDGETS
// @Security  		BearerAuth
// @Description This method gets budgets
// @Tags BUDGETS
// @Accept json
// @Produce json
// @Success 200 {object} models.GetListBudgetResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetListBudget(c *gin.Context) {

	resp, err := h.Service.GetListBudget(ctx, &pb.GetListBudgetRequest{})
	if err != nil {
		logger.Error("GetListBudget: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("GetListBudget: Successfully retrieved list budgets - ", len((resp.List)))
	c.JSON(200, resp)
}

// @Router /api/v1/budgets/{id} [put]
// @Summary UPDATE BUDGETS
// @Security  		BearerAuth
// @Description This method updates budgets
// @Tags BUDGETS
// @Accept json
// @Produce json
// @Param id path string true "Budget id"
// @Param budget body models.UpdateBudgetRequest true "budget"
// @Success 200 {object} models.UpdateBudgetResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) UpdateBudget(c *gin.Context) {

	req := pb.UpdateBudgetRequest{}
	req.BudgetId = c.Param("id")
	if err := c.BindJSON(&req); err != nil {
		logger.Error("UpdateBudget: JSON binding error - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Service.UpdateBudget(ctx, &req)
	if err != nil {
		logger.Error("UpdateBudget: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("UpdateBudhet: Successfully updated budget")
	c.JSON(200, resp)
}
