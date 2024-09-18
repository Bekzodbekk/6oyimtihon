package https

import (
	"api-gateway/internal/https/handler"
	auth "api-gateway/internal/https/middleware/authorization"
	rlimit "api-gateway/internal/https/middleware/rate-limiting"
	"api-gateway/internal/service"
	"crypto/tls"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Api-gateway service
// @version 1.0
// @description Api-gateway service
// @host localhost:9000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(service *service.ServiceRepositoryClient) *http.Server {

	r := gin.Default()

	apiHandler := handler.NewApiHandler(service)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(auth.MiddleWare())
	r.Use(rlimit.RateLimit())

	api := r.Group("/api/v1")
	{
		api.POST("/users", apiHandler.RegisterUser)
		api.POST("/users/login", apiHandler.LoginUser)
		api.POST("/users/verify", apiHandler.VerifyUser)
		api.PUT("/users/:id", apiHandler.UpdateUser)
		api.DELETE("/users/:id", apiHandler.DeleteUser)
		api.GET("/users/profile/:id", apiHandler.GetUserById)

		api.POST("/transactions/income", apiHandler.RegisterIncome)
		api.POST("/transactions/expense", apiHandler.RegisterExpense)
		api.GET("/transactions", apiHandler.GetListIncomeVSExpense)

		api.POST("/budgets", apiHandler.CreateBudget)
		api.PUT("/budgets/:id", apiHandler.UpdateBudget)
		api.GET("/budgets/:id", apiHandler.GetBudgetById)
		api.GET("/budgets", apiHandler.GetBudgets)
		api.DELETE("/budgets/:id", apiHandler.DeleteBudget)

		api.GET("/reports/income-expense", apiHandler.GetTotalReports)
		api.GET("/reports/spending-by-category", apiHandler.GetReportsByCategory)

	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:      ":9000",
		Handler:   r,
		TLSConfig: tlsConfig,
	}

	return srv
}
