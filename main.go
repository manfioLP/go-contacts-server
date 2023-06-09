package go_contacts_server

import (
	"contacts-server/api"
	"contacts-server/db"
	"github.com/gin-gonic/gin"
)

func init() {
	db.OpenConnection()
	defer db.CloseConnection()

	router := gin.Default()
	{
		router.GET("/", api.BasicPath)
		router.GET("/health", api.Health)
		router.GET("/user", api.GetUser)
		router.GET("/users", api.GetUsers)
		router.PUT("/user", api.UpdateUser)
		router.POST("/user", api.CreateUser)
		router.DELETE("/user", api.DeleteUser)
	}
	router.Run(":3003")
	//	log.Fatal(http.ListenAndServe(":3003", router))
}
