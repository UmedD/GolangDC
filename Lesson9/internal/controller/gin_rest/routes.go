package gin_rest

import "github.com/gin-gonic/gin"

func StartServer() {
	router := gin.Default()

	router.GET("/", Ping)
	router.GET("/accounts", GetAllAccounts)
	router.GET("/accounts/:id", GetAccountByID)
	router.POST("/create", CreateAccount)
	router.GET("/delete/:id", DeleteAccount)
	router.PATCH("/update/:id", UpdateAccount)
	router.Run(":8080")
}
