package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewDB struct {
	Name string `json:"name"`
}

func main() {
	router := gin.Default()
	router.GET("/db-management/databases", getMultipleDBs)
	router.GET("/db-management/databases/:id", getSingleDBs)
	router.POST("/db-management/databases", createSingleDB)
	router.PUT("/db-management/databases/:id", updateSingleDB)
	router.Run("localhost:8080")
}

func getMultipleDBs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Multiple DBs")
}

func getSingleDBs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Single DB: "+c.Param("id"))
}

func createSingleDB(c *gin.Context) {
	var newDB NewDB

	if err := c.ShouldBindJSON(&newDB); err != nil {
	}

	c.IndentedJSON(http.StatusCreated, newDB)
}

func updateSingleDB(c *gin.Context) {
	var newDB NewDB

	if err := c.ShouldBindJSON(&newDB); err != nil {
	}

	c.IndentedJSON(http.StatusOK, "Single DB: "+c.Param("id")+" -> "+newDB.Name)
}
