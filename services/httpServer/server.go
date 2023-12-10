package httpServer

import (
	"EmployeesApiService/data/stores"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type HttpServer struct {
	router *gin.Engine
	store  *stores.RedisStore
}

func NewServer(store *stores.RedisStore) HttpServer {

	router := gin.Default()
	initController(router, store)

	return HttpServer{router, store}
}

func (receiver HttpServer) Start() {
	err := receiver.router.Run(":8080")
	if err != nil {
		log.Panic(err.Error())
	}
}

func initController(router *gin.Engine, store *stores.RedisStore) {

	router.GET("/employees/", func(cnx *gin.Context) {
		employees, err := store.Employees()
		if err != nil {
			cnx.JSON(http.StatusNotFound, err)
		} else {
			cnx.JSON(http.StatusOK, employees)
		}
	})

	router.GET("/employee/:key", func(cnx *gin.Context) {
		key := cnx.Param("key")
		employee, err := store.Employee(key)
		if err != nil {
			cnx.JSON(http.StatusNotFound, err)
		} else {
			cnx.JSON(http.StatusOK, employee)
		}
	})
}