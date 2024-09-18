package handler

import (
	pb "api-gateway/genproto/budgetpb"
	"api-gateway/logger"

	"github.com/gin-gonic/gin"
)

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
